package main

import (
	"context"
	"database/sql"
	"fmt"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/ThreeDotsLabs/watermill-jetstream/pkg/jetstream"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	watermillmiddleware "github.com/ThreeDotsLabs/watermill/message/router/middleware"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/microcosm-cc/bluemonday"
	"github.com/nats-io/nats.go"

	adapters2 "github.com/edanko/launches-api/internal/adapters"
	ent2 "github.com/edanko/launches-api/internal/adapters/ent"
	"github.com/edanko/launches-api/internal/adapters/ent/migrate"
	"github.com/edanko/launches-api/internal/app"
	commands2 "github.com/edanko/launches-api/internal/app/commands"
	events2 "github.com/edanko/launches-api/internal/app/events"
	queries2 "github.com/edanko/launches-api/internal/app/queries"
	"github.com/edanko/launches-api/internal/config"
	ports2 "github.com/edanko/launches-api/internal/ports"
	"github.com/edanko/launches-api/pkg/application"
	httputils "github.com/edanko/launches-api/pkg/http"
	httpmw "github.com/edanko/launches-api/pkg/http/middleware"
	"github.com/edanko/launches-api/pkg/logs"
	"github.com/edanko/launches-api/pkg/metrics"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer stop()

	cfg := config.GetConfig()

	// if cfg.App.Environment == "development" {
	// 	// zeroLogger = zeroLogger.Level(zerolog.DebugLevel)
	// }
	metricsClient := metrics.NoOp{}
	logger := logs.NewZerologLogger(cfg.Logger.Level)
	watermillLogger := logs.NewWatermillLogger(logger)

	db, err := sql.Open("pgx", fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Database, cfg.DB.User, cfg.DB.Password, cfg.DB.SSLMode),
	)
	if err != nil {
		logger.Fatal("failed to open connection to database", err, nil)
	}
	defer db.Close()

	dbClient := ent2.NewClient(
		ent2.Driver(entsql.OpenDB(dialect.Postgres, db)),
		ent2.Log(func(i ...any) {
			logger.Debug(
				"",
				map[string]any{
					"message": fmt.Sprint(i...),
				})
		}),
	)
	defer dbClient.Close()

	if cfg.App.Environment == "development" {
		dbClient = dbClient.Debug()
	}

	err = dbClient.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		logger.Fatal("failed to create schema resources", err, nil)
	}

	// watermill
	natsURL := fmt.Sprintf("%s:%d", cfg.NATS.Host, cfg.NATS.Port)

	exactlyOnceDelivery := true

	natsOptions := []nats.Option{
		nats.RetryOnFailedConnect(true),
		nats.Timeout(30 * time.Second),
		nats.ReconnectWait(1 * time.Second),
	}

	jetstreamMarshalerUnmarshaler := jetstream.GobMarshaler{}

	jetstreamOptions := make([]nats.JSOpt, 0)

	jetstreamPublisher, err := jetstream.NewPublisher(
		jetstream.PublisherConfig{
			URL:              natsURL,
			Marshaler:        jetstreamMarshalerUnmarshaler,
			NatsOptions:      natsOptions,
			JetstreamOptions: jetstreamOptions,
			AutoProvision:    true,
			TrackMsgId:       exactlyOnceDelivery,
		},
		watermillLogger,
	)
	if err != nil {
		logger.Fatal("failed to create publisher", err, nil)
	}
	defer jetstreamPublisher.Close()

	subscribeOptions := []nats.SubOpt{
		nats.DeliverNew(),
		nats.AckExplicit(),
	}
	jetstreamSubscriber, err := jetstream.NewSubscriber(
		jetstream.SubscriberConfig{
			URL:              natsURL,
			Unmarshaler:      jetstreamMarshalerUnmarshaler,
			QueueGroup:       "launch",
			DurableName:      "durable-name",
			SubscribersCount: 10,
			AckWaitTimeout:   30 * time.Second,
			NatsOptions:      natsOptions,
			SubscribeOptions: subscribeOptions,
			JetstreamOptions: jetstreamOptions,
			CloseTimeout:     30 * time.Second,
			AutoProvision:    true,
			AckSync:          exactlyOnceDelivery,
		},
		watermillLogger,
	)
	if err != nil {
		logger.Fatal("failed to create subscriber", err, nil)
	}
	defer jetstreamSubscriber.Close()

	router, err := message.NewRouter(
		message.RouterConfig{
			CloseTimeout: 30 * time.Second,
		},
		watermillLogger,
	)
	if err != nil {
		logger.Fatal("failed to create router", err, nil)
	}
	defer router.Close()

	router.AddMiddleware(watermillmiddleware.Recoverer)

	commandEventMarshaler := cqrs.JSONMarshaler{}
	generateCommandsTopic := func(commandName string) string {
		return strings.Replace(commandName, ".", "_", 1)
	}
	commandsSubscriberConstructor := func(handlerName string) (message.Subscriber, error) {
		return jetstreamSubscriber, nil
	}
	generateEventsTopic := func(eventName string) string {
		return strings.Replace(eventName, ".", "_", 1)
	}
	eventsSubscriberConstructor := func(_ string) (message.Subscriber, error) {
		return jetstream.NewSubscriber(
			jetstream.SubscriberConfig{
				URL:              natsURL,
				Unmarshaler:      jetstreamMarshalerUnmarshaler,
				DurableName:      "durable-name",
				AckWaitTimeout:   30 * time.Second,
				NatsOptions:      natsOptions,
				SubscribeOptions: subscribeOptions,
				JetstreamOptions: jetstreamOptions,
				CloseTimeout:     30 * time.Second,
				AutoProvision:    true,
				AckSync:          exactlyOnceDelivery,
			},
			watermillLogger,
		)
	}

	commandBus, err := cqrs.NewCommandBus(
		jetstreamPublisher,
		generateCommandsTopic,
		commandEventMarshaler,
	)
	if err != nil {
		logger.Fatal("failed to create commands bus", err, nil)
	}

	eventBus, err := cqrs.NewEventBus(
		jetstreamPublisher,
		generateEventsTopic,
		commandEventMarshaler,
	)
	if err != nil {
		logger.Fatal("failed to create events bus", err, nil)
	}

	sanitizer := bluemonday.UGCPolicy()

	kindsRepository := adapters2.NewKindRepository(dbClient)
	ordersRepository := adapters2.NewOrderRepository(dbClient)
	// launchesRepository := adapters.NewLaunchRepository(dbClient)

	commandHandlers := []cqrs.CommandHandler{
		commands2.NewCreateKindHandler(eventBus, kindsRepository, sanitizer, logger, metricsClient),
		commands2.NewDeleteKindHandler(eventBus, kindsRepository, logger, metricsClient),
		commands2.NewChangeKindNameHandler(eventBus, kindsRepository, sanitizer, logger, metricsClient),
		commands2.NewChangeKindDescriptionHandler(eventBus, kindsRepository, sanitizer, logger, metricsClient),
		commands2.NewMakeKindPublishedHandler(eventBus, kindsRepository, logger, metricsClient),
		commands2.NewMakeKindDraftHandler(eventBus, kindsRepository, logger, metricsClient),

		commands2.NewCreateOrderHandler(eventBus, ordersRepository, sanitizer, logger, metricsClient),
		commands2.NewDeleteOrderHandler(eventBus, ordersRepository, logger, metricsClient),
		commands2.NewChangeOrderNameHandler(eventBus, ordersRepository, sanitizer, logger, metricsClient),
		commands2.NewChangeOrderDescriptionHandler(eventBus, ordersRepository, sanitizer, logger, metricsClient),
		commands2.NewMakeOrderPublishedHandler(eventBus, ordersRepository, logger, metricsClient),
		commands2.NewMakeOrderDraftHandler(eventBus, ordersRepository, logger, metricsClient),
	}

	commandProcessor, err := cqrs.NewCommandProcessor(
		commandHandlers,
		generateCommandsTopic,
		commandsSubscriberConstructor,
		commandEventMarshaler,
		watermillLogger,
	)
	if err != nil {
		logger.Fatal("failed to create commands processor", err, nil)
	}

	if err := commandProcessor.AddHandlersToRouter(router); err != nil {
		logger.Fatal("failed to add command handlers to router", err, nil)
	}

	eventHandlers := []cqrs.EventHandler{
		events2.KindCreatedHandler{},
		events2.KindDeletedHandler{},
		events2.KindNameChangedHandler{},
		events2.KindDescriptionChangedHandler{},
		events2.KindMadePublishedHandler{},
		events2.KindMadeDraftHandler{},

		events2.OrderCreatedHandler{},
		events2.OrderDeletedHandler{},
		events2.OrderNameChangedHandler{},
		events2.OrderDescriptionChangedHandler{},
		events2.OrderMadePublishedHandler{},
		events2.OrderMadeDraftHandler{},
	}

	eventProcessor, err := cqrs.NewEventProcessor(
		eventHandlers,
		generateEventsTopic,
		eventsSubscriberConstructor,
		commandEventMarshaler,
		watermillLogger,
	)
	if err != nil {
		logger.Fatal("failed to create event processor", err, nil)
	}

	if err := eventProcessor.AddHandlersToRouter(router); err != nil {
		logger.Fatal("failed to add event handlers to router", err, nil)
	}

	app := app.Application{
		CommandBus: commandBus,
		Queries: app.Queries{
			ListKinds:     queries2.NewListKindsHandler(kindsRepository, logger, metricsClient),
			GetKind:       queries2.NewGetKindHandler(kindsRepository, logger, metricsClient),
			GetKindByName: queries2.NewGetKindByNameHandler(kindsRepository, logger, metricsClient),

			ListOrders:     queries2.NewListOrdersHandler(ordersRepository, logger, metricsClient),
			GetOrder:       queries2.NewGetOrderHandler(ordersRepository, logger, metricsClient),
			GetOrderByName: queries2.NewGetOrderByNameHandler(ordersRepository, logger, metricsClient),
		},
	}

	go func(ctx context.Context) {
		err := router.Run(ctx)
		if err != nil {
			logger.Fatal("failed to run router", err, nil)
		}
	}(ctx)
	<-router.Running()
	// watermill

	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(httpmw.Identity())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info(
				"request",
				map[string]any{
					"remote_ip":     v.RemoteIP,
					"host":          v.Host,
					"uri":           v.URI,
					"protocol":      v.Protocol,
					"user_agent":    v.UserAgent,
					"method":        v.Method,
					"request_id":    v.RequestID,
					"latency":       v.Latency,
					"latency_human": v.Latency.String(),
					"status":        v.Status,
					"bytes_out":     v.ResponseSize,
				})

			return nil
		},
		LogLatency:      true,
		LogProtocol:     true,
		LogRemoteIP:     true,
		LogHost:         true,
		LogMethod:       true,
		LogURI:          true,
		LogRequestID:    true,
		LogUserAgent:    true,
		LogStatus:       true,
		LogResponseSize: true,
	}))

	apiRouter := e.Group("/v1")
	// apiRouter.Use(mw...)
	// _ = mw

	server := ports2.NewHTTPServer(app)
	ports2.RegisterHandlers(apiRouter, server)

	e.Server.ReadTimeout = cfg.HTTP.ReadTimeout
	e.Server.WriteTimeout = cfg.HTTP.WriteTimeout
	e.Server.IdleTimeout = cfg.HTTP.IdleTimeout

	a := application.New()
	a.AddAdapters(
		httputils.NewEchoAdapter(
			e,
			fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port),
			logger,
		),
	)

	if cfg.App.Environment == "development" {
		// a.AddAdapters(
		// 	application.NewDebugAdapter(
		// 		fmt.Sprintf("%s:%d", cfg.Debug.Host, cfg.Debug.Port),
		// 	),
		// )
	}

	a.WithShutdownTimeout(cfg.App.ShutdownTimeout)
	a.Run(ctx)
}

package commands

import (
	"context"

	"github.com/google/uuid"

	"github.com/edanko/launches-api/internal/app/events"
	order2 "github.com/edanko/launches-api/internal/domain/order"
	"github.com/edanko/launches-api/pkg/decorator"
	"github.com/edanko/launches-api/pkg/errors"
	"github.com/edanko/launches-api/pkg/logs"
)

type MakeOrderPublished struct {
	ID uuid.UUID `json:"id"`
}

type makeOrderPublishedHandler struct {
	eventBus   eventBus
	repository order2.Repository
}

type MakeOrderPublishedHandler decorator.CommandHandler

func NewMakeOrderPublishedHandler(
	eventBus eventBus,
	repo order2.Repository,
	logger logs.Logger,
	metricsClient decorator.MetricsClient,
) MakeOrderPublishedHandler {
	if repo == nil {
		panic("nil orderRepo")
	}

	return decorator.ApplyCommandDecorators(
		makeOrderPublishedHandler{
			eventBus:   eventBus,
			repository: repo,
		},
		// logger,
		metricsClient,
	)
}

func (h makeOrderPublishedHandler) HandlerName() string {
	return "make-order-published"
}

func (h makeOrderPublishedHandler) NewCommand() any {
	return &MakeOrderPublished{}
}

func (h makeOrderPublishedHandler) Handle(ctx context.Context, command any) error {
	c := command.(*MakeOrderPublished)

	err := h.repository.Update(
		ctx,
		c.ID,
		func(k *order2.Order) (*order2.Order, error) {
			err := k.MakePublished()
			if err != nil {
				return k, err
			}
			return k, nil
		},
	)
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-update-order")
	}

	err = h.eventBus.Publish(ctx, events.OrderMadePublished{
		ID: c.ID,
	})
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-publish-event")
	}

	return nil
}

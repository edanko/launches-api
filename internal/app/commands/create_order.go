package commands

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/edanko/launches-api/internal/app/events"
	order2 "github.com/edanko/launches-api/internal/domain/order"
	"github.com/edanko/launches-api/pkg/decorator"
	"github.com/edanko/launches-api/pkg/errors"
	"github.com/edanko/launches-api/pkg/logs"
	"github.com/edanko/launches-api/pkg/sanitizer"
)

type CreateOrder struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Status      string    `json:"status"`
}

type CreateOrderHandler decorator.CommandHandler

type createOrderHandler struct {
	eventBus   eventBus
	repository order2.Repository
	sanitizer  sanitizer.Sanitizer
}

func NewCreateOrderHandler(
	eventBus eventBus,
	repo order2.Repository,
	s sanitizer.Sanitizer,
	logger logs.Logger,
	metricsClient decorator.MetricsClient,
) CreateOrderHandler {
	if repo == nil {
		panic("nil orderRepo")
	}

	return decorator.ApplyCommandDecorators(
		createOrderHandler{
			eventBus:   eventBus,
			repository: repo,
			sanitizer:  s,
		},
		// logger,
		metricsClient,
	)
}

func (h createOrderHandler) HandlerName() string {
	return "create-order"
}

func (h createOrderHandler) NewCommand() any {
	return &CreateOrder{}
}

func (h createOrderHandler) Handle(ctx context.Context, command any) error {
	c := command.(*CreateOrder)

	c.Name = h.sanitizer.Sanitize(c.Name)
	c.Name = strings.TrimSpace(c.Name)

	if c.Description != nil {
		description := h.sanitizer.Sanitize(*c.Description)
		description = strings.TrimSpace(description)
		c.Description = &description
	}

	k, err := order2.NewOrder(
		c.ID,
		c.CreatedAt,
		c.UpdatedAt,
		c.Name,
		c.Description,
		c.Status,
	)
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-create-order")
	}

	exist, err := h.repository.Exist(ctx, c.Name)
	if err != nil {
		// h.logger.Warn().Err(err).Msg("error checking if order with the name already exists")
		return nil
	}
	if exist {
		// h.logger.Warn().Err(order.ErrOrderAlreadyExist).Msg("order with the name already exists")
		return nil
	}

	err = h.repository.Create(ctx, k)
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-create-order")
	}

	err = h.eventBus.Publish(ctx, events.OrderCreated{
		ID:          k.ID(),
		Name:        k.Name(),
		Description: k.Description(),
		Status:      k.Status().String(),
	})
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-publish-event")
	}

	return nil
}

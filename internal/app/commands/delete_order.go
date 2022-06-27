package commands

import (
	"context"

	"github.com/google/uuid"

	"github.com/edanko/launches-api/internal/app/events"
	"github.com/edanko/launches-api/internal/domain/order"
	"github.com/edanko/launches-api/pkg/decorator"
	"github.com/edanko/launches-api/pkg/errors"
	"github.com/edanko/launches-api/pkg/logs"
)

type DeleteOrder struct {
	ID uuid.UUID `json:"id"`
}

type DeleteOrderHandler decorator.CommandHandler

type deleteOrderHandler struct {
	eventBus   eventBus
	repository order.Repository
}

func NewDeleteOrderHandler(
	eventBus eventBus,
	repo order.Repository,
	logger logs.Logger,
	metricsClient decorator.MetricsClient,
) CreateOrderHandler {
	if repo == nil {
		panic("nil orderRepo")
	}

	return decorator.ApplyCommandDecorators(
		deleteOrderHandler{
			eventBus:   eventBus,
			repository: repo,
		},
		// logger,
		metricsClient,
	)
}

func (h deleteOrderHandler) HandlerName() string {
	return "delete-order"
}

func (h deleteOrderHandler) NewCommand() any {
	return &DeleteOrder{}
}

func (h deleteOrderHandler) Handle(ctx context.Context, command any) error {
	c := command.(*DeleteOrder)

	err := h.repository.Delete(ctx, c.ID)
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-delete-order")
	}

	err = h.eventBus.Publish(ctx, events.OrderDeleted{
		ID: c.ID,
	})
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-publish-event")
	}

	return nil
}

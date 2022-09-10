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

type MakeOrderDraft struct {
	ID uuid.UUID `json:"id"`
}

type makeOrderDraftHandler struct {
	eventBus   eventBus
	repository order2.Repository
}

type MakeOrderDraftHandler decorator.CommandHandler

func NewMakeOrderDraftHandler(
	eventBus eventBus,
	repo order2.Repository,
	logger logs.Logger,
	metricsClient decorator.MetricsClient,
) MakeOrderDraftHandler {
	if repo == nil {
		panic("nil orderRepo")
	}

	return decorator.ApplyCommandDecorators(
		makeOrderDraftHandler{
			eventBus:   eventBus,
			repository: repo,
		},
		// logger,
		metricsClient,
	)
}

func (h makeOrderDraftHandler) HandlerName() string {
	return "make-order-draft"
}

func (h makeOrderDraftHandler) NewCommand() any {
	return &MakeOrderDraft{}
}

func (h makeOrderDraftHandler) Handle(ctx context.Context, command any) error {
	c := command.(*MakeOrderDraft)

	err := h.repository.Update(
		ctx,
		c.ID,
		func(k *order2.Order) (*order2.Order, error) {
			err := k.MakeDraft()
			if err != nil {
				return k, err
			}
			return k, nil
		},
	)
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-update-order")
	}

	err = h.eventBus.Publish(ctx, events.OrderMadeDraft{
		ID: c.ID,
	})
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-publish-event")
	}

	return nil
}

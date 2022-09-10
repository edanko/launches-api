package commands

import (
	"context"
	"strings"

	"github.com/google/uuid"

	"github.com/edanko/launches-api/internal/app/events"
	order2 "github.com/edanko/launches-api/internal/domain/order"
	"github.com/edanko/launches-api/pkg/decorator"
	"github.com/edanko/launches-api/pkg/errors"
	"github.com/edanko/launches-api/pkg/logs"
	"github.com/edanko/launches-api/pkg/sanitizer"
)

type ChangeOrderDescription struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
}

type ChangeOrderDescriptionHandler decorator.CommandHandler

type changeOrderDescriptionHandler struct {
	eventBus   eventBus
	repository order2.Repository
	sanitizer  sanitizer.Sanitizer
}

func NewChangeOrderDescriptionHandler(
	eventBus eventBus,
	repo order2.Repository,
	s sanitizer.Sanitizer,
	logger logs.Logger,
	metricsClient decorator.MetricsClient,
) ChangeOrderDescriptionHandler {
	if repo == nil {
		panic("nil orderRepo")
	}

	return decorator.ApplyCommandDecorators(
		changeOrderDescriptionHandler{
			eventBus:   eventBus,
			repository: repo,
			sanitizer:  s,
		},
		// logger,
		metricsClient,
	)
}

func (h changeOrderDescriptionHandler) HandlerName() string {
	return "change-order-description"
}

func (h changeOrderDescriptionHandler) NewCommand() any {
	return &ChangeOrderDescription{}
}

func (h changeOrderDescriptionHandler) Handle(ctx context.Context, command any) error {
	c := command.(*ChangeOrderDescription)

	c.Description = h.sanitizer.Sanitize(c.Description)
	c.Description = strings.TrimSpace(c.Description)

	err := h.repository.Update(
		ctx,
		c.ID,
		func(o *order2.Order) (*order2.Order, error) {
			o.ChangeDescription(c.Description)
			return o, nil
		})
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-update-order")
	}

	err = h.eventBus.Publish(ctx, events.OrderDescriptionChanged{
		ID:          c.ID,
		Description: c.Description,
	})
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-publish-event")
	}

	return nil
}

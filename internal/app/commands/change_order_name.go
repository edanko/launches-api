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

type ChangeOrderName struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type ChangeOrderNameHandler decorator.CommandHandler

type changeOrderNameHandler struct {
	eventBus   eventBus
	repository order2.Repository
	sanitizer  sanitizer.Sanitizer
}

func NewChangeOrderNameHandler(
	eventBus eventBus,
	repo order2.Repository,
	s sanitizer.Sanitizer,
	logger logs.Logger,
	metricsClient decorator.MetricsClient,
) ChangeOrderNameHandler {
	if repo == nil {
		panic("nil orderRepo")
	}

	return decorator.ApplyCommandDecorators(
		changeOrderNameHandler{
			eventBus:   eventBus,
			repository: repo,
			sanitizer:  s,
		},
		// logger,
		metricsClient,
	)
}

func (h changeOrderNameHandler) HandlerName() string {
	return "change-order-name"
}

func (h changeOrderNameHandler) NewCommand() any {
	return &ChangeOrderName{}
}

func (h changeOrderNameHandler) Handle(ctx context.Context, command any) error {
	c := command.(*ChangeOrderName)

	c.Name = h.sanitizer.Sanitize(c.Name)
	c.Name = strings.TrimSpace(c.Name)
	c.Name = strings.ToUpper(c.Name)

	err := h.repository.Update(
		ctx,
		c.ID,
		func(o *order2.Order) (*order2.Order, error) {
			o.ChangeName(c.Name)
			return o, nil
		},
	)
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-update-order")
	}

	err = h.eventBus.Publish(ctx, events.OrderNameChanged{
		ID:   c.ID,
		Name: c.Name,
	})
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-publish-event")
	}

	return nil
}

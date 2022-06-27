package commands

import (
	"context"
	"strings"

	"github.com/google/uuid"

	"github.com/edanko/launches-api/internal/app/events"
	kind2 "github.com/edanko/launches-api/internal/domain/kind"
	"github.com/edanko/launches-api/pkg/decorator"
	"github.com/edanko/launches-api/pkg/errors"
	"github.com/edanko/launches-api/pkg/logs"
	"github.com/edanko/launches-api/pkg/sanitizer"
)

type ChangeKindName struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type ChangeKindNameHandler decorator.CommandHandler

type changeKindNameHandler struct {
	eventBus   eventBus
	repository kind2.Repository
	sanitizer  sanitizer.Sanitizer
}

func NewChangeKindNameHandler(
	eventBus eventBus,
	repo kind2.Repository,
	s sanitizer.Sanitizer,
	logger logs.Logger,
	metricsClient decorator.MetricsClient,
) ChangeKindNameHandler {
	if repo == nil {
		panic("nil kindRepo")
	}

	return decorator.ApplyCommandDecorators(
		changeKindNameHandler{
			eventBus:   eventBus,
			repository: repo,
			sanitizer:  s,
		},
		// logger,
		metricsClient,
	)
}

func (h changeKindNameHandler) HandlerName() string {
	return "change-kind-name"
}

func (h changeKindNameHandler) NewCommand() any {
	return &ChangeKindName{}
}

func (h changeKindNameHandler) Handle(ctx context.Context, command any) error {
	c := command.(*ChangeKindName)

	c.Name = h.sanitizer.Sanitize(c.Name)
	c.Name = strings.TrimSpace(c.Name)
	c.Name = strings.ToUpper(c.Name)

	err := h.repository.Update(
		ctx,
		c.ID,
		func(k *kind2.Kind) (*kind2.Kind, error) {
			k.ChangeName(c.Name)
			return k, nil

		},
	)
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-update-kind")
	}

	err = h.eventBus.Publish(ctx, events.KindNameChanged{
		ID:   c.ID,
		Name: c.Name,
	})
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-publish-event")
	}

	return nil
}

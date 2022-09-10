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

type ChangeKindDescription struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
}

type ChangeKindDescriptionHandler decorator.CommandHandler

type changeKindDescriptionHandler struct {
	eventBus   eventBus
	repository kind2.Repository
	sanitizer  sanitizer.Sanitizer
}

func NewChangeKindDescriptionHandler(
	eventBus eventBus,
	repo kind2.Repository,
	s sanitizer.Sanitizer,
	logger logs.Logger,
	metricsClient decorator.MetricsClient,
) ChangeKindDescriptionHandler {
	if repo == nil {
		panic("nil kindRepo")
	}

	return decorator.ApplyCommandDecorators(
		changeKindDescriptionHandler{
			eventBus:   eventBus,
			repository: repo,
			sanitizer:  s,
		},
		// logger,
		metricsClient,
	)
}

func (h changeKindDescriptionHandler) HandlerName() string {
	return "change-kind-description"
}

func (h changeKindDescriptionHandler) NewCommand() any {
	return &ChangeKindDescription{}
}

func (h changeKindDescriptionHandler) Handle(ctx context.Context, command any) error {
	c := command.(*ChangeKindDescription)

	c.Description = h.sanitizer.Sanitize(c.Description)
	c.Description = strings.TrimSpace(c.Description)

	err := h.repository.Update(
		ctx,
		c.ID,
		func(k *kind2.Kind) (*kind2.Kind, error) {
			k.ChangeDescription(c.Description)
			return k, nil
		},
	)
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-update-kind")
	}

	err = h.eventBus.Publish(ctx, events.KindDescriptionChanged{
		ID:          c.ID,
		Description: c.Description,
	})
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-publish-event")
	}

	return nil
}

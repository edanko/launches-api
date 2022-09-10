package commands

import (
	"context"

	"github.com/google/uuid"

	"github.com/edanko/launches-api/internal/app/events"
	kind2 "github.com/edanko/launches-api/internal/domain/kind"
	"github.com/edanko/launches-api/pkg/decorator"
	"github.com/edanko/launches-api/pkg/errors"
	"github.com/edanko/launches-api/pkg/logs"
)

type MakeKindPublished struct {
	ID uuid.UUID `json:"id"`
}

type makeKindPublishedHandler struct {
	eventBus   eventBus
	repository kind2.Repository
}

type MakeKindPublishedHandler decorator.CommandHandler

func NewMakeKindPublishedHandler(
	eventBus eventBus,
	repo kind2.Repository,
	logger logs.Logger,
	metricsClient decorator.MetricsClient,
) MakeKindPublishedHandler {
	if repo == nil {
		panic("nil kindRepo")
	}

	return decorator.ApplyCommandDecorators(
		makeKindPublishedHandler{
			eventBus:   eventBus,
			repository: repo,
		},
		// logger,
		metricsClient,
	)
}

func (h makeKindPublishedHandler) HandlerName() string {
	return "make-kind-published"
}

func (h makeKindPublishedHandler) NewCommand() any {
	return &MakeKindPublished{}
}

func (h makeKindPublishedHandler) Handle(ctx context.Context, command any) error {
	c := command.(*MakeKindPublished)

	err := h.repository.Update(
		ctx,
		c.ID,
		func(k *kind2.Kind) (*kind2.Kind, error) {
			err := k.MakePublished()
			if err != nil {
				return k, err
			}
			return k, nil
		},
	)
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-update-kind")
	}

	err = h.eventBus.Publish(ctx, events.KindMadePublished{
		ID: c.ID,
	})
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-publish-event")
	}

	return nil
}

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

type MakeKindDraft struct {
	ID uuid.UUID `json:"id"`
}

type makeKindDraftHandler struct {
	eventBus   eventBus
	repository kind2.Repository
}

type MakeKindDraftHandler decorator.CommandHandler

func NewMakeKindDraftHandler(
	eventBus eventBus,
	repo kind2.Repository,
	logger logs.Logger,
	metricsClient decorator.MetricsClient,
) MakeKindDraftHandler {
	if repo == nil {
		panic("nil kindRepo")
	}

	return decorator.ApplyCommandDecorators(
		makeKindDraftHandler{
			eventBus:   eventBus,
			repository: repo,
		},
		// logger,
		metricsClient,
	)
}

func (h makeKindDraftHandler) HandlerName() string {
	return "make-kind-draft"
}

func (h makeKindDraftHandler) NewCommand() any {
	return &MakeKindDraft{}
}

func (h makeKindDraftHandler) Handle(ctx context.Context, command any) error {
	c := command.(*MakeKindDraft)

	err := h.repository.Update(
		ctx,
		c.ID,
		func(k *kind2.Kind) (*kind2.Kind, error) {
			err := k.MakeDraft()
			if err != nil {
				return k, err
			}
			return k, nil
		},
	)
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-update-kind")
	}

	err = h.eventBus.Publish(ctx, events.KindMadeDraft{
		ID: c.ID,
	})
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-publish-event")
	}

	return nil
}

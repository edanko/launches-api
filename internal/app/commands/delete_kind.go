package commands

import (
	"context"

	"github.com/google/uuid"

	"github.com/edanko/launches-api/internal/app/events"
	"github.com/edanko/launches-api/internal/domain/kind"
	"github.com/edanko/launches-api/pkg/decorator"
	"github.com/edanko/launches-api/pkg/errors"
	"github.com/edanko/launches-api/pkg/logs"
)

type DeleteKind struct {
	ID uuid.UUID `json:"id"`
}

type DeleteKindHandler decorator.CommandHandler

type deleteKindHandler struct {
	eventBus   eventBus
	repository kind.Repository
}

func NewDeleteKindHandler(
	eventBus eventBus,
	repo kind.Repository,
	logger logs.Logger,
	metricsClient decorator.MetricsClient,
) CreateKindHandler {
	if repo == nil {
		panic("nil kindRepo")
	}

	return decorator.ApplyCommandDecorators(
		deleteKindHandler{
			eventBus:   eventBus,
			repository: repo,
		},
		// logger,
		metricsClient,
	)
}

func (h deleteKindHandler) HandlerName() string {
	return "delete-kind"
}

func (h deleteKindHandler) NewCommand() any {
	return &DeleteKind{}
}

func (h deleteKindHandler) Handle(ctx context.Context, command any) error {
	c := command.(*DeleteKind)

	err := h.repository.Delete(ctx, c.ID)
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-delete-kind")
	}

	err = h.eventBus.Publish(ctx, events.KindDeleted{
		ID: c.ID,
	})
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-publish-event")
	}

	return nil
}

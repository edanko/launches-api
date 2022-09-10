package commands

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/edanko/launches-api/internal/app/events"
	kind2 "github.com/edanko/launches-api/internal/domain/kind"
	"github.com/edanko/launches-api/pkg/decorator"
	"github.com/edanko/launches-api/pkg/errors"
	"github.com/edanko/launches-api/pkg/logs"
	"github.com/edanko/launches-api/pkg/sanitizer"
)

type CreateKind struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Status      string    `json:"status"`
}

type CreateKindHandler decorator.CommandHandler

type createKindHandler struct {
	eventBus   eventBus
	repository kind2.Repository
	sanitizer  sanitizer.Sanitizer
	logger     logs.Logger
}

func NewCreateKindHandler(
	eventBus eventBus,
	repo kind2.Repository,
	s sanitizer.Sanitizer,
	logger logs.Logger,
	metricsClient decorator.MetricsClient,
) CreateKindHandler {
	if repo == nil {
		panic("nil kindRepo")
	}

	return decorator.ApplyCommandDecorators(
		createKindHandler{
			eventBus:   eventBus,
			repository: repo,
			sanitizer:  s,
			logger:     logger,
		},
		metricsClient,
	)
}

func (h createKindHandler) HandlerName() string {
	return "create-kind"
}

func (h createKindHandler) NewCommand() any {
	return &CreateKind{}
}

func (h createKindHandler) Handle(ctx context.Context, command any) error {
	c := command.(*CreateKind)

	c.Name = h.sanitizer.Sanitize(c.Name)
	c.Name = strings.TrimSpace(c.Name)

	if c.Description != nil {
		description := h.sanitizer.Sanitize(*c.Description)
		description = strings.TrimSpace(description)
		c.Description = &description
	}

	k, err := kind2.NewKind(
		c.ID,
		c.CreatedAt,
		c.UpdatedAt,
		c.Name,
		c.Description,
		c.Status,
	)
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-create-kind")
	}

	exist, err := h.repository.Exist(ctx, c.Name)
	if err != nil {
		// h.logger.Warn().Err(err).Msg("error checking if kind with the name already exists")
		return nil
	}
	if exist {
		// h.logger.Warn().Err(kind.ErrKindAlreadyExist).Msg("kind with the name already exists")
		return nil
	}

	err = h.repository.Create(ctx, k)
	if err != nil {
		return errors.NewSlugError(err.Error(), "unable-to-create-kind")
	}

	err = h.eventBus.Publish(ctx, events.KindCreated{
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

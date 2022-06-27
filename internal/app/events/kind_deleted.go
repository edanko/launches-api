package events

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type KindDeleted struct {
	ID uuid.UUID `json:"id"`
}

type KindDeletedHandler struct {
	// 	Repository persistence.KindRepository
	// CommandBus *cqrs.CommandBus
}

func (k KindDeletedHandler) HandlerName() string {
	return "kind-deleted"
}

func (k KindDeletedHandler) NewEvent() any {
	return &KindDeleted{}
}

func (k KindDeletedHandler) Handle(_ context.Context, event any) error {
	// ctx, cancel := context.WithTimeout(ctx, time.Second*120)
	// defer cancel()

	e := event.(*KindDeleted)

	log.Info().Msg("KindDeleted events received " + e.ID.String())

	return nil
}

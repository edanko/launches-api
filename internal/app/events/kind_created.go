package events

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type KindCreated struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Status      string    `json:"status"`
}

type KindCreatedHandler struct {
	// 	Repository persistence.KindRepository
	// CommandBus *cqrs.CommandBus
}

func (k KindCreatedHandler) HandlerName() string {
	return "kind-created"
}

func (k KindCreatedHandler) NewEvent() any {
	return &KindCreated{}
}

func (k KindCreatedHandler) Handle(_ context.Context, _ any) error {
	// ctx, cancel := context.WithTimeout(ctx, time.Second*120)
	// defer cancel()

	// e := events.(*KindCreated)
	// _ = e

	log.Info().Msg("KindCreated events received")

	return nil
}

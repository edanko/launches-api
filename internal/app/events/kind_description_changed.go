package events

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type KindDescriptionChanged struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
}

type KindDescriptionChangedHandler struct {
	// 	Repository persistence.KindRepository
	// CommandBus *cqrs.CommandBus
}

func (k KindDescriptionChangedHandler) HandlerName() string {
	return "kind-description-changed"
}

func (k KindDescriptionChangedHandler) NewEvent() any {
	return &KindDescriptionChanged{}
}

func (k KindDescriptionChangedHandler) Handle(_ context.Context, event any) error {
	// ctx, cancel := context.WithTimeout(ctx, time.Second*120)
	// defer cancel()

	e := event.(*KindDescriptionChanged)

	log.Info().Msg("KindDescriptionChanged events received " + e.ID.String())

	return nil
}

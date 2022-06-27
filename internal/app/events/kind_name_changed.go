package events

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type KindNameChanged struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type KindNameChangedHandler struct {
	// 	Repository persistence.KindRepository
	// CommandBus *cqrs.CommandBus
}

func (k KindNameChangedHandler) HandlerName() string {
	return "kind-name-changed"
}

func (k KindNameChangedHandler) NewEvent() any {
	return &KindNameChanged{}
}

func (k KindNameChangedHandler) Handle(_ context.Context, event any) error {
	// ctx, cancel := context.WithTimeout(ctx, time.Second*120)
	// defer cancel()

	e := event.(*KindNameChanged)

	log.Info().Msg("KindNameChanged events received " + e.ID.String())

	return nil
}

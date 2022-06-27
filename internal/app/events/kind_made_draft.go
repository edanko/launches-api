package events

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type KindMadeDraft struct {
	ID uuid.UUID `json:"id"`
}

type KindMadeDraftHandler struct {
	// 	Repository persistence.KindRepository
	// CommandBus *cqrs.CommandBus
}

func (k KindMadeDraftHandler) HandlerName() string {
	return "kind-made-draft"
}

func (k KindMadeDraftHandler) NewEvent() any {
	return &KindMadeDraft{}
}

func (k KindMadeDraftHandler) Handle(_ context.Context, event any) error {
	// ctx, cancel := context.WithTimeout(ctx, time.Second*120)
	// defer cancel()

	e := event.(*KindMadeDraft)

	log.Info().Msg("KindMadeDraft events received " + e.ID.String())

	return nil
}

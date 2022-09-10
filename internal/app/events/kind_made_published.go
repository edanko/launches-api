package events

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type KindMadePublished struct {
	ID uuid.UUID `json:"id"`
}

type KindMadePublishedHandler struct {
	// 	Repository persistence.KindRepository
	// CommandBus *cqrs.CommandBus
}

func (k KindMadePublishedHandler) HandlerName() string {
	return "kind-made-published"
}

func (k KindMadePublishedHandler) NewEvent() any {
	return &KindMadePublished{}
}

func (k KindMadePublishedHandler) Handle(_ context.Context, event any) error {
	// ctx, cancel := context.WithTimeout(ctx, time.Second*120)
	// defer cancel()

	e := event.(*KindMadePublished)

	log.Info().Msg("KindMadePublished events received " + e.ID.String())

	return nil
}

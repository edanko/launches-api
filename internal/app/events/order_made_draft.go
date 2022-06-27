package events

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type OrderMadeDraft struct {
	ID uuid.UUID `json:"id"`
}

type OrderMadeDraftHandler struct {
	// 	Repository persistence.OrderRepository
	// CommandBus *cqrs.CommandBus
}

func (k OrderMadeDraftHandler) HandlerName() string {
	return "order-made-draft"
}

func (k OrderMadeDraftHandler) NewEvent() any {
	return &OrderMadeDraft{}
}

func (k OrderMadeDraftHandler) Handle(_ context.Context, event any) error {
	// ctx, cancel := context.WithTimeout(ctx, time.Second*120)
	// defer cancel()

	e := event.(*OrderMadeDraft)

	log.Info().Msg("OrderMadeDraft events received " + e.ID.String())

	return nil
}

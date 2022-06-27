package events

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type OrderDescriptionChanged struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
}

type OrderDescriptionChangedHandler struct {
	// 	Repository persistence.OrderRepository
	// CommandBus *cqrs.CommandBus
}

func (k OrderDescriptionChangedHandler) HandlerName() string {
	return "order-description-changed"
}

func (k OrderDescriptionChangedHandler) NewEvent() any {
	return &OrderDescriptionChanged{}
}

func (k OrderDescriptionChangedHandler) Handle(_ context.Context, event any) error {
	// ctx, cancel := context.WithTimeout(ctx, time.Second*120)
	// defer cancel()

	e := event.(*OrderDescriptionChanged)

	log.Info().Msg("OrderDescriptionChanged events received " + e.ID.String())

	return nil
}

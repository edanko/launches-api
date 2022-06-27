package events

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type OrderNameChanged struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type OrderNameChangedHandler struct {
	// 	Repository persistence.OrderRepository
	// CommandBus *cqrs.CommandBus
}

func (k OrderNameChangedHandler) HandlerName() string {
	return "order-name-changed"
}

func (k OrderNameChangedHandler) NewEvent() any {
	return &OrderNameChanged{}
}

func (k OrderNameChangedHandler) Handle(_ context.Context, event any) error {
	// ctx, cancel := context.WithTimeout(ctx, time.Second*120)
	// defer cancel()

	e := event.(*OrderNameChanged)

	log.Info().Msg("OrderNameChanged events received " + e.ID.String())

	return nil
}

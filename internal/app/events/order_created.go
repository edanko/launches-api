package events

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type OrderCreated struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Status      string    `json:"status"`
}

type OrderCreatedHandler struct {
	// 	Repository persistence.OrderRepository
	// CommandBus *cqrs.CommandBus
}

func (k OrderCreatedHandler) HandlerName() string {
	return "order-created"
}

func (k OrderCreatedHandler) NewEvent() any {
	return &OrderCreated{}
}

func (k OrderCreatedHandler) Handle(_ context.Context, _ any) error {
	// ctx, cancel := context.WithTimeout(ctx, time.Second*120)
	// defer cancel()

	// e := events.(*OrderCreated)
	// _ = e

	log.Info().Msg("OrderCreated events received")

	return nil
}

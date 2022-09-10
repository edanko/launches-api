package events

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type OrderMadePublished struct {
	ID uuid.UUID `json:"id"`
}

type OrderMadePublishedHandler struct {
	// 	Repository persistence.OrderRepository
	// CommandBus *cqrs.CommandBus
}

func (k OrderMadePublishedHandler) HandlerName() string {
	return "order-made-published"
}

func (k OrderMadePublishedHandler) NewEvent() any {
	return &OrderMadePublished{}
}

func (k OrderMadePublishedHandler) Handle(_ context.Context, event any) error {
	// ctx, cancel := context.WithTimeout(ctx, time.Second*120)
	// defer cancel()

	e := event.(*OrderMadePublished)

	log.Info().Msg("OrderMadePublished events received " + e.ID.String())

	return nil
}

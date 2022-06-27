package events

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type OrderDeleted struct {
	ID uuid.UUID `json:"id"`
}

type OrderDeletedHandler struct {
	// 	Repository persistence.OrderRepository
	// CommandBus *cqrs.CommandBus
}

func (k OrderDeletedHandler) HandlerName() string {
	return "order-deleted"
}

func (k OrderDeletedHandler) NewEvent() any {
	return &OrderDeleted{}
}

func (k OrderDeletedHandler) Handle(_ context.Context, event any) error {
	// ctx, cancel := context.WithTimeout(ctx, time.Second*120)
	// defer cancel()

	e := event.(*OrderDeleted)

	log.Info().Msg("OrderDeleted events received " + e.ID.String())

	return nil
}

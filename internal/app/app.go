package app

import (
	"context"

	queries2 "github.com/edanko/launches-api/internal/app/queries"
)

type mediator interface {
	Send(ctx context.Context, cmd any) error
}

type Application struct {
	CommandBus mediator
	Queries    Queries
}

type Queries struct {
	ListKinds     queries2.ListKindsHandler
	GetKind       queries2.GetKindHandler
	GetKindByName queries2.GetKindByNameHandler

	ListOrders     queries2.ListOrdersHandler
	GetOrder       queries2.GetOrderHandler
	GetOrderByName queries2.GetOrderByNameHandler
}

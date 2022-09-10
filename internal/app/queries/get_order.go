package queries

import (
	"context"

	"github.com/google/uuid"

	"github.com/edanko/launches-api/internal/domain/order"
	"github.com/edanko/launches-api/pkg/decorator"
	"github.com/edanko/launches-api/pkg/logs"
)

type GetOrderRequest struct {
	ID uuid.UUID
}

type GetOrderHandler decorator.QueryHandler[GetOrderRequest, Order]

type GetOrderReadModel interface {
	Get(ctx context.Context, id uuid.UUID) (*order.Order, error)
}

type getOrderHandler struct {
	readModel GetOrderReadModel
}

func (h getOrderHandler) Handle(
	ctx context.Context,
	query GetOrderRequest,
) (Order, error) {
	k, err := h.readModel.Get(ctx, query.ID)
	if err != nil {
		return Order{}, err
	}
	ret := mapOrderFromDomain(k)

	return ret, nil
}

func NewGetOrderHandler(
	readModel GetOrderReadModel,
	logger logs.Logger,
	metricsClient decorator.MetricsClient,
) GetOrderHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return decorator.ApplyQueryDecorators[GetOrderRequest, Order](
		getOrderHandler{readModel: readModel},
		// logger,
		metricsClient,
	)
}

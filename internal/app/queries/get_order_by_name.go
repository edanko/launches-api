package queries

import (
	"context"

	"github.com/edanko/launches-api/internal/domain/order"
	"github.com/edanko/launches-api/pkg/decorator"
	"github.com/edanko/launches-api/pkg/logs"
)

type GetOrderByNameRequest struct {
	Name string
}

type GetOrderByNameHandler decorator.QueryHandler[GetOrderByNameRequest, Order]

type GetOrderByNameReadModel interface {
	GetByName(ctx context.Context, name string) (*order.Order, error)
}

type getOrderByNameHandler struct {
	readModel GetOrderByNameReadModel
}

func (h getOrderByNameHandler) Handle(
	ctx context.Context,
	query GetOrderByNameRequest,
) (Order, error) {
	k, err := h.readModel.GetByName(ctx, query.Name)
	if err != nil {
		return Order{}, err
	}
	ret := mapOrderFromDomain(k)

	return ret, nil
}

func NewGetOrderByNameHandler(
	readModel GetOrderByNameReadModel,
	logger logs.Logger,
	metricsClient decorator.MetricsClient,
) GetOrderByNameHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return decorator.ApplyQueryDecorators[GetOrderByNameRequest, Order](
		getOrderByNameHandler{readModel: readModel},
		// logger,
		metricsClient,
	)
}

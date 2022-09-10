package queries

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/edanko/launches-api/pkg/decorator"
	"github.com/edanko/launches-api/pkg/logs"
)

type ListOrdersRequest struct {
	Limit     *int
	CreatedAt *time.Time
	ID        *uuid.UUID
	Status    *string
}

type ListOrdersHandler decorator.QueryHandler[ListOrdersRequest, []Order]

type ListOrdersReadModel interface {
	List(
		ctx context.Context,
		limit *int,
		createdAt *time.Time,
		id *uuid.UUID,
		status *string,
	) ([]Order, error)
}

type listOrdersHandler struct {
	readModel ListOrdersReadModel
}

func (h listOrdersHandler) Handle(
	ctx context.Context,
	query ListOrdersRequest,
) ([]Order, error) {
	return h.readModel.List(
		ctx,
		query.Limit,
		query.CreatedAt,
		query.ID,
		query.Status,
	)
}

func NewListOrdersHandler(
	readModel ListOrdersReadModel,
	logger logs.Logger,
	metricsClient decorator.MetricsClient,
) ListOrdersHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return decorator.ApplyQueryDecorators[ListOrdersRequest, []Order](
		listOrdersHandler{readModel: readModel},
		// logger,
		metricsClient,
	)
}

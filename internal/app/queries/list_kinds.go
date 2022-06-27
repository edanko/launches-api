package queries

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/edanko/launches-api/pkg/decorator"
	"github.com/edanko/launches-api/pkg/logs"
)

type ListKindsRequest struct {
	Limit     *int
	CreatedAt *time.Time
	ID        *uuid.UUID
	Status    *string
}

type ListKindsHandler decorator.QueryHandler[ListKindsRequest, []Kind]

type ListKindsReadModel interface {
	List(
		ctx context.Context,
		limit *int,
		createdAt *time.Time,
		id *uuid.UUID,
		status *string,
	) ([]Kind, error)
}

type listKindsHandler struct {
	readModel ListKindsReadModel
}

func (h listKindsHandler) Handle(
	ctx context.Context,
	query ListKindsRequest,
) ([]Kind, error) {
	return h.readModel.List(
		ctx,
		query.Limit,
		query.CreatedAt,
		query.ID,
		query.Status,
	)
}

func NewListKindsHandler(
	readModel ListKindsReadModel,
	logger logs.Logger,
	metricsClient decorator.MetricsClient,
) ListKindsHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return decorator.ApplyQueryDecorators[ListKindsRequest, []Kind](
		listKindsHandler{readModel: readModel},
		metricsClient,
	)
}

package queries

import (
	"context"

	"github.com/google/uuid"

	"github.com/edanko/launches-api/internal/domain/kind"
	"github.com/edanko/launches-api/pkg/decorator"
	"github.com/edanko/launches-api/pkg/logs"
)

type GetKindRequest struct {
	ID uuid.UUID
}

type GetKindHandler decorator.QueryHandler[GetKindRequest, Kind]

type GetKindReadModel interface {
	Get(ctx context.Context, id uuid.UUID) (*kind.Kind, error)
}

type getKindHandler struct {
	readModel GetKindReadModel
}

func (h getKindHandler) Handle(
	ctx context.Context,
	query GetKindRequest,
) (Kind, error) {
	k, err := h.readModel.Get(ctx, query.ID)
	if err != nil {
		return Kind{}, err
	}
	ret := mapKindFromDomain(k)

	return ret, nil
}

func NewGetKindHandler(
	readModel GetKindReadModel,
	logger logs.Logger,
	metricsClient decorator.MetricsClient,
) GetKindHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return decorator.ApplyQueryDecorators[GetKindRequest, Kind](
		getKindHandler{readModel: readModel},
		// logger,
		metricsClient,
	)
}

package queries

import (
	"context"

	"github.com/edanko/launches-api/internal/domain/kind"
	"github.com/edanko/launches-api/pkg/decorator"
	"github.com/edanko/launches-api/pkg/logs"
)

type GetKindByNameRequest struct {
	Name string
}

type GetKindByNameHandler decorator.QueryHandler[GetKindByNameRequest, Kind]

type GetKindByNameReadModel interface {
	GetByName(ctx context.Context, name string) (*kind.Kind, error)
}

type getKindByNameHandler struct {
	readModel GetKindByNameReadModel
}

func (h getKindByNameHandler) Handle(
	ctx context.Context,
	query GetKindByNameRequest,
) (Kind, error) {
	k, err := h.readModel.GetByName(ctx, query.Name)
	if err != nil {
		return Kind{}, err
	}
	ret := mapKindFromDomain(k)

	return ret, nil
}

func NewGetKindByNameHandler(
	readModel GetKindByNameReadModel,
	logger logs.Logger,
	metricsClient decorator.MetricsClient,
) GetKindByNameHandler {
	if readModel == nil {
		panic("nil readModel")
	}

	return decorator.ApplyQueryDecorators[GetKindByNameRequest, Kind](
		getKindByNameHandler{readModel: readModel},
		// logger,
		metricsClient,
	)
}

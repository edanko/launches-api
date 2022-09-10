package queries

import (
	"time"

	"github.com/google/uuid"

	"github.com/edanko/launches-api/internal/domain/kind"
	"github.com/edanko/launches-api/internal/domain/launch"
	"github.com/edanko/launches-api/internal/domain/order"
)

type Kind struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description *string
	Status      string
}

func mapKindFromDomain(d *kind.Kind) Kind {
	return Kind{
		ID:          d.ID(),
		CreatedAt:   d.CreatedAt(),
		UpdatedAt:   d.UpdatedAt(),
		Name:        d.Name(),
		Description: d.Description(),
		Status:      d.Status().String(),
	}
}

type Order struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description *string
	Status      string
}

func mapOrderFromDomain(d *order.Order) Order {
	return Order{
		ID:          d.ID(),
		CreatedAt:   d.CreatedAt(),
		UpdatedAt:   d.UpdatedAt(),
		Name:        d.Name(),
		Description: d.Description(),
		Status:      d.Status().String(),
	}
}

type Launch struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description *string
	Status      string
}

func mapLaunchFromDomain(d *launch.Launch) Launch {
	return Launch{
		ID:          d.ID(),
		CreatedAt:   d.CreatedAt(),
		UpdatedAt:   d.UpdatedAt(),
		Name:        d.Name(),
		Description: d.Description(),
		Status:      d.Status().String(),
	}
}

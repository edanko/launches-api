package order

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, k *Order) error
	Get(ctx context.Context, id uuid.UUID) (*Order, error)
	GetByName(ctx context.Context, name string) (*Order, error)
	Exist(ctx context.Context, name string) (bool, error)
	Update(ctx context.Context, id uuid.UUID, updateFn func(k *Order) (*Order, error)) error
	Delete(ctx context.Context, id uuid.UUID) error
}

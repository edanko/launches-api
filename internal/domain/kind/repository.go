package kind

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, k *Kind) error
	Get(ctx context.Context, id uuid.UUID) (*Kind, error)
	GetByName(ctx context.Context, name string) (*Kind, error)
	Exist(ctx context.Context, name string) (bool, error)
	// List(ctx context.Context, limit *int, createdAt *time.Time, id *uuid.UUID, status *string) ([]queries.Kind, error)
	Update(ctx context.Context, id uuid.UUID, updateFn func(k *Kind) (*Kind, error)) error
	Delete(ctx context.Context, id uuid.UUID) error
}

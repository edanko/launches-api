package launch

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, k *Launch) error
	Get(ctx context.Context, id uuid.UUID) (*Launch, error)
	GetByName(ctx context.Context, name string) (*Launch, error)
	Exist(ctx context.Context, name string) (bool, error)
	Update(ctx context.Context, id uuid.UUID, updateFn func(k *Launch) (*Launch, error)) error
	Delete(ctx context.Context, id uuid.UUID) error
}

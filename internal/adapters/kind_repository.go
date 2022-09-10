package adapters

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	ent2 "github.com/edanko/launches-api/internal/adapters/ent"
	kind2 "github.com/edanko/launches-api/internal/adapters/ent/kind"
	"github.com/edanko/launches-api/internal/app/queries"
	"github.com/edanko/launches-api/internal/domain/kind"
	"github.com/edanko/launches-api/internal/domain/launch"
)

type KindRepository struct {
	client *ent2.Client
}

var _ kind.Repository = (*KindRepository)(nil)

func NewKindRepository(c *ent2.Client) *KindRepository {
	return &KindRepository{
		client: c,
	}
}

func (r *KindRepository) Create(
	ctx context.Context,
	k *kind.Kind,
) error {
	return r.client.Kind.
		Create().
		SetID(k.ID()).
		SetCreatedAt(k.CreatedAt()).
		SetUpdatedAt(k.UpdatedAt()).
		SetName(k.Name()).
		SetNillableDescription(k.Description()).
		SetStatus(kind2.Status(k.Status().String())).
		Exec(ctx)
}

func (r *KindRepository) Get(
	ctx context.Context,
	id uuid.UUID,
) (*kind.Kind, error) {
	e, err := r.client.Kind.
		Query().
		Where(kind2.IDEQ(id)).
		Only(ctx)
	if ent2.IsNotFound(err) {
		return nil, kind.ErrKindNotFound
	}
	if err != nil {
		return nil, errors.Wrap(err, "unable to get actual kind by id")
	}

	return r.unmarshalKind(e), nil
}

func (r *KindRepository) GetByName(
	ctx context.Context,
	name string,
) (*kind.Kind, error) {
	e, err := r.client.Kind.
		Query().
		Where(kind2.NameEQ(name)).
		Only(ctx)
	if ent2.IsNotFound(err) {
		return nil, kind.ErrKindNotFound
	}
	if err != nil {
		return nil, errors.Wrap(err, "unable to get actual kind by name")
	}

	return r.unmarshalKind(e), nil
}

func (r *KindRepository) ListLaunches(ctx context.Context,
	kindName string,
) ([]*launch.Launch, error) {
	kindQuery := r.client.Kind.Query().Where(kind2.NameEQ(kindName)).WithLaunches()

	all, err := kindQuery.QueryLaunches().All(ctx)
	if err != nil {
		return nil, err
	}

	_ = all

	return nil, nil
}

func (r *KindRepository) List(
	ctx context.Context,
	limit *int,
	createdAt *time.Time,
	id *uuid.UUID,
	status *string,
) ([]queries.Kind, error) {
	kindQuery := r.client.Kind.Query().
		Order(
			ent2.Desc(kind2.FieldCreatedAt),
			ent2.Desc(kind2.FieldID),
		)

	if status != nil {
		kindQuery = kindQuery.Where(
			kind2.StatusEQ(kind2.Status(*status)),
		)
	}
	if limit != nil {
		kindQuery = kindQuery.Limit(*limit)
	}
	switch {
	case createdAt != nil && id != nil:
		kindQuery.Where(
			func(s *sql.Selector) {
				s.Where(
					sql.CompositeLT([]string{"created_at", "id"}, *createdAt, *id),
				)
			},
		)

	case id != nil:
		kindQuery = kindQuery.Where(
			kind2.IDLT(*id),
		)

	case createdAt != nil:
		kindQuery = kindQuery.Where(
			kind2.CreatedAtLTE(*createdAt),
		)
	}

	es, err := kindQuery.All(ctx)
	if err != nil {
		return nil, err
	}

	return r.kindModelsToQuery(es), nil
}

func (r *KindRepository) Update(
	ctx context.Context,
	id uuid.UUID,
	updateFn func(k *kind.Kind) (*kind.Kind, error),
) error {
	currentKind, err := r.Get(ctx, id)
	if err != nil {
		return err
	}

	tx, err := r.client.Tx(ctx)
	if err != nil {
		return err
	}

	updatedKind, err := updateFn(currentKind)
	if err != nil {
		return err
	}

	err = tx.Kind.UpdateOne(r.marshalKind(updatedKind)).Exec(ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return err
	}

	return tx.Commit()
}

func (r *KindRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	return r.client.Kind.DeleteOneID(id).Exec(ctx)
}

func (r *KindRepository) Exist(
	ctx context.Context,
	name string,
) (bool, error) {
	return r.client.Kind.Query().
		Where(kind2.NameEQ(name)).
		Exist(ctx)
}

func (r *KindRepository) marshalKind(k *kind.Kind) *ent2.Kind {
	return &ent2.Kind{
		ID:          k.ID(),
		CreatedAt:   k.CreatedAt(),
		UpdatedAt:   k.UpdatedAt(),
		Name:        k.Name(),
		Description: k.Description(),
		Status:      kind2.Status(k.Status().String()),
	}
}

func (r *KindRepository) unmarshalKind(e *ent2.Kind) *kind.Kind {
	return kind.UnmarshalFromDB(
		e.ID,
		e.CreatedAt,
		e.UpdatedAt,
		e.Name,
		e.Description,
		e.Status.String(),
	)
}

func (r *KindRepository) kindModelsToQuery(es []*ent2.Kind) []queries.Kind {
	return lo.Map[*ent2.Kind, queries.Kind](es, func(e *ent2.Kind, _ int) queries.Kind {
		return queries.Kind{
			ID:          e.ID,
			CreatedAt:   e.CreatedAt,
			UpdatedAt:   e.UpdatedAt,
			Name:        e.Name,
			Description: e.Description,
			Status:      e.Status.String(),
		}
	})
}

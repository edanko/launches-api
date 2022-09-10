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
	launch3 "github.com/edanko/launches-api/internal/adapters/ent/launch"
	"github.com/edanko/launches-api/internal/app/queries"
	launch2 "github.com/edanko/launches-api/internal/domain/launch"
)

type LaunchRepository struct {
	client *ent2.Client
}

var _ launch2.Repository = (*LaunchRepository)(nil)

func NewLaunchRepository(c *ent2.Client) *LaunchRepository {
	return &LaunchRepository{
		client: c,
	}
}

func (r *LaunchRepository) Create(
	ctx context.Context,
	k *launch2.Launch,
) error {
	return r.client.Launch.
		Create().
		SetID(k.ID()).
		SetCreatedAt(k.CreatedAt()).
		SetUpdatedAt(k.UpdatedAt()).
		// SetName(k.Name()).
		SetNillableDescription(k.Description()).
		SetStatus(launch3.Status(k.Status().String())).
		Exec(ctx)
}

func (r *LaunchRepository) Get(
	ctx context.Context,
	id uuid.UUID,
) (*launch2.Launch, error) {
	e, err := r.client.Launch.
		Query().
		Where(launch3.IDEQ(id)).
		Only(ctx)
	if ent2.IsNotFound(err) {
		return nil, launch2.ErrLaunchNotFound
	}
	if err != nil {
		return nil, errors.Wrap(err, "unable to get actual launch by id")
	}

	return r.unmarshalLaunch(e), nil
}

func (r *LaunchRepository) GetByName(
	ctx context.Context,
	name string,
) (*launch2.Launch, error) {
	e, err := r.client.Launch.
		Query().
		// Where(launch.NameEQ(name)).
		Only(ctx)
	if ent2.IsNotFound(err) {
		return nil, launch2.ErrLaunchNotFound
	}
	if err != nil {
		return nil, errors.Wrap(err, "unable to get actual launch by name")
	}

	return r.unmarshalLaunch(e), nil
}

func (r *LaunchRepository) List(
	ctx context.Context,
	limit *int,
	createdAt *time.Time,
	id *uuid.UUID,
	status *string,
) ([]queries.Launch, error) {
	launchQuery := r.client.Launch.Query().
		Order(
			ent2.Desc(launch3.FieldCreatedAt),
			ent2.Desc(launch3.FieldID),
		)

	if status != nil {
		launchQuery = launchQuery.Where(
			launch3.StatusEQ(launch3.Status(*status)),
		)
	}
	if limit != nil {
		launchQuery = launchQuery.Limit(*limit)
	}
	switch {
	case createdAt != nil && id != nil:
		launchQuery.Where(
			func(s *sql.Selector) {
				s.Where(
					sql.CompositeLT([]string{"created_at", "id"}, *createdAt, *id),
				)
			},
		)

	case id != nil:
		launchQuery = launchQuery.Where(
			launch3.IDLT(*id),
		)

	case createdAt != nil:
		launchQuery = launchQuery.Where(
			launch3.CreatedAtLTE(*createdAt),
		)
	}

	es, err := launchQuery.All(ctx)
	if err != nil {
		return nil, err
	}

	return r.launchModelsToQuery(es), nil
}

func (r *LaunchRepository) Update(
	ctx context.Context,
	id uuid.UUID,
	updateFn func(k *launch2.Launch) (*launch2.Launch, error),
) error {
	currentLaunch, err := r.Get(ctx, id)
	if err != nil {
		return err
	}

	tx, err := r.client.Tx(ctx)
	if err != nil {
		return err
	}

	updatedLaunch, err := updateFn(currentLaunch)
	if err != nil {
		return err
	}

	err = tx.Launch.UpdateOne(r.marshalLaunch(updatedLaunch)).Exec(ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return err
	}

	return tx.Commit()
}

func (r *LaunchRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	return r.client.Launch.DeleteOneID(id).Exec(ctx)
}

func (r *LaunchRepository) Exist(
	ctx context.Context,
	name string,
) (bool, error) {
	return r.client.Launch.Query().
		// Where(launch.NameEQ(name)).
		Exist(ctx)
}

func (r *LaunchRepository) marshalLaunch(k *launch2.Launch) *ent2.Launch {
	return &ent2.Launch{
		ID:        k.ID(),
		CreatedAt: k.CreatedAt(),
		UpdatedAt: k.UpdatedAt(),
		// Name:        k.Name(),
		Description: k.Description(),
		Status:      launch3.Status(k.Status().String()),
	}
}

func (r *LaunchRepository) unmarshalLaunch(e *ent2.Launch) *launch2.Launch {
	return launch2.UnmarshalFromDB(
		e.ID,
		e.CreatedAt,
		e.UpdatedAt,
		// e.Name,
		e.Description,
		e.Status.String(),
	)
}

func (r *LaunchRepository) launchModelsToQuery(es []*ent2.Launch) []queries.Launch {
	return lo.Map[*ent2.Launch, queries.Launch](es, func(e *ent2.Launch, _ int) queries.Launch {
		return queries.Launch{
			ID:        e.ID,
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
			// Name:        e.Name,
			Description: e.Description,
			Status:      e.Status.String(),
		}
	})
}

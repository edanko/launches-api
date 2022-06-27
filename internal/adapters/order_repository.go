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
	order3 "github.com/edanko/launches-api/internal/adapters/ent/order"
	"github.com/edanko/launches-api/internal/app/queries"
	order2 "github.com/edanko/launches-api/internal/domain/order"
)

type OrderRepository struct {
	client *ent2.Client
}

var _ order2.Repository = (*OrderRepository)(nil)

func NewOrderRepository(c *ent2.Client) *OrderRepository {
	return &OrderRepository{
		client: c,
	}
}

func (r *OrderRepository) Create(
	ctx context.Context,
	o *order2.Order,
) error {
	return r.client.Order.
		Create().
		SetID(o.ID()).
		SetCreatedAt(o.CreatedAt()).
		SetUpdatedAt(o.UpdatedAt()).
		SetName(o.Name()).
		SetNillableDescription(o.Description()).
		SetStatus(order3.Status(o.Status().String())).
		Exec(ctx)
}

func (r *OrderRepository) Get(
	ctx context.Context,
	id uuid.UUID,
) (*order2.Order, error) {
	e, err := r.client.Order.
		Query().
		Where(order3.IDEQ(id)).
		Only(ctx)
	if ent2.IsNotFound(err) {
		return nil, order2.ErrOrderNotFound
	}
	if err != nil {
		return nil, errors.Wrap(err, "unable to get actual order by id")
	}

	return r.unmarshalOrder(e), nil
}

func (r *OrderRepository) GetByName(
	ctx context.Context,
	name string,
) (*order2.Order, error) {
	e, err := r.client.Order.
		Query().
		Where(order3.NameEQ(name)).
		Only(ctx)
	if ent2.IsNotFound(err) {
		return nil, order2.ErrOrderNotFound
	}
	if err != nil {
		return nil, errors.Wrap(err, "unable to get actual order by name")
	}

	return r.unmarshalOrder(e), nil
}

func (r *OrderRepository) List(
	ctx context.Context,
	limit *int,
	createdAt *time.Time,
	id *uuid.UUID,
	status *string,
) ([]queries.Order, error) {
	orderQuery := r.client.Order.Query().
		Order(
			ent2.Desc(order3.FieldCreatedAt),
			ent2.Desc(order3.FieldID),
		)

	if status != nil {
		orderQuery = orderQuery.Where(
			order3.StatusEQ(order3.Status(*status)),
		)
	}
	if limit != nil {
		orderQuery = orderQuery.Limit(*limit)
	}
	switch {
	case createdAt != nil && id != nil:
		orderQuery.Where(
			func(s *sql.Selector) {
				s.Where(
					sql.CompositeLT([]string{"created_at", "id"}, *createdAt, *id),
				)
			},
		)

	case id != nil:
		orderQuery = orderQuery.Where(
			order3.IDLT(*id),
		)

	case createdAt != nil:
		orderQuery = orderQuery.Where(
			order3.CreatedAtLTE(*createdAt),
		)
	}

	es, err := orderQuery.All(ctx)
	if err != nil {
		return nil, err
	}

	return r.orderModelsToQuery(es), nil
}

func (r *OrderRepository) Update(
	ctx context.Context,
	id uuid.UUID,
	updateFn func(o *order2.Order) (*order2.Order, error),
) error {
	currentOrder, err := r.Get(ctx, id)
	if err != nil {
		return err
	}

	tx, err := r.client.Tx(ctx)
	if err != nil {
		return err
	}

	updatedOrder, err := updateFn(currentOrder)
	if err != nil {
		return err
	}

	err = tx.Order.UpdateOne(r.marshalOrder(updatedOrder)).Exec(ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return err
	}

	return tx.Commit()
}

func (r *OrderRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	return r.client.Order.DeleteOneID(id).Exec(ctx)
}

func (r *OrderRepository) Exist(
	ctx context.Context,
	name string,
) (bool, error) {
	return r.client.Order.Query().
		Where(order3.NameEQ(name)).
		Exist(ctx)
}

func (r *OrderRepository) marshalOrder(o *order2.Order) *ent2.Order {
	return &ent2.Order{
		ID:          o.ID(),
		CreatedAt:   o.CreatedAt(),
		UpdatedAt:   o.UpdatedAt(),
		Name:        o.Name(),
		Description: o.Description(),
		Status:      order3.Status(o.Status().String()),
	}
}

func (r *OrderRepository) unmarshalOrder(e *ent2.Order) *order2.Order {
	return order2.UnmarshalFromDB(
		e.ID,
		e.CreatedAt,
		e.UpdatedAt,
		e.Name,
		e.Description,
		e.Status.String(),
	)
}

func (r *OrderRepository) orderModelsToQuery(es []*ent2.Order) []queries.Order {
	return lo.Map[*ent2.Order, queries.Order](es, func(e *ent2.Order, _ int) queries.Order {
		return queries.Order{
			ID:          e.ID,
			CreatedAt:   e.CreatedAt,
			UpdatedAt:   e.UpdatedAt,
			Name:        e.Name,
			Description: e.Description,
			Status:      e.Status.String(),
		}
	})
}

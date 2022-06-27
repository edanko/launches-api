// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/edanko/nx/cmd/launch-api/internal/adapters/ent/launch"
	"github.com/edanko/nx/cmd/launch-api/internal/adapters/ent/predicate"
)

// LaunchDelete is the builder for deleting a Launch entity.
type LaunchDelete struct {
	config
	hooks    []Hook
	mutation *LaunchMutation
}

// Where appends a list predicates to the LaunchDelete builder.
func (ld *LaunchDelete) Where(ps ...predicate.Launch) *LaunchDelete {
	ld.mutation.Where(ps...)
	return ld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ld *LaunchDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ld.hooks) == 0 {
		affected, err = ld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LaunchMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ld.mutation = mutation
			affected, err = ld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ld.hooks) - 1; i >= 0; i-- {
			if ld.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ld *LaunchDelete) ExecX(ctx context.Context) int {
	n, err := ld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ld *LaunchDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: launch.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: launch.FieldID,
			},
		},
	}
	if ps := ld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ld.driver, _spec)
}

// LaunchDeleteOne is the builder for deleting a single Launch entity.
type LaunchDeleteOne struct {
	ld *LaunchDelete
}

// Exec executes the deletion query.
func (ldo *LaunchDeleteOne) Exec(ctx context.Context) error {
	n, err := ldo.ld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{launch.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ldo *LaunchDeleteOne) ExecX(ctx context.Context) {
	ldo.ld.ExecX(ctx)
}

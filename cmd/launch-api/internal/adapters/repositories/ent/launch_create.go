// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/edanko/nx/cmd/launch-api/internal/adapters/repositories/ent/launch"
	"github.com/google/uuid"
)

// LaunchCreate is the builder for creating a Launch entity.
type LaunchCreate struct {
	config
	mutation *LaunchMutation
	hooks    []Hook
}

// SetStatus sets the "status" field.
func (lc *LaunchCreate) SetStatus(l launch.Status) *LaunchCreate {
	lc.mutation.SetStatus(l)
	return lc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (lc *LaunchCreate) SetNillableStatus(l *launch.Status) *LaunchCreate {
	if l != nil {
		lc.SetStatus(*l)
	}
	return lc
}

// SetApplicant sets the "applicant" field.
func (lc *LaunchCreate) SetApplicant(s string) *LaunchCreate {
	lc.mutation.SetApplicant(s)
	return lc
}

// SetReason sets the "reason" field.
func (lc *LaunchCreate) SetReason(s string) *LaunchCreate {
	lc.mutation.SetReason(s)
	return lc
}

// SetDescription sets the "description" field.
func (lc *LaunchCreate) SetDescription(s string) *LaunchCreate {
	lc.mutation.SetDescription(s)
	return lc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (lc *LaunchCreate) SetNillableDescription(s *string) *LaunchCreate {
	if s != nil {
		lc.SetDescription(*s)
	}
	return lc
}

// SetID sets the "id" field.
func (lc *LaunchCreate) SetID(u uuid.UUID) *LaunchCreate {
	lc.mutation.SetID(u)
	return lc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (lc *LaunchCreate) SetNillableID(u *uuid.UUID) *LaunchCreate {
	if u != nil {
		lc.SetID(*u)
	}
	return lc
}

// Mutation returns the LaunchMutation object of the builder.
func (lc *LaunchCreate) Mutation() *LaunchMutation {
	return lc.mutation
}

// Save creates the Launch in the database.
func (lc *LaunchCreate) Save(ctx context.Context) (*Launch, error) {
	var (
		err  error
		node *Launch
	)
	lc.defaults()
	if len(lc.hooks) == 0 {
		if err = lc.check(); err != nil {
			return nil, err
		}
		node, err = lc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LaunchMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lc.check(); err != nil {
				return nil, err
			}
			lc.mutation = mutation
			if node, err = lc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(lc.hooks) - 1; i >= 0; i-- {
			if lc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LaunchCreate) SaveX(ctx context.Context) *Launch {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LaunchCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LaunchCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LaunchCreate) defaults() {
	if _, ok := lc.mutation.Status(); !ok {
		v := launch.DefaultStatus
		lc.mutation.SetStatus(v)
	}
	if _, ok := lc.mutation.ID(); !ok {
		v := launch.DefaultID()
		lc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LaunchCreate) check() error {
	if _, ok := lc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Launch.status"`)}
	}
	if v, ok := lc.mutation.Status(); ok {
		if err := launch.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Launch.status": %w`, err)}
		}
	}
	if _, ok := lc.mutation.Applicant(); !ok {
		return &ValidationError{Name: "applicant", err: errors.New(`ent: missing required field "Launch.applicant"`)}
	}
	if v, ok := lc.mutation.Applicant(); ok {
		if err := launch.ApplicantValidator(v); err != nil {
			return &ValidationError{Name: "applicant", err: fmt.Errorf(`ent: validator failed for field "Launch.applicant": %w`, err)}
		}
	}
	if _, ok := lc.mutation.Reason(); !ok {
		return &ValidationError{Name: "reason", err: errors.New(`ent: missing required field "Launch.reason"`)}
	}
	if v, ok := lc.mutation.Reason(); ok {
		if err := launch.ReasonValidator(v); err != nil {
			return &ValidationError{Name: "reason", err: fmt.Errorf(`ent: validator failed for field "Launch.reason": %w`, err)}
		}
	}
	return nil
}

func (lc *LaunchCreate) sqlSave(ctx context.Context) (*Launch, error) {
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (lc *LaunchCreate) createSpec() (*Launch, *sqlgraph.CreateSpec) {
	var (
		_node = &Launch{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: launch.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: launch.FieldID,
			},
		}
	)
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := lc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: launch.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := lc.mutation.Applicant(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: launch.FieldApplicant,
		})
		_node.Applicant = value
	}
	if value, ok := lc.mutation.Reason(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: launch.FieldReason,
		})
		_node.Reason = value
	}
	if value, ok := lc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: launch.FieldDescription,
		})
		_node.Description = &value
	}
	return _node, _spec
}

// LaunchCreateBulk is the builder for creating many Launch entities in bulk.
type LaunchCreateBulk struct {
	config
	builders []*LaunchCreate
}

// Save creates the Launch entities in the database.
func (lcb *LaunchCreateBulk) Save(ctx context.Context) ([]*Launch, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Launch, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LaunchMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LaunchCreateBulk) SaveX(ctx context.Context) []*Launch {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LaunchCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LaunchCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}

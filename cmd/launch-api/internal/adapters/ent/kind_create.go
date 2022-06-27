// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/edanko/nx/cmd/launch-api/internal/adapters/ent/kind"
	"github.com/edanko/nx/cmd/launch-api/internal/adapters/ent/launch"
	"github.com/google/uuid"
)

// KindCreate is the builder for creating a Kind entity.
type KindCreate struct {
	config
	mutation *KindMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (kc *KindCreate) SetCreatedAt(t time.Time) *KindCreate {
	kc.mutation.SetCreatedAt(t)
	return kc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (kc *KindCreate) SetNillableCreatedAt(t *time.Time) *KindCreate {
	if t != nil {
		kc.SetCreatedAt(*t)
	}
	return kc
}

// SetUpdatedAt sets the "updated_at" field.
func (kc *KindCreate) SetUpdatedAt(t time.Time) *KindCreate {
	kc.mutation.SetUpdatedAt(t)
	return kc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (kc *KindCreate) SetNillableUpdatedAt(t *time.Time) *KindCreate {
	if t != nil {
		kc.SetUpdatedAt(*t)
	}
	return kc
}

// SetName sets the "name" field.
func (kc *KindCreate) SetName(s string) *KindCreate {
	kc.mutation.SetName(s)
	return kc
}

// SetDescription sets the "description" field.
func (kc *KindCreate) SetDescription(s string) *KindCreate {
	kc.mutation.SetDescription(s)
	return kc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (kc *KindCreate) SetNillableDescription(s *string) *KindCreate {
	if s != nil {
		kc.SetDescription(*s)
	}
	return kc
}

// SetStatus sets the "status" field.
func (kc *KindCreate) SetStatus(k kind.Status) *KindCreate {
	kc.mutation.SetStatus(k)
	return kc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (kc *KindCreate) SetNillableStatus(k *kind.Status) *KindCreate {
	if k != nil {
		kc.SetStatus(*k)
	}
	return kc
}

// SetID sets the "id" field.
func (kc *KindCreate) SetID(u uuid.UUID) *KindCreate {
	kc.mutation.SetID(u)
	return kc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (kc *KindCreate) SetNillableID(u *uuid.UUID) *KindCreate {
	if u != nil {
		kc.SetID(*u)
	}
	return kc
}

// AddLaunchIDs adds the "launches" edge to the Launch entity by IDs.
func (kc *KindCreate) AddLaunchIDs(ids ...uuid.UUID) *KindCreate {
	kc.mutation.AddLaunchIDs(ids...)
	return kc
}

// AddLaunches adds the "launches" edges to the Launch entity.
func (kc *KindCreate) AddLaunches(l ...*Launch) *KindCreate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return kc.AddLaunchIDs(ids...)
}

// Mutation returns the KindMutation object of the builder.
func (kc *KindCreate) Mutation() *KindMutation {
	return kc.mutation
}

// Save creates the Kind in the database.
func (kc *KindCreate) Save(ctx context.Context) (*Kind, error) {
	var (
		err  error
		node *Kind
	)
	kc.defaults()
	if len(kc.hooks) == 0 {
		if err = kc.check(); err != nil {
			return nil, err
		}
		node, err = kc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KindMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kc.check(); err != nil {
				return nil, err
			}
			kc.mutation = mutation
			if node, err = kc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(kc.hooks) - 1; i >= 0; i-- {
			if kc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (kc *KindCreate) SaveX(ctx context.Context) *Kind {
	v, err := kc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kc *KindCreate) Exec(ctx context.Context) error {
	_, err := kc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kc *KindCreate) ExecX(ctx context.Context) {
	if err := kc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kc *KindCreate) defaults() {
	if _, ok := kc.mutation.CreatedAt(); !ok {
		v := kind.DefaultCreatedAt()
		kc.mutation.SetCreatedAt(v)
	}
	if _, ok := kc.mutation.UpdatedAt(); !ok {
		v := kind.DefaultUpdatedAt()
		kc.mutation.SetUpdatedAt(v)
	}
	if _, ok := kc.mutation.Status(); !ok {
		v := kind.DefaultStatus
		kc.mutation.SetStatus(v)
	}
	if _, ok := kc.mutation.ID(); !ok {
		v := kind.DefaultID()
		kc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kc *KindCreate) check() error {
	if _, ok := kc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Kind.created_at"`)}
	}
	if _, ok := kc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Kind.updated_at"`)}
	}
	if _, ok := kc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Kind.name"`)}
	}
	if v, ok := kc.mutation.Name(); ok {
		if err := kind.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Kind.name": %w`, err)}
		}
	}
	if _, ok := kc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Kind.status"`)}
	}
	if v, ok := kc.mutation.Status(); ok {
		if err := kind.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Kind.status": %w`, err)}
		}
	}
	return nil
}

func (kc *KindCreate) sqlSave(ctx context.Context) (*Kind, error) {
	_node, _spec := kc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kc.driver, _spec); err != nil {
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

func (kc *KindCreate) createSpec() (*Kind, *sqlgraph.CreateSpec) {
	var (
		_node = &Kind{config: kc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: kind.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: kind.FieldID,
			},
		}
	)
	if id, ok := kc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := kc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kind.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := kc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: kind.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := kc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kind.FieldName,
		})
		_node.Name = value
	}
	if value, ok := kc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kind.FieldDescription,
		})
		_node.Description = &value
	}
	if value, ok := kc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: kind.FieldStatus,
		})
		_node.Status = value
	}
	if nodes := kc.mutation.LaunchesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   kind.LaunchesTable,
			Columns: []string{kind.LaunchesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: launch.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// KindCreateBulk is the builder for creating many Kind entities in bulk.
type KindCreateBulk struct {
	config
	builders []*KindCreate
}

// Save creates the Kind entities in the database.
func (kcb *KindCreateBulk) Save(ctx context.Context) ([]*Kind, error) {
	specs := make([]*sqlgraph.CreateSpec, len(kcb.builders))
	nodes := make([]*Kind, len(kcb.builders))
	mutators := make([]Mutator, len(kcb.builders))
	for i := range kcb.builders {
		func(i int, root context.Context) {
			builder := kcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KindMutation)
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
					_, err = mutators[i+1].Mutate(root, kcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, kcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kcb *KindCreateBulk) SaveX(ctx context.Context) []*Kind {
	v, err := kcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kcb *KindCreateBulk) Exec(ctx context.Context) error {
	_, err := kcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kcb *KindCreateBulk) ExecX(ctx context.Context) {
	if err := kcb.Exec(ctx); err != nil {
		panic(err)
	}
}

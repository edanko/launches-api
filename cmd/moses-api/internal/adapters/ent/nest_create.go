// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/edanko/nx/cmd/moses-api/internal/adapters/ent/nest"
	"github.com/edanko/nx/cmd/moses-api/internal/adapters/ent/part"
	"github.com/edanko/nx/cmd/moses-api/internal/adapters/ent/remnant"
	"github.com/google/uuid"
)

// NestCreate is the builder for creating a Nest entity.
type NestCreate struct {
	config
	mutation *NestMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (nc *NestCreate) SetCreatedAt(t time.Time) *NestCreate {
	nc.mutation.SetCreatedAt(t)
	return nc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nc *NestCreate) SetNillableCreatedAt(t *time.Time) *NestCreate {
	if t != nil {
		nc.SetCreatedAt(*t)
	}
	return nc
}

// SetUpdatedAt sets the "updated_at" field.
func (nc *NestCreate) SetUpdatedAt(t time.Time) *NestCreate {
	nc.mutation.SetUpdatedAt(t)
	return nc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (nc *NestCreate) SetNillableUpdatedAt(t *time.Time) *NestCreate {
	if t != nil {
		nc.SetUpdatedAt(*t)
	}
	return nc
}

// SetName sets the "name" field.
func (nc *NestCreate) SetName(s string) *NestCreate {
	nc.mutation.SetName(s)
	return nc
}

// SetLength sets the "length" field.
func (nc *NestCreate) SetLength(f float64) *NestCreate {
	nc.mutation.SetLength(f)
	return nc
}

// SetID sets the "id" field.
func (nc *NestCreate) SetID(u uuid.UUID) *NestCreate {
	nc.mutation.SetID(u)
	return nc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (nc *NestCreate) SetNillableID(u *uuid.UUID) *NestCreate {
	if u != nil {
		nc.SetID(*u)
	}
	return nc
}

// AddPartIDs adds the "parts" edge to the Part entity by IDs.
func (nc *NestCreate) AddPartIDs(ids ...uuid.UUID) *NestCreate {
	nc.mutation.AddPartIDs(ids...)
	return nc
}

// AddParts adds the "parts" edges to the Part entity.
func (nc *NestCreate) AddParts(p ...*Part) *NestCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return nc.AddPartIDs(ids...)
}

// SetRemnantID sets the "remnant" edge to the Remnant entity by ID.
func (nc *NestCreate) SetRemnantID(id uuid.UUID) *NestCreate {
	nc.mutation.SetRemnantID(id)
	return nc
}

// SetNillableRemnantID sets the "remnant" edge to the Remnant entity by ID if the given value is not nil.
func (nc *NestCreate) SetNillableRemnantID(id *uuid.UUID) *NestCreate {
	if id != nil {
		nc = nc.SetRemnantID(*id)
	}
	return nc
}

// SetRemnant sets the "remnant" edge to the Remnant entity.
func (nc *NestCreate) SetRemnant(r *Remnant) *NestCreate {
	return nc.SetRemnantID(r.ID)
}

// SetRemnantUsedID sets the "remnant_used" edge to the Remnant entity by ID.
func (nc *NestCreate) SetRemnantUsedID(id uuid.UUID) *NestCreate {
	nc.mutation.SetRemnantUsedID(id)
	return nc
}

// SetNillableRemnantUsedID sets the "remnant_used" edge to the Remnant entity by ID if the given value is not nil.
func (nc *NestCreate) SetNillableRemnantUsedID(id *uuid.UUID) *NestCreate {
	if id != nil {
		nc = nc.SetRemnantUsedID(*id)
	}
	return nc
}

// SetRemnantUsed sets the "remnant_used" edge to the Remnant entity.
func (nc *NestCreate) SetRemnantUsed(r *Remnant) *NestCreate {
	return nc.SetRemnantUsedID(r.ID)
}

// Mutation returns the NestMutation object of the builder.
func (nc *NestCreate) Mutation() *NestMutation {
	return nc.mutation
}

// Save creates the Nest in the database.
func (nc *NestCreate) Save(ctx context.Context) (*Nest, error) {
	var (
		err  error
		node *Nest
	)
	nc.defaults()
	if len(nc.hooks) == 0 {
		if err = nc.check(); err != nil {
			return nil, err
		}
		node, err = nc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nc.check(); err != nil {
				return nil, err
			}
			nc.mutation = mutation
			if node, err = nc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(nc.hooks) - 1; i >= 0; i-- {
			if nc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, nc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Nest)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from NestMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NestCreate) SaveX(ctx context.Context) *Nest {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NestCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NestCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nc *NestCreate) defaults() {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		v := nest.DefaultCreatedAt()
		nc.mutation.SetCreatedAt(v)
	}
	if _, ok := nc.mutation.UpdatedAt(); !ok {
		v := nest.DefaultUpdatedAt()
		nc.mutation.SetUpdatedAt(v)
	}
	if _, ok := nc.mutation.ID(); !ok {
		v := nest.DefaultID()
		nc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NestCreate) check() error {
	if _, ok := nc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Nest.created_at"`)}
	}
	if _, ok := nc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Nest.updated_at"`)}
	}
	if _, ok := nc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Nest.name"`)}
	}
	if v, ok := nc.mutation.Name(); ok {
		if err := nest.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Nest.name": %w`, err)}
		}
	}
	if _, ok := nc.mutation.Length(); !ok {
		return &ValidationError{Name: "length", err: errors.New(`ent: missing required field "Nest.length"`)}
	}
	if v, ok := nc.mutation.Length(); ok {
		if err := nest.LengthValidator(v); err != nil {
			return &ValidationError{Name: "length", err: fmt.Errorf(`ent: validator failed for field "Nest.length": %w`, err)}
		}
	}
	if len(nc.mutation.PartsIDs()) == 0 {
		return &ValidationError{Name: "parts", err: errors.New(`ent: missing required edge "Nest.parts"`)}
	}
	return nil
}

func (nc *NestCreate) sqlSave(ctx context.Context) (*Nest, error) {
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
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

func (nc *NestCreate) createSpec() (*Nest, *sqlgraph.CreateSpec) {
	var (
		_node = &Nest{config: nc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: nest.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: nest.FieldID,
			},
		}
	)
	if id, ok := nc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := nc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: nest.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := nc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: nest.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := nc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: nest.FieldName,
		})
		_node.Name = value
	}
	if value, ok := nc.mutation.Length(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: nest.FieldLength,
		})
		_node.Length = value
	}
	if nodes := nc.mutation.PartsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   nest.PartsTable,
			Columns: nest.PartsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: part.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.RemnantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   nest.RemnantTable,
			Columns: []string{nest.RemnantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: remnant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.RemnantUsedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   nest.RemnantUsedTable,
			Columns: []string{nest.RemnantUsedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: remnant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.remnant_remnant_used = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// NestCreateBulk is the builder for creating many Nest entities in bulk.
type NestCreateBulk struct {
	config
	builders []*NestCreate
}

// Save creates the Nest entities in the database.
func (ncb *NestCreateBulk) Save(ctx context.Context) ([]*Nest, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Nest, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NestMutation)
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
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
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
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NestCreateBulk) SaveX(ctx context.Context) []*Nest {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NestCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NestCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}

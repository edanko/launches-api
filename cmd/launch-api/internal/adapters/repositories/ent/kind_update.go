// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/edanko/nx/cmd/launch-api/internal/adapters/repositories/ent/kind"
	"github.com/edanko/nx/cmd/launch-api/internal/adapters/repositories/ent/predicate"
)

// KindUpdate is the builder for updating Kind entities.
type KindUpdate struct {
	config
	hooks    []Hook
	mutation *KindMutation
}

// Where appends a list predicates to the KindUpdate builder.
func (ku *KindUpdate) Where(ps ...predicate.Kind) *KindUpdate {
	ku.mutation.Where(ps...)
	return ku
}

// SetName sets the "name" field.
func (ku *KindUpdate) SetName(s string) *KindUpdate {
	ku.mutation.SetName(s)
	return ku
}

// SetDescription sets the "description" field.
func (ku *KindUpdate) SetDescription(s string) *KindUpdate {
	ku.mutation.SetDescription(s)
	return ku
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ku *KindUpdate) SetNillableDescription(s *string) *KindUpdate {
	if s != nil {
		ku.SetDescription(*s)
	}
	return ku
}

// ClearDescription clears the value of the "description" field.
func (ku *KindUpdate) ClearDescription() *KindUpdate {
	ku.mutation.ClearDescription()
	return ku
}

// SetStatus sets the "status" field.
func (ku *KindUpdate) SetStatus(k kind.Status) *KindUpdate {
	ku.mutation.SetStatus(k)
	return ku
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ku *KindUpdate) SetNillableStatus(k *kind.Status) *KindUpdate {
	if k != nil {
		ku.SetStatus(*k)
	}
	return ku
}

// Mutation returns the KindMutation object of the builder.
func (ku *KindUpdate) Mutation() *KindMutation {
	return ku.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ku *KindUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ku.hooks) == 0 {
		if err = ku.check(); err != nil {
			return 0, err
		}
		affected, err = ku.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KindMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ku.check(); err != nil {
				return 0, err
			}
			ku.mutation = mutation
			affected, err = ku.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ku.hooks) - 1; i >= 0; i-- {
			if ku.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ku.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ku.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ku *KindUpdate) SaveX(ctx context.Context) int {
	affected, err := ku.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ku *KindUpdate) Exec(ctx context.Context) error {
	_, err := ku.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ku *KindUpdate) ExecX(ctx context.Context) {
	if err := ku.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ku *KindUpdate) check() error {
	if v, ok := ku.mutation.Name(); ok {
		if err := kind.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Kind.name": %w`, err)}
		}
	}
	if v, ok := ku.mutation.Status(); ok {
		if err := kind.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Kind.status": %w`, err)}
		}
	}
	return nil
}

func (ku *KindUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   kind.Table,
			Columns: kind.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: kind.FieldID,
			},
		},
	}
	if ps := ku.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ku.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kind.FieldName,
		})
	}
	if value, ok := ku.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kind.FieldDescription,
		})
	}
	if ku.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: kind.FieldDescription,
		})
	}
	if value, ok := ku.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: kind.FieldStatus,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ku.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{kind.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// KindUpdateOne is the builder for updating a single Kind entity.
type KindUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *KindMutation
}

// SetName sets the "name" field.
func (kuo *KindUpdateOne) SetName(s string) *KindUpdateOne {
	kuo.mutation.SetName(s)
	return kuo
}

// SetDescription sets the "description" field.
func (kuo *KindUpdateOne) SetDescription(s string) *KindUpdateOne {
	kuo.mutation.SetDescription(s)
	return kuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (kuo *KindUpdateOne) SetNillableDescription(s *string) *KindUpdateOne {
	if s != nil {
		kuo.SetDescription(*s)
	}
	return kuo
}

// ClearDescription clears the value of the "description" field.
func (kuo *KindUpdateOne) ClearDescription() *KindUpdateOne {
	kuo.mutation.ClearDescription()
	return kuo
}

// SetStatus sets the "status" field.
func (kuo *KindUpdateOne) SetStatus(k kind.Status) *KindUpdateOne {
	kuo.mutation.SetStatus(k)
	return kuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (kuo *KindUpdateOne) SetNillableStatus(k *kind.Status) *KindUpdateOne {
	if k != nil {
		kuo.SetStatus(*k)
	}
	return kuo
}

// Mutation returns the KindMutation object of the builder.
func (kuo *KindUpdateOne) Mutation() *KindMutation {
	return kuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (kuo *KindUpdateOne) Select(field string, fields ...string) *KindUpdateOne {
	kuo.fields = append([]string{field}, fields...)
	return kuo
}

// Save executes the query and returns the updated Kind entity.
func (kuo *KindUpdateOne) Save(ctx context.Context) (*Kind, error) {
	var (
		err  error
		node *Kind
	)
	if len(kuo.hooks) == 0 {
		if err = kuo.check(); err != nil {
			return nil, err
		}
		node, err = kuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KindMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kuo.check(); err != nil {
				return nil, err
			}
			kuo.mutation = mutation
			node, err = kuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(kuo.hooks) - 1; i >= 0; i-- {
			if kuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (kuo *KindUpdateOne) SaveX(ctx context.Context) *Kind {
	node, err := kuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (kuo *KindUpdateOne) Exec(ctx context.Context) error {
	_, err := kuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kuo *KindUpdateOne) ExecX(ctx context.Context) {
	if err := kuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kuo *KindUpdateOne) check() error {
	if v, ok := kuo.mutation.Name(); ok {
		if err := kind.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Kind.name": %w`, err)}
		}
	}
	if v, ok := kuo.mutation.Status(); ok {
		if err := kind.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Kind.status": %w`, err)}
		}
	}
	return nil
}

func (kuo *KindUpdateOne) sqlSave(ctx context.Context) (_node *Kind, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   kind.Table,
			Columns: kind.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: kind.FieldID,
			},
		},
	}
	id, ok := kuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Kind.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := kuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, kind.FieldID)
		for _, f := range fields {
			if !kind.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != kind.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := kuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kind.FieldName,
		})
	}
	if value, ok := kuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kind.FieldDescription,
		})
	}
	if kuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: kind.FieldDescription,
		})
	}
	if value, ok := kuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: kind.FieldStatus,
		})
	}
	_node = &Kind{config: kuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, kuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{kind.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

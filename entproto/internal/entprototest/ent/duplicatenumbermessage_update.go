// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/duplicatenumbermessage"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DuplicateNumberMessageUpdate is the builder for updating DuplicateNumberMessage entities.
type DuplicateNumberMessageUpdate struct {
	config
	hooks    []Hook
	mutation *DuplicateNumberMessageMutation
}

// Where appends a list predicates to the DuplicateNumberMessageUpdate builder.
func (dnmu *DuplicateNumberMessageUpdate) Where(ps ...predicate.DuplicateNumberMessage) *DuplicateNumberMessageUpdate {
	dnmu.mutation.Where(ps...)
	return dnmu
}

// SetHello sets the "hello" field.
func (dnmu *DuplicateNumberMessageUpdate) SetHello(s string) *DuplicateNumberMessageUpdate {
	dnmu.mutation.SetHello(s)
	return dnmu
}

// SetWorld sets the "world" field.
func (dnmu *DuplicateNumberMessageUpdate) SetWorld(s string) *DuplicateNumberMessageUpdate {
	dnmu.mutation.SetWorld(s)
	return dnmu
}

// Mutation returns the DuplicateNumberMessageMutation object of the builder.
func (dnmu *DuplicateNumberMessageUpdate) Mutation() *DuplicateNumberMessageMutation {
	return dnmu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dnmu *DuplicateNumberMessageUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dnmu.hooks) == 0 {
		affected, err = dnmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DuplicateNumberMessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dnmu.mutation = mutation
			affected, err = dnmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dnmu.hooks) - 1; i >= 0; i-- {
			if dnmu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dnmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dnmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (dnmu *DuplicateNumberMessageUpdate) SaveX(ctx context.Context) int {
	affected, err := dnmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dnmu *DuplicateNumberMessageUpdate) Exec(ctx context.Context) error {
	_, err := dnmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dnmu *DuplicateNumberMessageUpdate) ExecX(ctx context.Context) {
	if err := dnmu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (dnmu *DuplicateNumberMessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   duplicatenumbermessage.Table,
			Columns: duplicatenumbermessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: duplicatenumbermessage.FieldID,
			},
		},
	}
	if ps := dnmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dnmu.mutation.Hello(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: duplicatenumbermessage.FieldHello,
		})
	}
	if value, ok := dnmu.mutation.World(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: duplicatenumbermessage.FieldWorld,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dnmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{duplicatenumbermessage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// DuplicateNumberMessageUpdateOne is the builder for updating a single DuplicateNumberMessage entity.
type DuplicateNumberMessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DuplicateNumberMessageMutation
}

// SetHello sets the "hello" field.
func (dnmuo *DuplicateNumberMessageUpdateOne) SetHello(s string) *DuplicateNumberMessageUpdateOne {
	dnmuo.mutation.SetHello(s)
	return dnmuo
}

// SetWorld sets the "world" field.
func (dnmuo *DuplicateNumberMessageUpdateOne) SetWorld(s string) *DuplicateNumberMessageUpdateOne {
	dnmuo.mutation.SetWorld(s)
	return dnmuo
}

// Mutation returns the DuplicateNumberMessageMutation object of the builder.
func (dnmuo *DuplicateNumberMessageUpdateOne) Mutation() *DuplicateNumberMessageMutation {
	return dnmuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dnmuo *DuplicateNumberMessageUpdateOne) Select(field string, fields ...string) *DuplicateNumberMessageUpdateOne {
	dnmuo.fields = append([]string{field}, fields...)
	return dnmuo
}

// Save executes the query and returns the updated DuplicateNumberMessage entity.
func (dnmuo *DuplicateNumberMessageUpdateOne) Save(ctx context.Context) (*DuplicateNumberMessage, error) {
	var (
		err  error
		node *DuplicateNumberMessage
	)
	if len(dnmuo.hooks) == 0 {
		node, err = dnmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DuplicateNumberMessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dnmuo.mutation = mutation
			node, err = dnmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dnmuo.hooks) - 1; i >= 0; i-- {
			if dnmuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dnmuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dnmuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*DuplicateNumberMessage)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DuplicateNumberMessageMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (dnmuo *DuplicateNumberMessageUpdateOne) SaveX(ctx context.Context) *DuplicateNumberMessage {
	node, err := dnmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dnmuo *DuplicateNumberMessageUpdateOne) Exec(ctx context.Context) error {
	_, err := dnmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dnmuo *DuplicateNumberMessageUpdateOne) ExecX(ctx context.Context) {
	if err := dnmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (dnmuo *DuplicateNumberMessageUpdateOne) sqlSave(ctx context.Context) (_node *DuplicateNumberMessage, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   duplicatenumbermessage.Table,
			Columns: duplicatenumbermessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: duplicatenumbermessage.FieldID,
			},
		},
	}
	id, ok := dnmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DuplicateNumberMessage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dnmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, duplicatenumbermessage.FieldID)
		for _, f := range fields {
			if !duplicatenumbermessage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != duplicatenumbermessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dnmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dnmuo.mutation.Hello(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: duplicatenumbermessage.FieldHello,
		})
	}
	if value, ok := dnmuo.mutation.World(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: duplicatenumbermessage.FieldWorld,
		})
	}
	_node = &DuplicateNumberMessage{config: dnmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dnmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{duplicatenumbermessage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

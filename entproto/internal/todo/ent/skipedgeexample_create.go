// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/contrib/entproto/internal/todo/ent/skipedgeexample"
	"entgo.io/contrib/entproto/internal/todo/ent/user"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SkipEdgeExampleCreate is the builder for creating a SkipEdgeExample entity.
type SkipEdgeExampleCreate struct {
	config
	mutation *SkipEdgeExampleMutation
	hooks    []Hook
}

// SetUserID sets the "user" edge to the User entity by ID.
func (seec *SkipEdgeExampleCreate) SetUserID(id int) *SkipEdgeExampleCreate {
	seec.mutation.SetUserID(id)
	return seec
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (seec *SkipEdgeExampleCreate) SetNillableUserID(id *int) *SkipEdgeExampleCreate {
	if id != nil {
		seec = seec.SetUserID(*id)
	}
	return seec
}

// SetUser sets the "user" edge to the User entity.
func (seec *SkipEdgeExampleCreate) SetUser(u *User) *SkipEdgeExampleCreate {
	return seec.SetUserID(u.ID)
}

// Mutation returns the SkipEdgeExampleMutation object of the builder.
func (seec *SkipEdgeExampleCreate) Mutation() *SkipEdgeExampleMutation {
	return seec.mutation
}

// Save creates the SkipEdgeExample in the database.
func (seec *SkipEdgeExampleCreate) Save(ctx context.Context) (*SkipEdgeExample, error) {
	var (
		err  error
		node *SkipEdgeExample
	)
	if len(seec.hooks) == 0 {
		if err = seec.check(); err != nil {
			return nil, err
		}
		node, err = seec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SkipEdgeExampleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = seec.check(); err != nil {
				return nil, err
			}
			seec.mutation = mutation
			if node, err = seec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(seec.hooks) - 1; i >= 0; i-- {
			if seec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = seec.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, seec.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SkipEdgeExample)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SkipEdgeExampleMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (seec *SkipEdgeExampleCreate) SaveX(ctx context.Context) *SkipEdgeExample {
	v, err := seec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (seec *SkipEdgeExampleCreate) Exec(ctx context.Context) error {
	_, err := seec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (seec *SkipEdgeExampleCreate) ExecX(ctx context.Context) {
	if err := seec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (seec *SkipEdgeExampleCreate) check() error {
	return nil
}

func (seec *SkipEdgeExampleCreate) sqlSave(ctx context.Context) (*SkipEdgeExample, error) {
	_node, _spec := seec.createSpec()
	if err := sqlgraph.CreateNode(ctx, seec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (seec *SkipEdgeExampleCreate) createSpec() (*SkipEdgeExample, *sqlgraph.CreateSpec) {
	var (
		_node = &SkipEdgeExample{config: seec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: skipedgeexample.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: skipedgeexample.FieldID,
			},
		}
	)
	if nodes := seec.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   skipedgeexample.UserTable,
			Columns: []string{skipedgeexample.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_skip_edge = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SkipEdgeExampleCreateBulk is the builder for creating many SkipEdgeExample entities in bulk.
type SkipEdgeExampleCreateBulk struct {
	config
	builders []*SkipEdgeExampleCreate
}

// Save creates the SkipEdgeExample entities in the database.
func (seecb *SkipEdgeExampleCreateBulk) Save(ctx context.Context) ([]*SkipEdgeExample, error) {
	specs := make([]*sqlgraph.CreateSpec, len(seecb.builders))
	nodes := make([]*SkipEdgeExample, len(seecb.builders))
	mutators := make([]Mutator, len(seecb.builders))
	for i := range seecb.builders {
		func(i int, root context.Context) {
			builder := seecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SkipEdgeExampleMutation)
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
					_, err = mutators[i+1].Mutate(root, seecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, seecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, seecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (seecb *SkipEdgeExampleCreateBulk) SaveX(ctx context.Context) []*SkipEdgeExample {
	v, err := seecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (seecb *SkipEdgeExampleCreateBulk) Exec(ctx context.Context) error {
	_, err := seecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (seecb *SkipEdgeExampleCreateBulk) ExecX(ctx context.Context) {
	if err := seecb.Exec(ctx); err != nil {
		panic(err)
	}
}

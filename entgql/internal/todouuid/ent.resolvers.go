// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package todo

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"entgo.io/contrib/entgql/internal/todo/ent/todo"
	"entgo.io/contrib/entgql/internal/todouuid/ent"
	"github.com/google/uuid"
)

func (r *queryResolver) Node(ctx context.Context, id uuid.UUID) (ent.Noder, error) {
	return r.client.Noder(ctx, id, ent.WithFixedNodeType(todo.Table))
}

func (r *queryResolver) Nodes(ctx context.Context, ids []uuid.UUID) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids, ent.WithFixedNodeType(todo.Table))
}

func (r *queryResolver) Groups(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.GroupWhereInput) (*ent.GroupConnection, error) {
	return r.client.Group.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithGroupFilter(where.Filter),
		)
}

func (r *queryResolver) Todos(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TodoOrder, where *ent.TodoWhereInput) (*ent.TodoConnection, error) {
	return r.client.Todo.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithTodoOrder(orderBy),
			ent.WithTodoFilter(where.Filter),
		)
}

func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return r.client.User.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithUserFilter(where.Filter),
		)
}

func (r *todoResolver) Status(ctx context.Context, obj *ent.Todo) (todo.Status, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *createTodoInputResolver) Status(ctx context.Context, obj *ent.CreateTodoInput, data todo.Status) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoWhereInputResolver) Status(ctx context.Context, obj *ent.TodoWhereInput, data *todo.Status) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoWhereInputResolver) StatusNeq(ctx context.Context, obj *ent.TodoWhereInput, data *todo.Status) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoWhereInputResolver) StatusIn(ctx context.Context, obj *ent.TodoWhereInput, data []todo.Status) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoWhereInputResolver) StatusNotIn(ctx context.Context, obj *ent.TodoWhereInput, data []todo.Status) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *updateTodoInputResolver) Status(ctx context.Context, obj *ent.UpdateTodoInput, data *todo.Status) error {
	panic(fmt.Errorf("not implemented"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

// CreateTodoInput returns CreateTodoInputResolver implementation.
func (r *Resolver) CreateTodoInput() CreateTodoInputResolver { return &createTodoInputResolver{r} }

// TodoWhereInput returns TodoWhereInputResolver implementation.
func (r *Resolver) TodoWhereInput() TodoWhereInputResolver { return &todoWhereInputResolver{r} }

// UpdateTodoInput returns UpdateTodoInputResolver implementation.
func (r *Resolver) UpdateTodoInput() UpdateTodoInputResolver { return &updateTodoInputResolver{r} }

type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type createTodoInputResolver struct{ *Resolver }
type todoWhereInputResolver struct{ *Resolver }
type updateTodoInputResolver struct{ *Resolver }

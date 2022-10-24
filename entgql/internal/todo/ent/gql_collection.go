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
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"

	"entgo.io/contrib/entgql/internal/todo/ent/category"
	"entgo.io/contrib/entgql/internal/todo/ent/friendship"
	"entgo.io/contrib/entgql/internal/todo/ent/group"
	"entgo.io/contrib/entgql/internal/todo/ent/todo"
	"entgo.io/contrib/entgql/internal/todo/ent/user"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *CategoryQuery) CollectFields(ctx context.Context, satisfies ...string) (*CategoryQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return c, nil
	}
	if err := c.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *CategoryQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	c = c.limitSelection(op, field, satisfies...)

	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "todos":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &TodoQuery{config: c.config}
			)
			args := newTodoPaginateArgs(fieldArgs(ctx, new(TodoWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newTodoPager(args.opts)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					c.loadTotal = append(c.loadTotal, func(ctx context.Context, nodes []*Category) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID int `sql:"category_id"`
							Count  int `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(category.TodosColumn, ids...))
						})
						if err := query.GroupBy(category.TodosColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[int]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				} else {
					c.loadTotal = append(c.loadTotal, func(_ context.Context, nodes []*Category) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Todos)
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}

			query = pager.applyCursors(query, args.after, args.before)
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(category.TodosColumn, limit, pager.orderExpr(args.last != nil))
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query, args.last != nil)
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, op, *field, path, satisfies...); err != nil {
					return err
				}
			}
			c.WithNamedTodos(alias, func(wq *TodoQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

func (c *CategoryQuery) limitSelection(op *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *CategoryQuery {
	selectFields := []string{category.FieldID}
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {

		case "text":
			selectFields = append(selectFields, category.FieldText)

		case "status":
			selectFields = append(selectFields, category.FieldStatus)

		case "config":
			selectFields = append(selectFields, category.FieldConfig)

		case "duration":
			selectFields = append(selectFields, category.FieldDuration)

		case "count":
			selectFields = append(selectFields, category.FieldCount)

		case "strings":
			selectFields = append(selectFields, category.FieldStrings)

		}
	}
	return c.Select(selectFields...).CategoryQuery
}

type categoryPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []CategoryPaginateOption
}

func newCategoryPaginateArgs(rv map[string]interface{}) *categoryPaginateArgs {
	args := &categoryPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &CategoryOrder{Field: &CategoryOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithCategoryOrder(order))
			}
		case *CategoryOrder:
			if v != nil {
				args.opts = append(args.opts, WithCategoryOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*CategoryWhereInput); ok {
		args.opts = append(args.opts, WithCategoryFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (f *FriendshipQuery) CollectFields(ctx context.Context, satisfies ...string) (*FriendshipQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return f, nil
	}
	if err := f.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return f, nil
}

func (f *FriendshipQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	f = f.limitSelection(op, field, satisfies...)

	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "user":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: f.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			f.withUser = query
		case "friend":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: f.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			f.withFriend = query
		}
	}
	return nil
}

func (f *FriendshipQuery) limitSelection(op *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *FriendshipQuery {
	selectFields := []string{friendship.FieldID}
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {

		case "createdAt":
			selectFields = append(selectFields, friendship.FieldCreatedAt)

		case "userID":
			selectFields = append(selectFields, friendship.FieldUserID)

		case "friendID":
			selectFields = append(selectFields, friendship.FieldFriendID)

		}
	}
	return f.Select(selectFields...).FriendshipQuery
}

type friendshipPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []FriendshipPaginateOption
}

func newFriendshipPaginateArgs(rv map[string]interface{}) *friendshipPaginateArgs {
	args := &friendshipPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*FriendshipWhereInput); ok {
		args.opts = append(args.opts, WithFriendshipFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (gr *GroupQuery) CollectFields(ctx context.Context, satisfies ...string) (*GroupQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return gr, nil
	}
	if err := gr.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return gr, nil
}

func (gr *GroupQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	gr = gr.limitSelection(op, field, satisfies...)

	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "users":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: gr.config}
			)
			args := newUserPaginateArgs(fieldArgs(ctx, new(UserWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newUserPager(args.opts)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					gr.loadTotal = append(gr.loadTotal, func(ctx context.Context, nodes []*Group) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID int `sql:"group_id"`
							Count  int `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							joinT := sql.Table(group.UsersTable)
							s.Join(joinT).On(s.C(user.FieldID), joinT.C(group.UsersPrimaryKey[0]))
							s.Where(sql.InValues(joinT.C(group.UsersPrimaryKey[1]), ids...))
							s.Select(joinT.C(group.UsersPrimaryKey[1]), sql.Count("*"))
							s.GroupBy(joinT.C(group.UsersPrimaryKey[1]))
						})
						if err := query.Select().Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[int]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				} else {
					gr.loadTotal = append(gr.loadTotal, func(_ context.Context, nodes []*Group) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Users)
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}

			query = pager.applyCursors(query, args.after, args.before)
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(group.UsersPrimaryKey[1], limit, pager.orderExpr(args.last != nil))
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query, args.last != nil)
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, op, *field, path, satisfies...); err != nil {
					return err
				}
			}
			gr.WithNamedUsers(alias, func(wq *UserQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

func (gr *GroupQuery) limitSelection(op *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *GroupQuery {
	selectFields := []string{group.FieldID}
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {

		case "name":
			selectFields = append(selectFields, group.FieldName)

		}
	}
	return gr.Select(selectFields...).GroupQuery
}

type groupPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []GroupPaginateOption
}

func newGroupPaginateArgs(rv map[string]interface{}) *groupPaginateArgs {
	args := &groupPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*GroupWhereInput); ok {
		args.opts = append(args.opts, WithGroupFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TodoQuery) CollectFields(ctx context.Context, satisfies ...string) (*TodoQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return t, nil
	}
	if err := t.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *TodoQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	t = t.limitSelection(op, field, satisfies...)

	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "parent":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &TodoQuery{config: t.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			t.withParent = query
		case "children":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &TodoQuery{config: t.config}
			)
			args := newTodoPaginateArgs(fieldArgs(ctx, new(TodoWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newTodoPager(args.opts)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					t.loadTotal = append(t.loadTotal, func(ctx context.Context, nodes []*Todo) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID int `sql:"todo_children"`
							Count  int `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(todo.ChildrenColumn, ids...))
						})
						if err := query.GroupBy(todo.ChildrenColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[int]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[1] == nil {
								nodes[i].Edges.totalCount[1] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[1][alias] = n
						}
						return nil
					})
				} else {
					t.loadTotal = append(t.loadTotal, func(_ context.Context, nodes []*Todo) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Children)
							if nodes[i].Edges.totalCount[1] == nil {
								nodes[i].Edges.totalCount[1] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[1][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}

			query = pager.applyCursors(query, args.after, args.before)
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(todo.ChildrenColumn, limit, pager.orderExpr(args.last != nil))
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query, args.last != nil)
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, op, *field, path, satisfies...); err != nil {
					return err
				}
			}
			t.WithNamedChildren(alias, func(wq *TodoQuery) {
				*wq = *query
			})
		case "category":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &CategoryQuery{config: t.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			t.withCategory = query
		}
	}
	return nil
}

func (t *TodoQuery) limitSelection(op *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *TodoQuery {
	selectFields := []string{todo.FieldID}
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {

		case "createdAt":
			selectFields = append(selectFields, todo.FieldCreatedAt)

		case "status":
			selectFields = append(selectFields, todo.FieldStatus)

		case "priorityOrder":
			selectFields = append(selectFields, todo.FieldPriority)

		case "text":
			selectFields = append(selectFields, todo.FieldText)

		case "categoryID", "category_id", "categoryX":
			selectFields = append(selectFields, todo.FieldCategoryID)

		case "init":
			selectFields = append(selectFields, todo.FieldInit)

		}
	}
	return t.Select(selectFields...).TodoQuery
}

type todoPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []TodoPaginateOption
}

func newTodoPaginateArgs(rv map[string]interface{}) *todoPaginateArgs {
	args := &todoPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &TodoOrder{Field: &TodoOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithTodoOrder(order))
			}
		case *TodoOrder:
			if v != nil {
				args.opts = append(args.opts, WithTodoOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*TodoWhereInput); ok {
		args.opts = append(args.opts, WithTodoFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	u = u.limitSelection(op, field, satisfies...)

	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "groups":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &GroupQuery{config: u.config}
			)
			args := newGroupPaginateArgs(fieldArgs(ctx, new(GroupWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newGroupPager(args.opts)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					u.loadTotal = append(u.loadTotal, func(ctx context.Context, nodes []*User) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID int `sql:"user_id"`
							Count  int `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							joinT := sql.Table(user.GroupsTable)
							s.Join(joinT).On(s.C(group.FieldID), joinT.C(user.GroupsPrimaryKey[1]))
							s.Where(sql.InValues(joinT.C(user.GroupsPrimaryKey[0]), ids...))
							s.Select(joinT.C(user.GroupsPrimaryKey[0]), sql.Count("*"))
							s.GroupBy(joinT.C(user.GroupsPrimaryKey[0]))
						})
						if err := query.Select().Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[int]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				} else {
					u.loadTotal = append(u.loadTotal, func(_ context.Context, nodes []*User) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Groups)
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}

			query = pager.applyCursors(query, args.after, args.before)
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(user.GroupsPrimaryKey[0], limit, pager.orderExpr(args.last != nil))
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query, args.last != nil)
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, op, *field, path, satisfies...); err != nil {
					return err
				}
			}
			u.WithNamedGroups(alias, func(wq *GroupQuery) {
				*wq = *query
			})
		case "friends":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: u.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedFriends(alias, func(wq *UserQuery) {
				*wq = *query
			})
		case "friendships":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &FriendshipQuery{config: u.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedFriendships(alias, func(wq *FriendshipQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

func (u *UserQuery) limitSelection(op *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *UserQuery {
	selectFields := []string{user.FieldID}
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {

		case "name":
			selectFields = append(selectFields, user.FieldName)

		case "password":
			selectFields = append(selectFields, user.FieldPassword)

		}
	}
	return u.Select(selectFields...).UserQuery
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]interface{}) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*UserWhereInput); ok {
		args.opts = append(args.opts, WithUserFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput interface{}, path ...string) map[string]interface{} {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	for _, name := range path {
		var field *graphql.CollectedField
		for _, f := range graphql.CollectFields(oc, fc.Field.Selections, nil) {
			if f.Alias == name {
				field = &f
				break
			}
		}
		if field == nil {
			return nil
		}
		cf, err := fc.Child(ctx, *field)
		if err != nil {
			args := field.ArgumentMap(oc.Variables)
			return unmarshalArgs(ctx, whereInput, args)
		}
		fc = cf
	}
	return fc.Args
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput interface{}, args map[string]interface{}) map[string]interface{} {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

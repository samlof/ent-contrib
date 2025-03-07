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
	"errors"
	"fmt"
	"log"

	"entgo.io/contrib/entgql/internal/todogotype/ent/migrate"
	"entgo.io/contrib/entgql/internal/todogotype/ent/schema/bigintgql"
	"entgo.io/contrib/entgql/internal/todogotype/ent/schema/uintgql"

	"entgo.io/contrib/entgql/internal/todogotype/ent/category"
	"entgo.io/contrib/entgql/internal/todogotype/ent/friendship"
	"entgo.io/contrib/entgql/internal/todogotype/ent/group"
	"entgo.io/contrib/entgql/internal/todogotype/ent/pet"
	"entgo.io/contrib/entgql/internal/todogotype/ent/todo"
	"entgo.io/contrib/entgql/internal/todogotype/ent/user"
	"entgo.io/contrib/entgql/internal/todogotype/ent/verysecret"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Category is the client for interacting with the Category builders.
	Category *CategoryClient
	// Friendship is the client for interacting with the Friendship builders.
	Friendship *FriendshipClient
	// Group is the client for interacting with the Group builders.
	Group *GroupClient
	// Pet is the client for interacting with the Pet builders.
	Pet *PetClient
	// Todo is the client for interacting with the Todo builders.
	Todo *TodoClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// VerySecret is the client for interacting with the VerySecret builders.
	VerySecret *VerySecretClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Category = NewCategoryClient(c.config)
	c.Friendship = NewFriendshipClient(c.config)
	c.Group = NewGroupClient(c.config)
	c.Pet = NewPetClient(c.config)
	c.Todo = NewTodoClient(c.config)
	c.User = NewUserClient(c.config)
	c.VerySecret = NewVerySecretClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Category:   NewCategoryClient(cfg),
		Friendship: NewFriendshipClient(cfg),
		Group:      NewGroupClient(cfg),
		Pet:        NewPetClient(cfg),
		Todo:       NewTodoClient(cfg),
		User:       NewUserClient(cfg),
		VerySecret: NewVerySecretClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Category:   NewCategoryClient(cfg),
		Friendship: NewFriendshipClient(cfg),
		Group:      NewGroupClient(cfg),
		Pet:        NewPetClient(cfg),
		Todo:       NewTodoClient(cfg),
		User:       NewUserClient(cfg),
		VerySecret: NewVerySecretClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Category.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Category.Use(hooks...)
	c.Friendship.Use(hooks...)
	c.Group.Use(hooks...)
	c.Pet.Use(hooks...)
	c.Todo.Use(hooks...)
	c.User.Use(hooks...)
	c.VerySecret.Use(hooks...)
}

// CategoryClient is a client for the Category schema.
type CategoryClient struct {
	config
}

// NewCategoryClient returns a client for the Category from the given config.
func NewCategoryClient(c config) *CategoryClient {
	return &CategoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `category.Hooks(f(g(h())))`.
func (c *CategoryClient) Use(hooks ...Hook) {
	c.hooks.Category = append(c.hooks.Category, hooks...)
}

// Create returns a builder for creating a Category entity.
func (c *CategoryClient) Create() *CategoryCreate {
	mutation := newCategoryMutation(c.config, OpCreate)
	return &CategoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Category entities.
func (c *CategoryClient) CreateBulk(builders ...*CategoryCreate) *CategoryCreateBulk {
	return &CategoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Category.
func (c *CategoryClient) Update() *CategoryUpdate {
	mutation := newCategoryMutation(c.config, OpUpdate)
	return &CategoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CategoryClient) UpdateOne(ca *Category) *CategoryUpdateOne {
	mutation := newCategoryMutation(c.config, OpUpdateOne, withCategory(ca))
	return &CategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CategoryClient) UpdateOneID(id bigintgql.BigInt) *CategoryUpdateOne {
	mutation := newCategoryMutation(c.config, OpUpdateOne, withCategoryID(id))
	return &CategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Category.
func (c *CategoryClient) Delete() *CategoryDelete {
	mutation := newCategoryMutation(c.config, OpDelete)
	return &CategoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CategoryClient) DeleteOne(ca *Category) *CategoryDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *CategoryClient) DeleteOneID(id bigintgql.BigInt) *CategoryDeleteOne {
	builder := c.Delete().Where(category.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CategoryDeleteOne{builder}
}

// Query returns a query builder for Category.
func (c *CategoryClient) Query() *CategoryQuery {
	return &CategoryQuery{
		config: c.config,
	}
}

// Get returns a Category entity by its id.
func (c *CategoryClient) Get(ctx context.Context, id bigintgql.BigInt) (*Category, error) {
	return c.Query().Where(category.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CategoryClient) GetX(ctx context.Context, id bigintgql.BigInt) *Category {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTodos queries the todos edge of a Category.
func (c *CategoryClient) QueryTodos(ca *Category) *TodoQuery {
	query := &TodoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(category.Table, category.FieldID, id),
			sqlgraph.To(todo.Table, todo.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, category.TodosTable, category.TodosColumn),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CategoryClient) Hooks() []Hook {
	return c.hooks.Category
}

// FriendshipClient is a client for the Friendship schema.
type FriendshipClient struct {
	config
}

// NewFriendshipClient returns a client for the Friendship from the given config.
func NewFriendshipClient(c config) *FriendshipClient {
	return &FriendshipClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `friendship.Hooks(f(g(h())))`.
func (c *FriendshipClient) Use(hooks ...Hook) {
	c.hooks.Friendship = append(c.hooks.Friendship, hooks...)
}

// Create returns a builder for creating a Friendship entity.
func (c *FriendshipClient) Create() *FriendshipCreate {
	mutation := newFriendshipMutation(c.config, OpCreate)
	return &FriendshipCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Friendship entities.
func (c *FriendshipClient) CreateBulk(builders ...*FriendshipCreate) *FriendshipCreateBulk {
	return &FriendshipCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Friendship.
func (c *FriendshipClient) Update() *FriendshipUpdate {
	mutation := newFriendshipMutation(c.config, OpUpdate)
	return &FriendshipUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FriendshipClient) UpdateOne(f *Friendship) *FriendshipUpdateOne {
	mutation := newFriendshipMutation(c.config, OpUpdateOne, withFriendship(f))
	return &FriendshipUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FriendshipClient) UpdateOneID(id string) *FriendshipUpdateOne {
	mutation := newFriendshipMutation(c.config, OpUpdateOne, withFriendshipID(id))
	return &FriendshipUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Friendship.
func (c *FriendshipClient) Delete() *FriendshipDelete {
	mutation := newFriendshipMutation(c.config, OpDelete)
	return &FriendshipDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *FriendshipClient) DeleteOne(f *Friendship) *FriendshipDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *FriendshipClient) DeleteOneID(id string) *FriendshipDeleteOne {
	builder := c.Delete().Where(friendship.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FriendshipDeleteOne{builder}
}

// Query returns a query builder for Friendship.
func (c *FriendshipClient) Query() *FriendshipQuery {
	return &FriendshipQuery{
		config: c.config,
	}
}

// Get returns a Friendship entity by its id.
func (c *FriendshipClient) Get(ctx context.Context, id string) (*Friendship, error) {
	return c.Query().Where(friendship.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FriendshipClient) GetX(ctx context.Context, id string) *Friendship {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Friendship.
func (c *FriendshipClient) QueryUser(f *Friendship) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(friendship.Table, friendship.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, friendship.UserTable, friendship.UserColumn),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFriend queries the friend edge of a Friendship.
func (c *FriendshipClient) QueryFriend(f *Friendship) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(friendship.Table, friendship.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, friendship.FriendTable, friendship.FriendColumn),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FriendshipClient) Hooks() []Hook {
	return c.hooks.Friendship
}

// GroupClient is a client for the Group schema.
type GroupClient struct {
	config
}

// NewGroupClient returns a client for the Group from the given config.
func NewGroupClient(c config) *GroupClient {
	return &GroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `group.Hooks(f(g(h())))`.
func (c *GroupClient) Use(hooks ...Hook) {
	c.hooks.Group = append(c.hooks.Group, hooks...)
}

// Create returns a builder for creating a Group entity.
func (c *GroupClient) Create() *GroupCreate {
	mutation := newGroupMutation(c.config, OpCreate)
	return &GroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Group entities.
func (c *GroupClient) CreateBulk(builders ...*GroupCreate) *GroupCreateBulk {
	return &GroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Group.
func (c *GroupClient) Update() *GroupUpdate {
	mutation := newGroupMutation(c.config, OpUpdate)
	return &GroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GroupClient) UpdateOne(gr *Group) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroup(gr))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupClient) UpdateOneID(id string) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroupID(id))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Group.
func (c *GroupClient) Delete() *GroupDelete {
	mutation := newGroupMutation(c.config, OpDelete)
	return &GroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GroupClient) DeleteOne(gr *Group) *GroupDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *GroupClient) DeleteOneID(id string) *GroupDeleteOne {
	builder := c.Delete().Where(group.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GroupDeleteOne{builder}
}

// Query returns a query builder for Group.
func (c *GroupClient) Query() *GroupQuery {
	return &GroupQuery{
		config: c.config,
	}
}

// Get returns a Group entity by its id.
func (c *GroupClient) Get(ctx context.Context, id string) (*Group, error) {
	return c.Query().Where(group.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupClient) GetX(ctx context.Context, id string) *Group {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a Group.
func (c *GroupClient) QueryUsers(gr *Group) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := gr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(group.Table, group.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, group.UsersTable, group.UsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(gr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GroupClient) Hooks() []Hook {
	return c.hooks.Group
}

// PetClient is a client for the Pet schema.
type PetClient struct {
	config
}

// NewPetClient returns a client for the Pet from the given config.
func NewPetClient(c config) *PetClient {
	return &PetClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `pet.Hooks(f(g(h())))`.
func (c *PetClient) Use(hooks ...Hook) {
	c.hooks.Pet = append(c.hooks.Pet, hooks...)
}

// Create returns a builder for creating a Pet entity.
func (c *PetClient) Create() *PetCreate {
	mutation := newPetMutation(c.config, OpCreate)
	return &PetCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Pet entities.
func (c *PetClient) CreateBulk(builders ...*PetCreate) *PetCreateBulk {
	return &PetCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Pet.
func (c *PetClient) Update() *PetUpdate {
	mutation := newPetMutation(c.config, OpUpdate)
	return &PetUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PetClient) UpdateOne(pe *Pet) *PetUpdateOne {
	mutation := newPetMutation(c.config, OpUpdateOne, withPet(pe))
	return &PetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PetClient) UpdateOneID(id uintgql.Uint64) *PetUpdateOne {
	mutation := newPetMutation(c.config, OpUpdateOne, withPetID(id))
	return &PetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Pet.
func (c *PetClient) Delete() *PetDelete {
	mutation := newPetMutation(c.config, OpDelete)
	return &PetDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PetClient) DeleteOne(pe *Pet) *PetDeleteOne {
	return c.DeleteOneID(pe.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *PetClient) DeleteOneID(id uintgql.Uint64) *PetDeleteOne {
	builder := c.Delete().Where(pet.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PetDeleteOne{builder}
}

// Query returns a query builder for Pet.
func (c *PetClient) Query() *PetQuery {
	return &PetQuery{
		config: c.config,
	}
}

// Get returns a Pet entity by its id.
func (c *PetClient) Get(ctx context.Context, id uintgql.Uint64) (*Pet, error) {
	return c.Query().Where(pet.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PetClient) GetX(ctx context.Context, id uintgql.Uint64) *Pet {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PetClient) Hooks() []Hook {
	return c.hooks.Pet
}

// TodoClient is a client for the Todo schema.
type TodoClient struct {
	config
}

// NewTodoClient returns a client for the Todo from the given config.
func NewTodoClient(c config) *TodoClient {
	return &TodoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `todo.Hooks(f(g(h())))`.
func (c *TodoClient) Use(hooks ...Hook) {
	c.hooks.Todo = append(c.hooks.Todo, hooks...)
}

// Create returns a builder for creating a Todo entity.
func (c *TodoClient) Create() *TodoCreate {
	mutation := newTodoMutation(c.config, OpCreate)
	return &TodoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Todo entities.
func (c *TodoClient) CreateBulk(builders ...*TodoCreate) *TodoCreateBulk {
	return &TodoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Todo.
func (c *TodoClient) Update() *TodoUpdate {
	mutation := newTodoMutation(c.config, OpUpdate)
	return &TodoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TodoClient) UpdateOne(t *Todo) *TodoUpdateOne {
	mutation := newTodoMutation(c.config, OpUpdateOne, withTodo(t))
	return &TodoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TodoClient) UpdateOneID(id string) *TodoUpdateOne {
	mutation := newTodoMutation(c.config, OpUpdateOne, withTodoID(id))
	return &TodoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Todo.
func (c *TodoClient) Delete() *TodoDelete {
	mutation := newTodoMutation(c.config, OpDelete)
	return &TodoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TodoClient) DeleteOne(t *Todo) *TodoDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *TodoClient) DeleteOneID(id string) *TodoDeleteOne {
	builder := c.Delete().Where(todo.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TodoDeleteOne{builder}
}

// Query returns a query builder for Todo.
func (c *TodoClient) Query() *TodoQuery {
	return &TodoQuery{
		config: c.config,
	}
}

// Get returns a Todo entity by its id.
func (c *TodoClient) Get(ctx context.Context, id string) (*Todo, error) {
	return c.Query().Where(todo.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TodoClient) GetX(ctx context.Context, id string) *Todo {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a Todo.
func (c *TodoClient) QueryParent(t *Todo) *TodoQuery {
	query := &TodoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(todo.Table, todo.FieldID, id),
			sqlgraph.To(todo.Table, todo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, todo.ParentTable, todo.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChildren queries the children edge of a Todo.
func (c *TodoClient) QueryChildren(t *Todo) *TodoQuery {
	query := &TodoQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(todo.Table, todo.FieldID, id),
			sqlgraph.To(todo.Table, todo.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, todo.ChildrenTable, todo.ChildrenColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCategory queries the category edge of a Todo.
func (c *TodoClient) QueryCategory(t *Todo) *CategoryQuery {
	query := &CategoryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(todo.Table, todo.FieldID, id),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, todo.CategoryTable, todo.CategoryColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySecret queries the secret edge of a Todo.
func (c *TodoClient) QuerySecret(t *Todo) *VerySecretQuery {
	query := &VerySecretQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(todo.Table, todo.FieldID, id),
			sqlgraph.To(verysecret.Table, verysecret.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, todo.SecretTable, todo.SecretColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TodoClient) Hooks() []Hook {
	return c.hooks.Todo
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id string) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id string) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id string) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id string) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryGroups queries the groups edge of a User.
func (c *UserClient) QueryGroups(u *User) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.GroupsTable, user.GroupsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFriends queries the friends edge of a User.
func (c *UserClient) QueryFriends(u *User) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.FriendsTable, user.FriendsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFriendships queries the friendships edge of a User.
func (c *UserClient) QueryFriendships(u *User) *FriendshipQuery {
	query := &FriendshipQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(friendship.Table, friendship.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, user.FriendshipsTable, user.FriendshipsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// VerySecretClient is a client for the VerySecret schema.
type VerySecretClient struct {
	config
}

// NewVerySecretClient returns a client for the VerySecret from the given config.
func NewVerySecretClient(c config) *VerySecretClient {
	return &VerySecretClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `verysecret.Hooks(f(g(h())))`.
func (c *VerySecretClient) Use(hooks ...Hook) {
	c.hooks.VerySecret = append(c.hooks.VerySecret, hooks...)
}

// Create returns a builder for creating a VerySecret entity.
func (c *VerySecretClient) Create() *VerySecretCreate {
	mutation := newVerySecretMutation(c.config, OpCreate)
	return &VerySecretCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of VerySecret entities.
func (c *VerySecretClient) CreateBulk(builders ...*VerySecretCreate) *VerySecretCreateBulk {
	return &VerySecretCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for VerySecret.
func (c *VerySecretClient) Update() *VerySecretUpdate {
	mutation := newVerySecretMutation(c.config, OpUpdate)
	return &VerySecretUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VerySecretClient) UpdateOne(vs *VerySecret) *VerySecretUpdateOne {
	mutation := newVerySecretMutation(c.config, OpUpdateOne, withVerySecret(vs))
	return &VerySecretUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VerySecretClient) UpdateOneID(id string) *VerySecretUpdateOne {
	mutation := newVerySecretMutation(c.config, OpUpdateOne, withVerySecretID(id))
	return &VerySecretUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for VerySecret.
func (c *VerySecretClient) Delete() *VerySecretDelete {
	mutation := newVerySecretMutation(c.config, OpDelete)
	return &VerySecretDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *VerySecretClient) DeleteOne(vs *VerySecret) *VerySecretDeleteOne {
	return c.DeleteOneID(vs.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *VerySecretClient) DeleteOneID(id string) *VerySecretDeleteOne {
	builder := c.Delete().Where(verysecret.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VerySecretDeleteOne{builder}
}

// Query returns a query builder for VerySecret.
func (c *VerySecretClient) Query() *VerySecretQuery {
	return &VerySecretQuery{
		config: c.config,
	}
}

// Get returns a VerySecret entity by its id.
func (c *VerySecretClient) Get(ctx context.Context, id string) (*VerySecret, error) {
	return c.Query().Where(verysecret.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VerySecretClient) GetX(ctx context.Context, id string) *VerySecret {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *VerySecretClient) Hooks() []Hook {
	return c.hooks.VerySecret
}

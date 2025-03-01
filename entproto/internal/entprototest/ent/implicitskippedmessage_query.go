// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/contrib/entproto/internal/entprototest/ent/implicitskippedmessage"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ImplicitSkippedMessageQuery is the builder for querying ImplicitSkippedMessage entities.
type ImplicitSkippedMessageQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ImplicitSkippedMessage
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ImplicitSkippedMessageQuery builder.
func (ismq *ImplicitSkippedMessageQuery) Where(ps ...predicate.ImplicitSkippedMessage) *ImplicitSkippedMessageQuery {
	ismq.predicates = append(ismq.predicates, ps...)
	return ismq
}

// Limit adds a limit step to the query.
func (ismq *ImplicitSkippedMessageQuery) Limit(limit int) *ImplicitSkippedMessageQuery {
	ismq.limit = &limit
	return ismq
}

// Offset adds an offset step to the query.
func (ismq *ImplicitSkippedMessageQuery) Offset(offset int) *ImplicitSkippedMessageQuery {
	ismq.offset = &offset
	return ismq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ismq *ImplicitSkippedMessageQuery) Unique(unique bool) *ImplicitSkippedMessageQuery {
	ismq.unique = &unique
	return ismq
}

// Order adds an order step to the query.
func (ismq *ImplicitSkippedMessageQuery) Order(o ...OrderFunc) *ImplicitSkippedMessageQuery {
	ismq.order = append(ismq.order, o...)
	return ismq
}

// First returns the first ImplicitSkippedMessage entity from the query.
// Returns a *NotFoundError when no ImplicitSkippedMessage was found.
func (ismq *ImplicitSkippedMessageQuery) First(ctx context.Context) (*ImplicitSkippedMessage, error) {
	nodes, err := ismq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{implicitskippedmessage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ismq *ImplicitSkippedMessageQuery) FirstX(ctx context.Context) *ImplicitSkippedMessage {
	node, err := ismq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ImplicitSkippedMessage ID from the query.
// Returns a *NotFoundError when no ImplicitSkippedMessage ID was found.
func (ismq *ImplicitSkippedMessageQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ismq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{implicitskippedmessage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ismq *ImplicitSkippedMessageQuery) FirstIDX(ctx context.Context) int {
	id, err := ismq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ImplicitSkippedMessage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ImplicitSkippedMessage entity is found.
// Returns a *NotFoundError when no ImplicitSkippedMessage entities are found.
func (ismq *ImplicitSkippedMessageQuery) Only(ctx context.Context) (*ImplicitSkippedMessage, error) {
	nodes, err := ismq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{implicitskippedmessage.Label}
	default:
		return nil, &NotSingularError{implicitskippedmessage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ismq *ImplicitSkippedMessageQuery) OnlyX(ctx context.Context) *ImplicitSkippedMessage {
	node, err := ismq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ImplicitSkippedMessage ID in the query.
// Returns a *NotSingularError when more than one ImplicitSkippedMessage ID is found.
// Returns a *NotFoundError when no entities are found.
func (ismq *ImplicitSkippedMessageQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ismq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{implicitskippedmessage.Label}
	default:
		err = &NotSingularError{implicitskippedmessage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ismq *ImplicitSkippedMessageQuery) OnlyIDX(ctx context.Context) int {
	id, err := ismq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ImplicitSkippedMessages.
func (ismq *ImplicitSkippedMessageQuery) All(ctx context.Context) ([]*ImplicitSkippedMessage, error) {
	if err := ismq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ismq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ismq *ImplicitSkippedMessageQuery) AllX(ctx context.Context) []*ImplicitSkippedMessage {
	nodes, err := ismq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ImplicitSkippedMessage IDs.
func (ismq *ImplicitSkippedMessageQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ismq.Select(implicitskippedmessage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ismq *ImplicitSkippedMessageQuery) IDsX(ctx context.Context) []int {
	ids, err := ismq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ismq *ImplicitSkippedMessageQuery) Count(ctx context.Context) (int, error) {
	if err := ismq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ismq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ismq *ImplicitSkippedMessageQuery) CountX(ctx context.Context) int {
	count, err := ismq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ismq *ImplicitSkippedMessageQuery) Exist(ctx context.Context) (bool, error) {
	if err := ismq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ismq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ismq *ImplicitSkippedMessageQuery) ExistX(ctx context.Context) bool {
	exist, err := ismq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ImplicitSkippedMessageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ismq *ImplicitSkippedMessageQuery) Clone() *ImplicitSkippedMessageQuery {
	if ismq == nil {
		return nil
	}
	return &ImplicitSkippedMessageQuery{
		config:     ismq.config,
		limit:      ismq.limit,
		offset:     ismq.offset,
		order:      append([]OrderFunc{}, ismq.order...),
		predicates: append([]predicate.ImplicitSkippedMessage{}, ismq.predicates...),
		// clone intermediate query.
		sql:    ismq.sql.Clone(),
		path:   ismq.path,
		unique: ismq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (ismq *ImplicitSkippedMessageQuery) GroupBy(field string, fields ...string) *ImplicitSkippedMessageGroupBy {
	grbuild := &ImplicitSkippedMessageGroupBy{config: ismq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ismq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ismq.sqlQuery(ctx), nil
	}
	grbuild.label = implicitskippedmessage.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (ismq *ImplicitSkippedMessageQuery) Select(fields ...string) *ImplicitSkippedMessageSelect {
	ismq.fields = append(ismq.fields, fields...)
	selbuild := &ImplicitSkippedMessageSelect{ImplicitSkippedMessageQuery: ismq}
	selbuild.label = implicitskippedmessage.Label
	selbuild.flds, selbuild.scan = &ismq.fields, selbuild.Scan
	return selbuild
}

func (ismq *ImplicitSkippedMessageQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ismq.fields {
		if !implicitskippedmessage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ismq.path != nil {
		prev, err := ismq.path(ctx)
		if err != nil {
			return err
		}
		ismq.sql = prev
	}
	return nil
}

func (ismq *ImplicitSkippedMessageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ImplicitSkippedMessage, error) {
	var (
		nodes   = []*ImplicitSkippedMessage{}
		withFKs = ismq.withFKs
		_spec   = ismq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, implicitskippedmessage.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ImplicitSkippedMessage).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ImplicitSkippedMessage{config: ismq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ismq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ismq *ImplicitSkippedMessageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ismq.querySpec()
	_spec.Node.Columns = ismq.fields
	if len(ismq.fields) > 0 {
		_spec.Unique = ismq.unique != nil && *ismq.unique
	}
	return sqlgraph.CountNodes(ctx, ismq.driver, _spec)
}

func (ismq *ImplicitSkippedMessageQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := ismq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (ismq *ImplicitSkippedMessageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   implicitskippedmessage.Table,
			Columns: implicitskippedmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: implicitskippedmessage.FieldID,
			},
		},
		From:   ismq.sql,
		Unique: true,
	}
	if unique := ismq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ismq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, implicitskippedmessage.FieldID)
		for i := range fields {
			if fields[i] != implicitskippedmessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ismq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ismq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ismq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ismq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ismq *ImplicitSkippedMessageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ismq.driver.Dialect())
	t1 := builder.Table(implicitskippedmessage.Table)
	columns := ismq.fields
	if len(columns) == 0 {
		columns = implicitskippedmessage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ismq.sql != nil {
		selector = ismq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ismq.unique != nil && *ismq.unique {
		selector.Distinct()
	}
	for _, p := range ismq.predicates {
		p(selector)
	}
	for _, p := range ismq.order {
		p(selector)
	}
	if offset := ismq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ismq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ImplicitSkippedMessageGroupBy is the group-by builder for ImplicitSkippedMessage entities.
type ImplicitSkippedMessageGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ismgb *ImplicitSkippedMessageGroupBy) Aggregate(fns ...AggregateFunc) *ImplicitSkippedMessageGroupBy {
	ismgb.fns = append(ismgb.fns, fns...)
	return ismgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ismgb *ImplicitSkippedMessageGroupBy) Scan(ctx context.Context, v any) error {
	query, err := ismgb.path(ctx)
	if err != nil {
		return err
	}
	ismgb.sql = query
	return ismgb.sqlScan(ctx, v)
}

func (ismgb *ImplicitSkippedMessageGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range ismgb.fields {
		if !implicitskippedmessage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ismgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ismgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ismgb *ImplicitSkippedMessageGroupBy) sqlQuery() *sql.Selector {
	selector := ismgb.sql.Select()
	aggregation := make([]string, 0, len(ismgb.fns))
	for _, fn := range ismgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ismgb.fields)+len(ismgb.fns))
		for _, f := range ismgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ismgb.fields...)...)
}

// ImplicitSkippedMessageSelect is the builder for selecting fields of ImplicitSkippedMessage entities.
type ImplicitSkippedMessageSelect struct {
	*ImplicitSkippedMessageQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (isms *ImplicitSkippedMessageSelect) Scan(ctx context.Context, v any) error {
	if err := isms.prepareQuery(ctx); err != nil {
		return err
	}
	isms.sql = isms.ImplicitSkippedMessageQuery.sqlQuery(ctx)
	return isms.sqlScan(ctx, v)
}

func (isms *ImplicitSkippedMessageSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := isms.sql.Query()
	if err := isms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

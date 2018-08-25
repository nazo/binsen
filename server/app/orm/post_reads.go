// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package orm

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// PostRead is an object representing the database table.
type PostRead struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	PostID    int64     `boil:"post_id" json:"post_id" toml:"post_id" yaml:"post_id"`
	UserID    int64     `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *postReadR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L postReadL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PostReadColumns = struct {
	ID        string
	PostID    string
	UserID    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	PostID:    "post_id",
	UserID:    "user_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// PostReadRels is where relationship names are stored.
var PostReadRels = struct {
	Post string
	User string
}{
	Post: "Post",
	User: "User",
}

// postReadR is where relationships are stored.
type postReadR struct {
	Post *Post
	User *User
}

// NewStruct creates a new relationship struct
func (*postReadR) NewStruct() *postReadR {
	return &postReadR{}
}

// postReadL is where Load methods for each relationship are stored.
type postReadL struct{}

var (
	postReadColumns               = []string{"id", "post_id", "user_id", "created_at", "updated_at"}
	postReadColumnsWithoutDefault = []string{"post_id", "user_id", "created_at", "updated_at"}
	postReadColumnsWithDefault    = []string{"id"}
	postReadPrimaryKeyColumns     = []string{"id"}
)

type (
	// PostReadSlice is an alias for a slice of pointers to PostRead.
	// This should generally be used opposed to []PostRead.
	PostReadSlice []*PostRead
	// PostReadHook is the signature for custom PostRead hook methods
	PostReadHook func(context.Context, boil.ContextExecutor, *PostRead) error

	postReadQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	postReadType                 = reflect.TypeOf(&PostRead{})
	postReadMapping              = queries.MakeStructMapping(postReadType)
	postReadPrimaryKeyMapping, _ = queries.BindMapping(postReadType, postReadMapping, postReadPrimaryKeyColumns)
	postReadInsertCacheMut       sync.RWMutex
	postReadInsertCache          = make(map[string]insertCache)
	postReadUpdateCacheMut       sync.RWMutex
	postReadUpdateCache          = make(map[string]updateCache)
	postReadUpsertCacheMut       sync.RWMutex
	postReadUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var postReadBeforeInsertHooks []PostReadHook
var postReadBeforeUpdateHooks []PostReadHook
var postReadBeforeDeleteHooks []PostReadHook
var postReadBeforeUpsertHooks []PostReadHook

var postReadAfterInsertHooks []PostReadHook
var postReadAfterSelectHooks []PostReadHook
var postReadAfterUpdateHooks []PostReadHook
var postReadAfterDeleteHooks []PostReadHook
var postReadAfterUpsertHooks []PostReadHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PostRead) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range postReadBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PostRead) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range postReadBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PostRead) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range postReadBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PostRead) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range postReadBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PostRead) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range postReadAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PostRead) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range postReadAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PostRead) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range postReadAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PostRead) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range postReadAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PostRead) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range postReadAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPostReadHook registers your hook function for all future operations.
func AddPostReadHook(hookPoint boil.HookPoint, postReadHook PostReadHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		postReadBeforeInsertHooks = append(postReadBeforeInsertHooks, postReadHook)
	case boil.BeforeUpdateHook:
		postReadBeforeUpdateHooks = append(postReadBeforeUpdateHooks, postReadHook)
	case boil.BeforeDeleteHook:
		postReadBeforeDeleteHooks = append(postReadBeforeDeleteHooks, postReadHook)
	case boil.BeforeUpsertHook:
		postReadBeforeUpsertHooks = append(postReadBeforeUpsertHooks, postReadHook)
	case boil.AfterInsertHook:
		postReadAfterInsertHooks = append(postReadAfterInsertHooks, postReadHook)
	case boil.AfterSelectHook:
		postReadAfterSelectHooks = append(postReadAfterSelectHooks, postReadHook)
	case boil.AfterUpdateHook:
		postReadAfterUpdateHooks = append(postReadAfterUpdateHooks, postReadHook)
	case boil.AfterDeleteHook:
		postReadAfterDeleteHooks = append(postReadAfterDeleteHooks, postReadHook)
	case boil.AfterUpsertHook:
		postReadAfterUpsertHooks = append(postReadAfterUpsertHooks, postReadHook)
	}
}

// One returns a single postRead record from the query.
func (q postReadQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PostRead, error) {
	o := &PostRead{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: failed to execute a one query for post_reads")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PostRead records from the query.
func (q postReadQuery) All(ctx context.Context, exec boil.ContextExecutor) (PostReadSlice, error) {
	var o []*PostRead

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "orm: failed to assign all query results to PostRead slice")
	}

	if len(postReadAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PostRead records in the query.
func (q postReadQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to count post_reads rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q postReadQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "orm: failed to check if post_reads exists")
	}

	return count > 0, nil
}

// Post pointed to by the foreign key.
func (o *PostRead) Post(mods ...qm.QueryMod) postQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.PostID),
	}

	queryMods = append(queryMods, mods...)

	query := Posts(queryMods...)
	queries.SetFrom(query.Query, "\"posts\"")

	return query
}

// User pointed to by the foreign key.
func (o *PostRead) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// LoadPost allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (postReadL) LoadPost(ctx context.Context, e boil.ContextExecutor, singular bool, maybePostRead interface{}, mods queries.Applicator) error {
	var slice []*PostRead
	var object *PostRead

	if singular {
		object = maybePostRead.(*PostRead)
	} else {
		slice = *maybePostRead.(*[]*PostRead)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &postReadR{}
		}
		args = append(args, object.PostID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &postReadR{}
			}

			for _, a := range args {
				if a == obj.PostID {
					continue Outer
				}
			}

			args = append(args, obj.PostID)
		}
	}

	query := NewQuery(qm.From(`posts`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Post")
	}

	var resultSlice []*Post
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Post")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for posts")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for posts")
	}

	if len(postReadAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Post = foreign
		if foreign.R == nil {
			foreign.R = &postR{}
		}
		foreign.R.PostReads = append(foreign.R.PostReads, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PostID == foreign.ID {
				local.R.Post = foreign
				if foreign.R == nil {
					foreign.R = &postR{}
				}
				foreign.R.PostReads = append(foreign.R.PostReads, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (postReadL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybePostRead interface{}, mods queries.Applicator) error {
	var slice []*PostRead
	var object *PostRead

	if singular {
		object = maybePostRead.(*PostRead)
	} else {
		slice = *maybePostRead.(*[]*PostRead)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &postReadR{}
		}
		args = append(args, object.UserID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &postReadR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)
		}
	}

	query := NewQuery(qm.From(`users`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(postReadAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.PostReads = append(foreign.R.PostReads, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.PostReads = append(foreign.R.PostReads, local)
				break
			}
		}
	}

	return nil
}

// SetPost of the postRead to the related item.
// Sets o.R.Post to related.
// Adds o to related.R.PostReads.
func (o *PostRead) SetPost(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Post) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"post_reads\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"post_id"}),
		strmangle.WhereClause("\"", "\"", 2, postReadPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PostID = related.ID
	if o.R == nil {
		o.R = &postReadR{
			Post: related,
		}
	} else {
		o.R.Post = related
	}

	if related.R == nil {
		related.R = &postR{
			PostReads: PostReadSlice{o},
		}
	} else {
		related.R.PostReads = append(related.R.PostReads, o)
	}

	return nil
}

// SetUser of the postRead to the related item.
// Sets o.R.User to related.
// Adds o to related.R.PostReads.
func (o *PostRead) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"post_reads\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, postReadPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &postReadR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			PostReads: PostReadSlice{o},
		}
	} else {
		related.R.PostReads = append(related.R.PostReads, o)
	}

	return nil
}

// PostReads retrieves all the records using an executor.
func PostReads(mods ...qm.QueryMod) postReadQuery {
	mods = append(mods, qm.From("\"post_reads\""))
	return postReadQuery{NewQuery(mods...)}
}

// FindPostRead retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPostRead(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*PostRead, error) {
	postReadObj := &PostRead{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"post_reads\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, postReadObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: unable to select from post_reads")
	}

	return postReadObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PostRead) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no post_reads provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(postReadColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	postReadInsertCacheMut.RLock()
	cache, cached := postReadInsertCache[key]
	postReadInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			postReadColumns,
			postReadColumnsWithDefault,
			postReadColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(postReadType, postReadMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(postReadType, postReadMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"post_reads\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"post_reads\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "orm: unable to insert into post_reads")
	}

	if !cached {
		postReadInsertCacheMut.Lock()
		postReadInsertCache[key] = cache
		postReadInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the PostRead.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PostRead) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	postReadUpdateCacheMut.RLock()
	cache, cached := postReadUpdateCache[key]
	postReadUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			postReadColumns,
			postReadPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("orm: unable to update post_reads, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"post_reads\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, postReadPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(postReadType, postReadMapping, append(wl, postReadPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update post_reads row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by update for post_reads")
	}

	if !cached {
		postReadUpdateCacheMut.Lock()
		postReadUpdateCache[key] = cache
		postReadUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q postReadQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all for post_reads")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected for post_reads")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PostReadSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("orm: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), postReadPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"post_reads\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, postReadPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all in postRead slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected all in update all postRead")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PostRead) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no post_reads provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(postReadColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	postReadUpsertCacheMut.RLock()
	cache, cached := postReadUpsertCache[key]
	postReadUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			postReadColumns,
			postReadColumnsWithDefault,
			postReadColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			postReadColumns,
			postReadPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("orm: unable to upsert post_reads, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(postReadPrimaryKeyColumns))
			copy(conflict, postReadPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"post_reads\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(postReadType, postReadMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(postReadType, postReadMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "orm: unable to upsert post_reads")
	}

	if !cached {
		postReadUpsertCacheMut.Lock()
		postReadUpsertCache[key] = cache
		postReadUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single PostRead record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PostRead) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("orm: no PostRead provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), postReadPrimaryKeyMapping)
	sql := "DELETE FROM \"post_reads\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete from post_reads")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by delete for post_reads")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q postReadQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("orm: no postReadQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from post_reads")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for post_reads")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PostReadSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("orm: no PostRead slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(postReadBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), postReadPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"post_reads\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, postReadPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from postRead slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for post_reads")
	}

	if len(postReadAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PostRead) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPostRead(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PostReadSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PostReadSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), postReadPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"post_reads\".* FROM \"post_reads\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, postReadPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "orm: unable to reload all in PostReadSlice")
	}

	*o = slice

	return nil
}

// PostReadExists checks if the PostRead row exists.
func PostReadExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"post_reads\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "orm: unable to check if post_reads exists")
	}

	return exists, nil
}

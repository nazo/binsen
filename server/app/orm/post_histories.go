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

// PostHistory is an object representing the database table.
type PostHistory struct {
	ID           int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	Version      int64     `boil:"version" json:"version" toml:"version" yaml:"version"`
	PostID       int64     `boil:"post_id" json:"post_id" toml:"post_id" yaml:"post_id"`
	Title        string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	Body         string    `boil:"body" json:"body" toml:"body" yaml:"body"`
	EditorUserID int64     `boil:"editor_user_id" json:"editor_user_id" toml:"editor_user_id" yaml:"editor_user_id"`
	EditStatus   int       `boil:"edit_status" json:"edit_status" toml:"edit_status" yaml:"edit_status"`
	CreatedAt    time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *postHistoryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L postHistoryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PostHistoryColumns = struct {
	ID           string
	Version      string
	PostID       string
	Title        string
	Body         string
	EditorUserID string
	EditStatus   string
	CreatedAt    string
	UpdatedAt    string
}{
	ID:           "id",
	Version:      "version",
	PostID:       "post_id",
	Title:        "title",
	Body:         "body",
	EditorUserID: "editor_user_id",
	EditStatus:   "edit_status",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// PostHistoryRels is where relationship names are stored.
var PostHistoryRels = struct {
	Post       string
	EditorUser string
}{
	Post:       "Post",
	EditorUser: "EditorUser",
}

// postHistoryR is where relationships are stored.
type postHistoryR struct {
	Post       *Post
	EditorUser *User
}

// NewStruct creates a new relationship struct
func (*postHistoryR) NewStruct() *postHistoryR {
	return &postHistoryR{}
}

// postHistoryL is where Load methods for each relationship are stored.
type postHistoryL struct{}

var (
	postHistoryColumns               = []string{"id", "version", "post_id", "title", "body", "editor_user_id", "edit_status", "created_at", "updated_at"}
	postHistoryColumnsWithoutDefault = []string{"version", "post_id", "title", "body", "editor_user_id", "edit_status", "created_at", "updated_at"}
	postHistoryColumnsWithDefault    = []string{"id"}
	postHistoryPrimaryKeyColumns     = []string{"id"}
)

type (
	// PostHistorySlice is an alias for a slice of pointers to PostHistory.
	// This should generally be used opposed to []PostHistory.
	PostHistorySlice []*PostHistory
	// PostHistoryHook is the signature for custom PostHistory hook methods
	PostHistoryHook func(context.Context, boil.ContextExecutor, *PostHistory) error

	postHistoryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	postHistoryType                 = reflect.TypeOf(&PostHistory{})
	postHistoryMapping              = queries.MakeStructMapping(postHistoryType)
	postHistoryPrimaryKeyMapping, _ = queries.BindMapping(postHistoryType, postHistoryMapping, postHistoryPrimaryKeyColumns)
	postHistoryInsertCacheMut       sync.RWMutex
	postHistoryInsertCache          = make(map[string]insertCache)
	postHistoryUpdateCacheMut       sync.RWMutex
	postHistoryUpdateCache          = make(map[string]updateCache)
	postHistoryUpsertCacheMut       sync.RWMutex
	postHistoryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var postHistoryBeforeInsertHooks []PostHistoryHook
var postHistoryBeforeUpdateHooks []PostHistoryHook
var postHistoryBeforeDeleteHooks []PostHistoryHook
var postHistoryBeforeUpsertHooks []PostHistoryHook

var postHistoryAfterInsertHooks []PostHistoryHook
var postHistoryAfterSelectHooks []PostHistoryHook
var postHistoryAfterUpdateHooks []PostHistoryHook
var postHistoryAfterDeleteHooks []PostHistoryHook
var postHistoryAfterUpsertHooks []PostHistoryHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PostHistory) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range postHistoryBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PostHistory) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range postHistoryBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PostHistory) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range postHistoryBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PostHistory) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range postHistoryBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PostHistory) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range postHistoryAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PostHistory) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range postHistoryAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PostHistory) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range postHistoryAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PostHistory) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range postHistoryAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PostHistory) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range postHistoryAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPostHistoryHook registers your hook function for all future operations.
func AddPostHistoryHook(hookPoint boil.HookPoint, postHistoryHook PostHistoryHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		postHistoryBeforeInsertHooks = append(postHistoryBeforeInsertHooks, postHistoryHook)
	case boil.BeforeUpdateHook:
		postHistoryBeforeUpdateHooks = append(postHistoryBeforeUpdateHooks, postHistoryHook)
	case boil.BeforeDeleteHook:
		postHistoryBeforeDeleteHooks = append(postHistoryBeforeDeleteHooks, postHistoryHook)
	case boil.BeforeUpsertHook:
		postHistoryBeforeUpsertHooks = append(postHistoryBeforeUpsertHooks, postHistoryHook)
	case boil.AfterInsertHook:
		postHistoryAfterInsertHooks = append(postHistoryAfterInsertHooks, postHistoryHook)
	case boil.AfterSelectHook:
		postHistoryAfterSelectHooks = append(postHistoryAfterSelectHooks, postHistoryHook)
	case boil.AfterUpdateHook:
		postHistoryAfterUpdateHooks = append(postHistoryAfterUpdateHooks, postHistoryHook)
	case boil.AfterDeleteHook:
		postHistoryAfterDeleteHooks = append(postHistoryAfterDeleteHooks, postHistoryHook)
	case boil.AfterUpsertHook:
		postHistoryAfterUpsertHooks = append(postHistoryAfterUpsertHooks, postHistoryHook)
	}
}

// One returns a single postHistory record from the query.
func (q postHistoryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PostHistory, error) {
	o := &PostHistory{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: failed to execute a one query for post_histories")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PostHistory records from the query.
func (q postHistoryQuery) All(ctx context.Context, exec boil.ContextExecutor) (PostHistorySlice, error) {
	var o []*PostHistory

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "orm: failed to assign all query results to PostHistory slice")
	}

	if len(postHistoryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PostHistory records in the query.
func (q postHistoryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to count post_histories rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q postHistoryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "orm: failed to check if post_histories exists")
	}

	return count > 0, nil
}

// Post pointed to by the foreign key.
func (o *PostHistory) Post(mods ...qm.QueryMod) postQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.PostID),
	}

	queryMods = append(queryMods, mods...)

	query := Posts(queryMods...)
	queries.SetFrom(query.Query, "\"posts\"")

	return query
}

// EditorUser pointed to by the foreign key.
func (o *PostHistory) EditorUser(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.EditorUserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// LoadPost allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (postHistoryL) LoadPost(ctx context.Context, e boil.ContextExecutor, singular bool, maybePostHistory interface{}, mods queries.Applicator) error {
	var slice []*PostHistory
	var object *PostHistory

	if singular {
		object = maybePostHistory.(*PostHistory)
	} else {
		slice = *maybePostHistory.(*[]*PostHistory)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &postHistoryR{}
		}
		args = append(args, object.PostID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &postHistoryR{}
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

	if len(postHistoryAfterSelectHooks) != 0 {
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
		foreign.R.PostHistories = append(foreign.R.PostHistories, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PostID == foreign.ID {
				local.R.Post = foreign
				if foreign.R == nil {
					foreign.R = &postR{}
				}
				foreign.R.PostHistories = append(foreign.R.PostHistories, local)
				break
			}
		}
	}

	return nil
}

// LoadEditorUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (postHistoryL) LoadEditorUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybePostHistory interface{}, mods queries.Applicator) error {
	var slice []*PostHistory
	var object *PostHistory

	if singular {
		object = maybePostHistory.(*PostHistory)
	} else {
		slice = *maybePostHistory.(*[]*PostHistory)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &postHistoryR{}
		}
		args = append(args, object.EditorUserID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &postHistoryR{}
			}

			for _, a := range args {
				if a == obj.EditorUserID {
					continue Outer
				}
			}

			args = append(args, obj.EditorUserID)
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

	if len(postHistoryAfterSelectHooks) != 0 {
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
		object.R.EditorUser = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.EditorUserPostHistories = append(foreign.R.EditorUserPostHistories, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.EditorUserID == foreign.ID {
				local.R.EditorUser = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.EditorUserPostHistories = append(foreign.R.EditorUserPostHistories, local)
				break
			}
		}
	}

	return nil
}

// SetPost of the postHistory to the related item.
// Sets o.R.Post to related.
// Adds o to related.R.PostHistories.
func (o *PostHistory) SetPost(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Post) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"post_histories\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"post_id"}),
		strmangle.WhereClause("\"", "\"", 2, postHistoryPrimaryKeyColumns),
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
		o.R = &postHistoryR{
			Post: related,
		}
	} else {
		o.R.Post = related
	}

	if related.R == nil {
		related.R = &postR{
			PostHistories: PostHistorySlice{o},
		}
	} else {
		related.R.PostHistories = append(related.R.PostHistories, o)
	}

	return nil
}

// SetEditorUser of the postHistory to the related item.
// Sets o.R.EditorUser to related.
// Adds o to related.R.EditorUserPostHistories.
func (o *PostHistory) SetEditorUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"post_histories\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"editor_user_id"}),
		strmangle.WhereClause("\"", "\"", 2, postHistoryPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.EditorUserID = related.ID
	if o.R == nil {
		o.R = &postHistoryR{
			EditorUser: related,
		}
	} else {
		o.R.EditorUser = related
	}

	if related.R == nil {
		related.R = &userR{
			EditorUserPostHistories: PostHistorySlice{o},
		}
	} else {
		related.R.EditorUserPostHistories = append(related.R.EditorUserPostHistories, o)
	}

	return nil
}

// PostHistories retrieves all the records using an executor.
func PostHistories(mods ...qm.QueryMod) postHistoryQuery {
	mods = append(mods, qm.From("\"post_histories\""))
	return postHistoryQuery{NewQuery(mods...)}
}

// FindPostHistory retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPostHistory(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*PostHistory, error) {
	postHistoryObj := &PostHistory{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"post_histories\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, postHistoryObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: unable to select from post_histories")
	}

	return postHistoryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PostHistory) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no post_histories provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(postHistoryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	postHistoryInsertCacheMut.RLock()
	cache, cached := postHistoryInsertCache[key]
	postHistoryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			postHistoryColumns,
			postHistoryColumnsWithDefault,
			postHistoryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(postHistoryType, postHistoryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(postHistoryType, postHistoryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"post_histories\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"post_histories\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "orm: unable to insert into post_histories")
	}

	if !cached {
		postHistoryInsertCacheMut.Lock()
		postHistoryInsertCache[key] = cache
		postHistoryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the PostHistory.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PostHistory) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	postHistoryUpdateCacheMut.RLock()
	cache, cached := postHistoryUpdateCache[key]
	postHistoryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			postHistoryColumns,
			postHistoryPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("orm: unable to update post_histories, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"post_histories\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, postHistoryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(postHistoryType, postHistoryMapping, append(wl, postHistoryPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "orm: unable to update post_histories row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by update for post_histories")
	}

	if !cached {
		postHistoryUpdateCacheMut.Lock()
		postHistoryUpdateCache[key] = cache
		postHistoryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q postHistoryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all for post_histories")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected for post_histories")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PostHistorySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), postHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"post_histories\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, postHistoryPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all in postHistory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected all in update all postHistory")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PostHistory) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no post_histories provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(postHistoryColumnsWithDefault, o)

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

	postHistoryUpsertCacheMut.RLock()
	cache, cached := postHistoryUpsertCache[key]
	postHistoryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			postHistoryColumns,
			postHistoryColumnsWithDefault,
			postHistoryColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			postHistoryColumns,
			postHistoryPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("orm: unable to upsert post_histories, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(postHistoryPrimaryKeyColumns))
			copy(conflict, postHistoryPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"post_histories\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(postHistoryType, postHistoryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(postHistoryType, postHistoryMapping, ret)
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
		return errors.Wrap(err, "orm: unable to upsert post_histories")
	}

	if !cached {
		postHistoryUpsertCacheMut.Lock()
		postHistoryUpsertCache[key] = cache
		postHistoryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single PostHistory record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PostHistory) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("orm: no PostHistory provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), postHistoryPrimaryKeyMapping)
	sql := "DELETE FROM \"post_histories\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete from post_histories")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by delete for post_histories")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q postHistoryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("orm: no postHistoryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from post_histories")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for post_histories")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PostHistorySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("orm: no PostHistory slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(postHistoryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), postHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"post_histories\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, postHistoryPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from postHistory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for post_histories")
	}

	if len(postHistoryAfterDeleteHooks) != 0 {
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
func (o *PostHistory) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPostHistory(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PostHistorySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PostHistorySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), postHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"post_histories\".* FROM \"post_histories\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, postHistoryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "orm: unable to reload all in PostHistorySlice")
	}

	*o = slice

	return nil
}

// PostHistoryExists checks if the PostHistory row exists.
func PostHistoryExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"post_histories\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "orm: unable to check if post_histories exists")
	}

	return exists, nil
}
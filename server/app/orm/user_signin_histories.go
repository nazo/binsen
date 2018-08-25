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
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// UserSigninHistory is an object representing the database table.
type UserSigninHistory struct {
	ID          int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID      int64       `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Provider    int         `boil:"provider" json:"provider" toml:"provider" yaml:"provider"`
	Ipv4Address null.String `boil:"ipv4_address" json:"ipv4_address,omitempty" toml:"ipv4_address" yaml:"ipv4_address,omitempty"`
	CreatedAt   time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *userSigninHistoryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userSigninHistoryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserSigninHistoryColumns = struct {
	ID          string
	UserID      string
	Provider    string
	Ipv4Address string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	UserID:      "user_id",
	Provider:    "provider",
	Ipv4Address: "ipv4_address",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// UserSigninHistoryRels is where relationship names are stored.
var UserSigninHistoryRels = struct {
	User string
}{
	User: "User",
}

// userSigninHistoryR is where relationships are stored.
type userSigninHistoryR struct {
	User *User
}

// NewStruct creates a new relationship struct
func (*userSigninHistoryR) NewStruct() *userSigninHistoryR {
	return &userSigninHistoryR{}
}

// userSigninHistoryL is where Load methods for each relationship are stored.
type userSigninHistoryL struct{}

var (
	userSigninHistoryColumns               = []string{"id", "user_id", "provider", "ipv4_address", "created_at", "updated_at"}
	userSigninHistoryColumnsWithoutDefault = []string{"user_id", "provider", "ipv4_address", "created_at", "updated_at"}
	userSigninHistoryColumnsWithDefault    = []string{"id"}
	userSigninHistoryPrimaryKeyColumns     = []string{"id"}
)

type (
	// UserSigninHistorySlice is an alias for a slice of pointers to UserSigninHistory.
	// This should generally be used opposed to []UserSigninHistory.
	UserSigninHistorySlice []*UserSigninHistory
	// UserSigninHistoryHook is the signature for custom UserSigninHistory hook methods
	UserSigninHistoryHook func(context.Context, boil.ContextExecutor, *UserSigninHistory) error

	userSigninHistoryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userSigninHistoryType                 = reflect.TypeOf(&UserSigninHistory{})
	userSigninHistoryMapping              = queries.MakeStructMapping(userSigninHistoryType)
	userSigninHistoryPrimaryKeyMapping, _ = queries.BindMapping(userSigninHistoryType, userSigninHistoryMapping, userSigninHistoryPrimaryKeyColumns)
	userSigninHistoryInsertCacheMut       sync.RWMutex
	userSigninHistoryInsertCache          = make(map[string]insertCache)
	userSigninHistoryUpdateCacheMut       sync.RWMutex
	userSigninHistoryUpdateCache          = make(map[string]updateCache)
	userSigninHistoryUpsertCacheMut       sync.RWMutex
	userSigninHistoryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var userSigninHistoryBeforeInsertHooks []UserSigninHistoryHook
var userSigninHistoryBeforeUpdateHooks []UserSigninHistoryHook
var userSigninHistoryBeforeDeleteHooks []UserSigninHistoryHook
var userSigninHistoryBeforeUpsertHooks []UserSigninHistoryHook

var userSigninHistoryAfterInsertHooks []UserSigninHistoryHook
var userSigninHistoryAfterSelectHooks []UserSigninHistoryHook
var userSigninHistoryAfterUpdateHooks []UserSigninHistoryHook
var userSigninHistoryAfterDeleteHooks []UserSigninHistoryHook
var userSigninHistoryAfterUpsertHooks []UserSigninHistoryHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserSigninHistory) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range userSigninHistoryBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserSigninHistory) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range userSigninHistoryBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserSigninHistory) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range userSigninHistoryBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserSigninHistory) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range userSigninHistoryBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserSigninHistory) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range userSigninHistoryAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserSigninHistory) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range userSigninHistoryAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserSigninHistory) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range userSigninHistoryAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserSigninHistory) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range userSigninHistoryAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserSigninHistory) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range userSigninHistoryAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserSigninHistoryHook registers your hook function for all future operations.
func AddUserSigninHistoryHook(hookPoint boil.HookPoint, userSigninHistoryHook UserSigninHistoryHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		userSigninHistoryBeforeInsertHooks = append(userSigninHistoryBeforeInsertHooks, userSigninHistoryHook)
	case boil.BeforeUpdateHook:
		userSigninHistoryBeforeUpdateHooks = append(userSigninHistoryBeforeUpdateHooks, userSigninHistoryHook)
	case boil.BeforeDeleteHook:
		userSigninHistoryBeforeDeleteHooks = append(userSigninHistoryBeforeDeleteHooks, userSigninHistoryHook)
	case boil.BeforeUpsertHook:
		userSigninHistoryBeforeUpsertHooks = append(userSigninHistoryBeforeUpsertHooks, userSigninHistoryHook)
	case boil.AfterInsertHook:
		userSigninHistoryAfterInsertHooks = append(userSigninHistoryAfterInsertHooks, userSigninHistoryHook)
	case boil.AfterSelectHook:
		userSigninHistoryAfterSelectHooks = append(userSigninHistoryAfterSelectHooks, userSigninHistoryHook)
	case boil.AfterUpdateHook:
		userSigninHistoryAfterUpdateHooks = append(userSigninHistoryAfterUpdateHooks, userSigninHistoryHook)
	case boil.AfterDeleteHook:
		userSigninHistoryAfterDeleteHooks = append(userSigninHistoryAfterDeleteHooks, userSigninHistoryHook)
	case boil.AfterUpsertHook:
		userSigninHistoryAfterUpsertHooks = append(userSigninHistoryAfterUpsertHooks, userSigninHistoryHook)
	}
}

// One returns a single userSigninHistory record from the query.
func (q userSigninHistoryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserSigninHistory, error) {
	o := &UserSigninHistory{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: failed to execute a one query for user_signin_histories")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UserSigninHistory records from the query.
func (q userSigninHistoryQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserSigninHistorySlice, error) {
	var o []*UserSigninHistory

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "orm: failed to assign all query results to UserSigninHistory slice")
	}

	if len(userSigninHistoryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UserSigninHistory records in the query.
func (q userSigninHistoryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to count user_signin_histories rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userSigninHistoryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "orm: failed to check if user_signin_histories exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *UserSigninHistory) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userSigninHistoryL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserSigninHistory interface{}, mods queries.Applicator) error {
	var slice []*UserSigninHistory
	var object *UserSigninHistory

	if singular {
		object = maybeUserSigninHistory.(*UserSigninHistory)
	} else {
		slice = *maybeUserSigninHistory.(*[]*UserSigninHistory)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userSigninHistoryR{}
		}
		args = append(args, object.UserID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userSigninHistoryR{}
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

	if len(userSigninHistoryAfterSelectHooks) != 0 {
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
		foreign.R.UserSigninHistories = append(foreign.R.UserSigninHistories, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UserSigninHistories = append(foreign.R.UserSigninHistories, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the userSigninHistory to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserSigninHistories.
func (o *UserSigninHistory) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_signin_histories\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, userSigninHistoryPrimaryKeyColumns),
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
		o.R = &userSigninHistoryR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			UserSigninHistories: UserSigninHistorySlice{o},
		}
	} else {
		related.R.UserSigninHistories = append(related.R.UserSigninHistories, o)
	}

	return nil
}

// UserSigninHistories retrieves all the records using an executor.
func UserSigninHistories(mods ...qm.QueryMod) userSigninHistoryQuery {
	mods = append(mods, qm.From("\"user_signin_histories\""))
	return userSigninHistoryQuery{NewQuery(mods...)}
}

// FindUserSigninHistory retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserSigninHistory(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*UserSigninHistory, error) {
	userSigninHistoryObj := &UserSigninHistory{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"user_signin_histories\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, userSigninHistoryObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: unable to select from user_signin_histories")
	}

	return userSigninHistoryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserSigninHistory) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no user_signin_histories provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(userSigninHistoryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userSigninHistoryInsertCacheMut.RLock()
	cache, cached := userSigninHistoryInsertCache[key]
	userSigninHistoryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userSigninHistoryColumns,
			userSigninHistoryColumnsWithDefault,
			userSigninHistoryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userSigninHistoryType, userSigninHistoryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userSigninHistoryType, userSigninHistoryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"user_signin_histories\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"user_signin_histories\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "orm: unable to insert into user_signin_histories")
	}

	if !cached {
		userSigninHistoryInsertCacheMut.Lock()
		userSigninHistoryInsertCache[key] = cache
		userSigninHistoryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UserSigninHistory.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserSigninHistory) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userSigninHistoryUpdateCacheMut.RLock()
	cache, cached := userSigninHistoryUpdateCache[key]
	userSigninHistoryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userSigninHistoryColumns,
			userSigninHistoryPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("orm: unable to update user_signin_histories, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"user_signin_histories\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userSigninHistoryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userSigninHistoryType, userSigninHistoryMapping, append(wl, userSigninHistoryPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "orm: unable to update user_signin_histories row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by update for user_signin_histories")
	}

	if !cached {
		userSigninHistoryUpdateCacheMut.Lock()
		userSigninHistoryUpdateCache[key] = cache
		userSigninHistoryUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userSigninHistoryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all for user_signin_histories")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected for user_signin_histories")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserSigninHistorySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userSigninHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"user_signin_histories\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, userSigninHistoryPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all in userSigninHistory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected all in update all userSigninHistory")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserSigninHistory) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no user_signin_histories provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userSigninHistoryColumnsWithDefault, o)

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

	userSigninHistoryUpsertCacheMut.RLock()
	cache, cached := userSigninHistoryUpsertCache[key]
	userSigninHistoryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userSigninHistoryColumns,
			userSigninHistoryColumnsWithDefault,
			userSigninHistoryColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			userSigninHistoryColumns,
			userSigninHistoryPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("orm: unable to upsert user_signin_histories, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(userSigninHistoryPrimaryKeyColumns))
			copy(conflict, userSigninHistoryPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"user_signin_histories\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(userSigninHistoryType, userSigninHistoryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userSigninHistoryType, userSigninHistoryMapping, ret)
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
		return errors.Wrap(err, "orm: unable to upsert user_signin_histories")
	}

	if !cached {
		userSigninHistoryUpsertCacheMut.Lock()
		userSigninHistoryUpsertCache[key] = cache
		userSigninHistoryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UserSigninHistory record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserSigninHistory) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("orm: no UserSigninHistory provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userSigninHistoryPrimaryKeyMapping)
	sql := "DELETE FROM \"user_signin_histories\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete from user_signin_histories")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by delete for user_signin_histories")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userSigninHistoryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("orm: no userSigninHistoryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from user_signin_histories")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for user_signin_histories")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserSigninHistorySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("orm: no UserSigninHistory slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(userSigninHistoryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userSigninHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"user_signin_histories\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userSigninHistoryPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from userSigninHistory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for user_signin_histories")
	}

	if len(userSigninHistoryAfterDeleteHooks) != 0 {
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
func (o *UserSigninHistory) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserSigninHistory(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserSigninHistorySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserSigninHistorySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userSigninHistoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"user_signin_histories\".* FROM \"user_signin_histories\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userSigninHistoryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "orm: unable to reload all in UserSigninHistorySlice")
	}

	*o = slice

	return nil
}

// UserSigninHistoryExists checks if the UserSigninHistory row exists.
func UserSigninHistoryExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"user_signin_histories\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "orm: unable to check if user_signin_histories exists")
	}

	return exists, nil
}
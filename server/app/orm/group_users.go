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

// GroupUser is an object representing the database table.
type GroupUser struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID    int64     `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	GroupID   int64     `boil:"group_id" json:"group_id" toml:"group_id" yaml:"group_id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *groupUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L groupUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GroupUserColumns = struct {
	ID        string
	UserID    string
	GroupID   string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	GroupID:   "group_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// GroupUserRels is where relationship names are stored.
var GroupUserRels = struct {
	User  string
	Group string
}{
	User:  "User",
	Group: "Group",
}

// groupUserR is where relationships are stored.
type groupUserR struct {
	User  *User
	Group *Group
}

// NewStruct creates a new relationship struct
func (*groupUserR) NewStruct() *groupUserR {
	return &groupUserR{}
}

// groupUserL is where Load methods for each relationship are stored.
type groupUserL struct{}

var (
	groupUserColumns               = []string{"id", "user_id", "group_id", "created_at", "updated_at"}
	groupUserColumnsWithoutDefault = []string{"user_id", "group_id", "created_at", "updated_at"}
	groupUserColumnsWithDefault    = []string{"id"}
	groupUserPrimaryKeyColumns     = []string{"id"}
)

type (
	// GroupUserSlice is an alias for a slice of pointers to GroupUser.
	// This should generally be used opposed to []GroupUser.
	GroupUserSlice []*GroupUser
	// GroupUserHook is the signature for custom GroupUser hook methods
	GroupUserHook func(context.Context, boil.ContextExecutor, *GroupUser) error

	groupUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	groupUserType                 = reflect.TypeOf(&GroupUser{})
	groupUserMapping              = queries.MakeStructMapping(groupUserType)
	groupUserPrimaryKeyMapping, _ = queries.BindMapping(groupUserType, groupUserMapping, groupUserPrimaryKeyColumns)
	groupUserInsertCacheMut       sync.RWMutex
	groupUserInsertCache          = make(map[string]insertCache)
	groupUserUpdateCacheMut       sync.RWMutex
	groupUserUpdateCache          = make(map[string]updateCache)
	groupUserUpsertCacheMut       sync.RWMutex
	groupUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var groupUserBeforeInsertHooks []GroupUserHook
var groupUserBeforeUpdateHooks []GroupUserHook
var groupUserBeforeDeleteHooks []GroupUserHook
var groupUserBeforeUpsertHooks []GroupUserHook

var groupUserAfterInsertHooks []GroupUserHook
var groupUserAfterSelectHooks []GroupUserHook
var groupUserAfterUpdateHooks []GroupUserHook
var groupUserAfterDeleteHooks []GroupUserHook
var groupUserAfterUpsertHooks []GroupUserHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *GroupUser) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range groupUserBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *GroupUser) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range groupUserBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *GroupUser) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range groupUserBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *GroupUser) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range groupUserBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *GroupUser) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range groupUserAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *GroupUser) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range groupUserAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *GroupUser) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range groupUserAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *GroupUser) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range groupUserAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *GroupUser) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range groupUserAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGroupUserHook registers your hook function for all future operations.
func AddGroupUserHook(hookPoint boil.HookPoint, groupUserHook GroupUserHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		groupUserBeforeInsertHooks = append(groupUserBeforeInsertHooks, groupUserHook)
	case boil.BeforeUpdateHook:
		groupUserBeforeUpdateHooks = append(groupUserBeforeUpdateHooks, groupUserHook)
	case boil.BeforeDeleteHook:
		groupUserBeforeDeleteHooks = append(groupUserBeforeDeleteHooks, groupUserHook)
	case boil.BeforeUpsertHook:
		groupUserBeforeUpsertHooks = append(groupUserBeforeUpsertHooks, groupUserHook)
	case boil.AfterInsertHook:
		groupUserAfterInsertHooks = append(groupUserAfterInsertHooks, groupUserHook)
	case boil.AfterSelectHook:
		groupUserAfterSelectHooks = append(groupUserAfterSelectHooks, groupUserHook)
	case boil.AfterUpdateHook:
		groupUserAfterUpdateHooks = append(groupUserAfterUpdateHooks, groupUserHook)
	case boil.AfterDeleteHook:
		groupUserAfterDeleteHooks = append(groupUserAfterDeleteHooks, groupUserHook)
	case boil.AfterUpsertHook:
		groupUserAfterUpsertHooks = append(groupUserAfterUpsertHooks, groupUserHook)
	}
}

// One returns a single groupUser record from the query.
func (q groupUserQuery) One(ctx context.Context, exec boil.ContextExecutor) (*GroupUser, error) {
	o := &GroupUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: failed to execute a one query for group_users")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all GroupUser records from the query.
func (q groupUserQuery) All(ctx context.Context, exec boil.ContextExecutor) (GroupUserSlice, error) {
	var o []*GroupUser

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "orm: failed to assign all query results to GroupUser slice")
	}

	if len(groupUserAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all GroupUser records in the query.
func (q groupUserQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to count group_users rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q groupUserQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "orm: failed to check if group_users exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *GroupUser) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// Group pointed to by the foreign key.
func (o *GroupUser) Group(mods ...qm.QueryMod) groupQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.GroupID),
	}

	queryMods = append(queryMods, mods...)

	query := Groups(queryMods...)
	queries.SetFrom(query.Query, "\"groups\"")

	return query
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (groupUserL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGroupUser interface{}, mods queries.Applicator) error {
	var slice []*GroupUser
	var object *GroupUser

	if singular {
		object = maybeGroupUser.(*GroupUser)
	} else {
		slice = *maybeGroupUser.(*[]*GroupUser)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &groupUserR{}
		}
		args = append(args, object.UserID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &groupUserR{}
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

	if len(groupUserAfterSelectHooks) != 0 {
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
		foreign.R.GroupUsers = append(foreign.R.GroupUsers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.GroupUsers = append(foreign.R.GroupUsers, local)
				break
			}
		}
	}

	return nil
}

// LoadGroup allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (groupUserL) LoadGroup(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGroupUser interface{}, mods queries.Applicator) error {
	var slice []*GroupUser
	var object *GroupUser

	if singular {
		object = maybeGroupUser.(*GroupUser)
	} else {
		slice = *maybeGroupUser.(*[]*GroupUser)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &groupUserR{}
		}
		args = append(args, object.GroupID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &groupUserR{}
			}

			for _, a := range args {
				if a == obj.GroupID {
					continue Outer
				}
			}

			args = append(args, obj.GroupID)
		}
	}

	query := NewQuery(qm.From(`groups`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Group")
	}

	var resultSlice []*Group
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Group")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for groups")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for groups")
	}

	if len(groupUserAfterSelectHooks) != 0 {
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
		object.R.Group = foreign
		if foreign.R == nil {
			foreign.R = &groupR{}
		}
		foreign.R.GroupUsers = append(foreign.R.GroupUsers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.GroupID == foreign.ID {
				local.R.Group = foreign
				if foreign.R == nil {
					foreign.R = &groupR{}
				}
				foreign.R.GroupUsers = append(foreign.R.GroupUsers, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the groupUser to the related item.
// Sets o.R.User to related.
// Adds o to related.R.GroupUsers.
func (o *GroupUser) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"group_users\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, groupUserPrimaryKeyColumns),
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
		o.R = &groupUserR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			GroupUsers: GroupUserSlice{o},
		}
	} else {
		related.R.GroupUsers = append(related.R.GroupUsers, o)
	}

	return nil
}

// SetGroup of the groupUser to the related item.
// Sets o.R.Group to related.
// Adds o to related.R.GroupUsers.
func (o *GroupUser) SetGroup(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Group) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"group_users\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"group_id"}),
		strmangle.WhereClause("\"", "\"", 2, groupUserPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.GroupID = related.ID
	if o.R == nil {
		o.R = &groupUserR{
			Group: related,
		}
	} else {
		o.R.Group = related
	}

	if related.R == nil {
		related.R = &groupR{
			GroupUsers: GroupUserSlice{o},
		}
	} else {
		related.R.GroupUsers = append(related.R.GroupUsers, o)
	}

	return nil
}

// GroupUsers retrieves all the records using an executor.
func GroupUsers(mods ...qm.QueryMod) groupUserQuery {
	mods = append(mods, qm.From("\"group_users\""))
	return groupUserQuery{NewQuery(mods...)}
}

// FindGroupUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGroupUser(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*GroupUser, error) {
	groupUserObj := &GroupUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"group_users\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, groupUserObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: unable to select from group_users")
	}

	return groupUserObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *GroupUser) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no group_users provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(groupUserColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	groupUserInsertCacheMut.RLock()
	cache, cached := groupUserInsertCache[key]
	groupUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			groupUserColumns,
			groupUserColumnsWithDefault,
			groupUserColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(groupUserType, groupUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(groupUserType, groupUserMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"group_users\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"group_users\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "orm: unable to insert into group_users")
	}

	if !cached {
		groupUserInsertCacheMut.Lock()
		groupUserInsertCache[key] = cache
		groupUserInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the GroupUser.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *GroupUser) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	groupUserUpdateCacheMut.RLock()
	cache, cached := groupUserUpdateCache[key]
	groupUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			groupUserColumns,
			groupUserPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("orm: unable to update group_users, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"group_users\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, groupUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(groupUserType, groupUserMapping, append(wl, groupUserPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "orm: unable to update group_users row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by update for group_users")
	}

	if !cached {
		groupUserUpdateCacheMut.Lock()
		groupUserUpdateCache[key] = cache
		groupUserUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q groupUserQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all for group_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected for group_users")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GroupUserSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), groupUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"group_users\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, groupUserPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all in groupUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected all in update all groupUser")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *GroupUser) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no group_users provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(groupUserColumnsWithDefault, o)

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

	groupUserUpsertCacheMut.RLock()
	cache, cached := groupUserUpsertCache[key]
	groupUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			groupUserColumns,
			groupUserColumnsWithDefault,
			groupUserColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			groupUserColumns,
			groupUserPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("orm: unable to upsert group_users, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(groupUserPrimaryKeyColumns))
			copy(conflict, groupUserPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"group_users\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(groupUserType, groupUserMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(groupUserType, groupUserMapping, ret)
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
		return errors.Wrap(err, "orm: unable to upsert group_users")
	}

	if !cached {
		groupUserUpsertCacheMut.Lock()
		groupUserUpsertCache[key] = cache
		groupUserUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single GroupUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *GroupUser) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("orm: no GroupUser provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), groupUserPrimaryKeyMapping)
	sql := "DELETE FROM \"group_users\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete from group_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by delete for group_users")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q groupUserQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("orm: no groupUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from group_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for group_users")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GroupUserSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("orm: no GroupUser slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(groupUserBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), groupUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"group_users\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, groupUserPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from groupUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for group_users")
	}

	if len(groupUserAfterDeleteHooks) != 0 {
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
func (o *GroupUser) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindGroupUser(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GroupUserSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GroupUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), groupUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"group_users\".* FROM \"group_users\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, groupUserPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "orm: unable to reload all in GroupUserSlice")
	}

	*o = slice

	return nil
}

// GroupUserExists checks if the GroupUser row exists.
func GroupUserExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"group_users\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "orm: unable to check if group_users exists")
	}

	return exists, nil
}

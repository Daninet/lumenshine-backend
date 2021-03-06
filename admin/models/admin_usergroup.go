// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
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

// AdminUsergroup is an object representing the database table.
type AdminUsergroup struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID    int       `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	GroupID   int       `boil:"group_id" json:"group_id" toml:"group_id" yaml:"group_id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	UpdatedBy string    `boil:"updated_by" json:"updated_by" toml:"updated_by" yaml:"updated_by"`

	R *adminUsergroupR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L adminUsergroupL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AdminUsergroupColumns = struct {
	ID        string
	UserID    string
	GroupID   string
	CreatedAt string
	UpdatedAt string
	UpdatedBy string
}{
	ID:        "id",
	UserID:    "user_id",
	GroupID:   "group_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	UpdatedBy: "updated_by",
}

// AdminUsergroupRels is where relationship names are stored.
var AdminUsergroupRels = struct {
	User  string
	Group string
}{
	User:  "User",
	Group: "Group",
}

// adminUsergroupR is where relationships are stored.
type adminUsergroupR struct {
	User  *AdminUser
	Group *AdminGroup
}

// NewStruct creates a new relationship struct
func (*adminUsergroupR) NewStruct() *adminUsergroupR {
	return &adminUsergroupR{}
}

// adminUsergroupL is where Load methods for each relationship are stored.
type adminUsergroupL struct{}

var (
	adminUsergroupColumns               = []string{"id", "user_id", "group_id", "created_at", "updated_at", "updated_by"}
	adminUsergroupColumnsWithoutDefault = []string{"user_id", "group_id", "updated_by"}
	adminUsergroupColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	adminUsergroupPrimaryKeyColumns     = []string{"id"}
)

type (
	// AdminUsergroupSlice is an alias for a slice of pointers to AdminUsergroup.
	// This should generally be used opposed to []AdminUsergroup.
	AdminUsergroupSlice []*AdminUsergroup
	// AdminUsergroupHook is the signature for custom AdminUsergroup hook methods
	AdminUsergroupHook func(boil.Executor, *AdminUsergroup) error

	adminUsergroupQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	adminUsergroupType                 = reflect.TypeOf(&AdminUsergroup{})
	adminUsergroupMapping              = queries.MakeStructMapping(adminUsergroupType)
	adminUsergroupPrimaryKeyMapping, _ = queries.BindMapping(adminUsergroupType, adminUsergroupMapping, adminUsergroupPrimaryKeyColumns)
	adminUsergroupInsertCacheMut       sync.RWMutex
	adminUsergroupInsertCache          = make(map[string]insertCache)
	adminUsergroupUpdateCacheMut       sync.RWMutex
	adminUsergroupUpdateCache          = make(map[string]updateCache)
	adminUsergroupUpsertCacheMut       sync.RWMutex
	adminUsergroupUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var adminUsergroupBeforeInsertHooks []AdminUsergroupHook
var adminUsergroupBeforeUpdateHooks []AdminUsergroupHook
var adminUsergroupBeforeDeleteHooks []AdminUsergroupHook
var adminUsergroupBeforeUpsertHooks []AdminUsergroupHook

var adminUsergroupAfterInsertHooks []AdminUsergroupHook
var adminUsergroupAfterSelectHooks []AdminUsergroupHook
var adminUsergroupAfterUpdateHooks []AdminUsergroupHook
var adminUsergroupAfterDeleteHooks []AdminUsergroupHook
var adminUsergroupAfterUpsertHooks []AdminUsergroupHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AdminUsergroup) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range adminUsergroupBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AdminUsergroup) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range adminUsergroupBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AdminUsergroup) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range adminUsergroupBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AdminUsergroup) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range adminUsergroupBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AdminUsergroup) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range adminUsergroupAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AdminUsergroup) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range adminUsergroupAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AdminUsergroup) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range adminUsergroupAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AdminUsergroup) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range adminUsergroupAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AdminUsergroup) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range adminUsergroupAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAdminUsergroupHook registers your hook function for all future operations.
func AddAdminUsergroupHook(hookPoint boil.HookPoint, adminUsergroupHook AdminUsergroupHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		adminUsergroupBeforeInsertHooks = append(adminUsergroupBeforeInsertHooks, adminUsergroupHook)
	case boil.BeforeUpdateHook:
		adminUsergroupBeforeUpdateHooks = append(adminUsergroupBeforeUpdateHooks, adminUsergroupHook)
	case boil.BeforeDeleteHook:
		adminUsergroupBeforeDeleteHooks = append(adminUsergroupBeforeDeleteHooks, adminUsergroupHook)
	case boil.BeforeUpsertHook:
		adminUsergroupBeforeUpsertHooks = append(adminUsergroupBeforeUpsertHooks, adminUsergroupHook)
	case boil.AfterInsertHook:
		adminUsergroupAfterInsertHooks = append(adminUsergroupAfterInsertHooks, adminUsergroupHook)
	case boil.AfterSelectHook:
		adminUsergroupAfterSelectHooks = append(adminUsergroupAfterSelectHooks, adminUsergroupHook)
	case boil.AfterUpdateHook:
		adminUsergroupAfterUpdateHooks = append(adminUsergroupAfterUpdateHooks, adminUsergroupHook)
	case boil.AfterDeleteHook:
		adminUsergroupAfterDeleteHooks = append(adminUsergroupAfterDeleteHooks, adminUsergroupHook)
	case boil.AfterUpsertHook:
		adminUsergroupAfterUpsertHooks = append(adminUsergroupAfterUpsertHooks, adminUsergroupHook)
	}
}

// OneG returns a single adminUsergroup record from the query using the global executor.
func (q adminUsergroupQuery) OneG() (*AdminUsergroup, error) {
	return q.One(boil.GetDB())
}

// One returns a single adminUsergroup record from the query.
func (q adminUsergroupQuery) One(exec boil.Executor) (*AdminUsergroup, error) {
	o := &AdminUsergroup{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for admin_usergroup")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all AdminUsergroup records from the query using the global executor.
func (q adminUsergroupQuery) AllG() (AdminUsergroupSlice, error) {
	return q.All(boil.GetDB())
}

// All returns all AdminUsergroup records from the query.
func (q adminUsergroupQuery) All(exec boil.Executor) (AdminUsergroupSlice, error) {
	var o []*AdminUsergroup

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AdminUsergroup slice")
	}

	if len(adminUsergroupAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all AdminUsergroup records in the query, and panics on error.
func (q adminUsergroupQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all AdminUsergroup records in the query.
func (q adminUsergroupQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count admin_usergroup rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q adminUsergroupQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q adminUsergroupQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if admin_usergroup exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *AdminUsergroup) User(mods ...qm.QueryMod) adminUserQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := AdminUsers(queryMods...)
	queries.SetFrom(query.Query, "\"admin_user\"")

	return query
}

// Group pointed to by the foreign key.
func (o *AdminUsergroup) Group(mods ...qm.QueryMod) adminGroupQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.GroupID),
	}

	queryMods = append(queryMods, mods...)

	query := AdminGroups(queryMods...)
	queries.SetFrom(query.Query, "\"admin_group\"")

	return query
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (adminUsergroupL) LoadUser(e boil.Executor, singular bool, maybeAdminUsergroup interface{}, mods queries.Applicator) error {
	var slice []*AdminUsergroup
	var object *AdminUsergroup

	if singular {
		object = maybeAdminUsergroup.(*AdminUsergroup)
	} else {
		slice = *maybeAdminUsergroup.(*[]*AdminUsergroup)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &adminUsergroupR{}
		}
		args = append(args, object.UserID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &adminUsergroupR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)
		}
	}

	query := NewQuery(qm.From(`admin_user`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AdminUser")
	}

	var resultSlice []*AdminUser
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AdminUser")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for admin_user")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for admin_user")
	}

	if len(adminUsergroupAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
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
			foreign.R = &adminUserR{}
		}
		foreign.R.UserAdminUsergroups = append(foreign.R.UserAdminUsergroups, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &adminUserR{}
				}
				foreign.R.UserAdminUsergroups = append(foreign.R.UserAdminUsergroups, local)
				break
			}
		}
	}

	return nil
}

// LoadGroup allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (adminUsergroupL) LoadGroup(e boil.Executor, singular bool, maybeAdminUsergroup interface{}, mods queries.Applicator) error {
	var slice []*AdminUsergroup
	var object *AdminUsergroup

	if singular {
		object = maybeAdminUsergroup.(*AdminUsergroup)
	} else {
		slice = *maybeAdminUsergroup.(*[]*AdminUsergroup)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &adminUsergroupR{}
		}
		args = append(args, object.GroupID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &adminUsergroupR{}
			}

			for _, a := range args {
				if a == obj.GroupID {
					continue Outer
				}
			}

			args = append(args, obj.GroupID)
		}
	}

	query := NewQuery(qm.From(`admin_group`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AdminGroup")
	}

	var resultSlice []*AdminGroup
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AdminGroup")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for admin_group")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for admin_group")
	}

	if len(adminUsergroupAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
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
			foreign.R = &adminGroupR{}
		}
		foreign.R.GroupAdminUsergroups = append(foreign.R.GroupAdminUsergroups, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.GroupID == foreign.ID {
				local.R.Group = foreign
				if foreign.R == nil {
					foreign.R = &adminGroupR{}
				}
				foreign.R.GroupAdminUsergroups = append(foreign.R.GroupAdminUsergroups, local)
				break
			}
		}
	}

	return nil
}

// SetUserG of the adminUsergroup to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserAdminUsergroups.
// Uses the global database handle.
func (o *AdminUsergroup) SetUserG(insert bool, related *AdminUser) error {
	return o.SetUser(boil.GetDB(), insert, related)
}

// SetUser of the adminUsergroup to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserAdminUsergroups.
func (o *AdminUsergroup) SetUser(exec boil.Executor, insert bool, related *AdminUser) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"admin_usergroup\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, adminUsergroupPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &adminUsergroupR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &adminUserR{
			UserAdminUsergroups: AdminUsergroupSlice{o},
		}
	} else {
		related.R.UserAdminUsergroups = append(related.R.UserAdminUsergroups, o)
	}

	return nil
}

// SetGroupG of the adminUsergroup to the related item.
// Sets o.R.Group to related.
// Adds o to related.R.GroupAdminUsergroups.
// Uses the global database handle.
func (o *AdminUsergroup) SetGroupG(insert bool, related *AdminGroup) error {
	return o.SetGroup(boil.GetDB(), insert, related)
}

// SetGroup of the adminUsergroup to the related item.
// Sets o.R.Group to related.
// Adds o to related.R.GroupAdminUsergroups.
func (o *AdminUsergroup) SetGroup(exec boil.Executor, insert bool, related *AdminGroup) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"admin_usergroup\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"group_id"}),
		strmangle.WhereClause("\"", "\"", 2, adminUsergroupPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.GroupID = related.ID
	if o.R == nil {
		o.R = &adminUsergroupR{
			Group: related,
		}
	} else {
		o.R.Group = related
	}

	if related.R == nil {
		related.R = &adminGroupR{
			GroupAdminUsergroups: AdminUsergroupSlice{o},
		}
	} else {
		related.R.GroupAdminUsergroups = append(related.R.GroupAdminUsergroups, o)
	}

	return nil
}

// AdminUsergroups retrieves all the records using an executor.
func AdminUsergroups(mods ...qm.QueryMod) adminUsergroupQuery {
	mods = append(mods, qm.From("\"admin_usergroup\""))
	return adminUsergroupQuery{NewQuery(mods...)}
}

// FindAdminUsergroupG retrieves a single record by ID.
func FindAdminUsergroupG(iD int, selectCols ...string) (*AdminUsergroup, error) {
	return FindAdminUsergroup(boil.GetDB(), iD, selectCols...)
}

// FindAdminUsergroup retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAdminUsergroup(exec boil.Executor, iD int, selectCols ...string) (*AdminUsergroup, error) {
	adminUsergroupObj := &AdminUsergroup{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"admin_usergroup\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, adminUsergroupObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from admin_usergroup")
	}

	return adminUsergroupObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *AdminUsergroup) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AdminUsergroup) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no admin_usergroup provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(adminUsergroupColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	adminUsergroupInsertCacheMut.RLock()
	cache, cached := adminUsergroupInsertCache[key]
	adminUsergroupInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			adminUsergroupColumns,
			adminUsergroupColumnsWithDefault,
			adminUsergroupColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(adminUsergroupType, adminUsergroupMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(adminUsergroupType, adminUsergroupMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"admin_usergroup\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"admin_usergroup\" %sDEFAULT VALUES%s"
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
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into admin_usergroup")
	}

	if !cached {
		adminUsergroupInsertCacheMut.Lock()
		adminUsergroupInsertCache[key] = cache
		adminUsergroupInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single AdminUsergroup record using the global executor.
// See Update for more documentation.
func (o *AdminUsergroup) UpdateG(columns boil.Columns) (int64, error) {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the AdminUsergroup.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AdminUsergroup) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	adminUsergroupUpdateCacheMut.RLock()
	cache, cached := adminUsergroupUpdateCache[key]
	adminUsergroupUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			adminUsergroupColumns,
			adminUsergroupPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update admin_usergroup, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"admin_usergroup\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, adminUsergroupPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(adminUsergroupType, adminUsergroupMapping, append(wl, adminUsergroupPrimaryKeyColumns...))
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
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update admin_usergroup row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for admin_usergroup")
	}

	if !cached {
		adminUsergroupUpdateCacheMut.Lock()
		adminUsergroupUpdateCache[key] = cache
		adminUsergroupUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q adminUsergroupQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for admin_usergroup")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for admin_usergroup")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AdminUsergroupSlice) UpdateAllG(cols M) (int64, error) {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AdminUsergroupSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), adminUsergroupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"admin_usergroup\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, adminUsergroupPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in adminUsergroup slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all adminUsergroup")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *AdminUsergroup) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AdminUsergroup) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no admin_usergroup provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(adminUsergroupColumnsWithDefault, o)

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

	adminUsergroupUpsertCacheMut.RLock()
	cache, cached := adminUsergroupUpsertCache[key]
	adminUsergroupUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			adminUsergroupColumns,
			adminUsergroupColumnsWithDefault,
			adminUsergroupColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			adminUsergroupColumns,
			adminUsergroupPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert admin_usergroup, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(adminUsergroupPrimaryKeyColumns))
			copy(conflict, adminUsergroupPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"admin_usergroup\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(adminUsergroupType, adminUsergroupMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(adminUsergroupType, adminUsergroupMapping, ret)
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
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert admin_usergroup")
	}

	if !cached {
		adminUsergroupUpsertCacheMut.Lock()
		adminUsergroupUpsertCache[key] = cache
		adminUsergroupUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteG deletes a single AdminUsergroup record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *AdminUsergroup) DeleteG() (int64, error) {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single AdminUsergroup record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AdminUsergroup) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AdminUsergroup provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), adminUsergroupPrimaryKeyMapping)
	sql := "DELETE FROM \"admin_usergroup\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from admin_usergroup")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for admin_usergroup")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q adminUsergroupQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no adminUsergroupQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from admin_usergroup")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for admin_usergroup")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o AdminUsergroupSlice) DeleteAllG() (int64, error) {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AdminUsergroupSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AdminUsergroup slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(adminUsergroupBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), adminUsergroupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"admin_usergroup\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, adminUsergroupPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from adminUsergroup slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for admin_usergroup")
	}

	if len(adminUsergroupAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *AdminUsergroup) ReloadG() error {
	if o == nil {
		return errors.New("models: no AdminUsergroup provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AdminUsergroup) Reload(exec boil.Executor) error {
	ret, err := FindAdminUsergroup(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AdminUsergroupSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty AdminUsergroupSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AdminUsergroupSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AdminUsergroupSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), adminUsergroupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"admin_usergroup\".* FROM \"admin_usergroup\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, adminUsergroupPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AdminUsergroupSlice")
	}

	*o = slice

	return nil
}

// AdminUsergroupExistsG checks if the AdminUsergroup row exists.
func AdminUsergroupExistsG(iD int) (bool, error) {
	return AdminUsergroupExists(boil.GetDB(), iD)
}

// AdminUsergroupExists checks if the AdminUsergroup row exists.
func AdminUsergroupExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"admin_usergroup\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if admin_usergroup exists")
	}

	return exists, nil
}

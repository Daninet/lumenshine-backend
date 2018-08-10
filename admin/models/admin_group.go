// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// AdminGroup is an object representing the database table.
type AdminGroup struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	UpdatedBy string    `boil:"updated_by" json:"updated_by" toml:"updated_by" yaml:"updated_by"`

	R *adminGroupR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L adminGroupL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AdminGroupColumns = struct {
	ID        string
	Name      string
	CreatedAt string
	UpdatedAt string
	UpdatedBy string
}{
	ID:        "id",
	Name:      "name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	UpdatedBy: "updated_by",
}

// adminGroupR is where relationships are stored.
type adminGroupR struct {
	GroupAdminUsergroups AdminUsergroupSlice
}

// adminGroupL is where Load methods for each relationship are stored.
type adminGroupL struct{}

var (
	adminGroupColumns               = []string{"id", "name", "created_at", "updated_at", "updated_by"}
	adminGroupColumnsWithoutDefault = []string{"name", "updated_by"}
	adminGroupColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	adminGroupPrimaryKeyColumns     = []string{"id"}
)

type (
	// AdminGroupSlice is an alias for a slice of pointers to AdminGroup.
	// This should generally be used opposed to []AdminGroup.
	AdminGroupSlice []*AdminGroup
	// AdminGroupHook is the signature for custom AdminGroup hook methods
	AdminGroupHook func(boil.Executor, *AdminGroup) error

	adminGroupQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	adminGroupType                 = reflect.TypeOf(&AdminGroup{})
	adminGroupMapping              = queries.MakeStructMapping(adminGroupType)
	adminGroupPrimaryKeyMapping, _ = queries.BindMapping(adminGroupType, adminGroupMapping, adminGroupPrimaryKeyColumns)
	adminGroupInsertCacheMut       sync.RWMutex
	adminGroupInsertCache          = make(map[string]insertCache)
	adminGroupUpdateCacheMut       sync.RWMutex
	adminGroupUpdateCache          = make(map[string]updateCache)
	adminGroupUpsertCacheMut       sync.RWMutex
	adminGroupUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var adminGroupBeforeInsertHooks []AdminGroupHook
var adminGroupBeforeUpdateHooks []AdminGroupHook
var adminGroupBeforeDeleteHooks []AdminGroupHook
var adminGroupBeforeUpsertHooks []AdminGroupHook

var adminGroupAfterInsertHooks []AdminGroupHook
var adminGroupAfterSelectHooks []AdminGroupHook
var adminGroupAfterUpdateHooks []AdminGroupHook
var adminGroupAfterDeleteHooks []AdminGroupHook
var adminGroupAfterUpsertHooks []AdminGroupHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AdminGroup) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range adminGroupBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AdminGroup) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range adminGroupBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AdminGroup) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range adminGroupBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AdminGroup) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range adminGroupBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AdminGroup) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range adminGroupAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AdminGroup) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range adminGroupAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AdminGroup) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range adminGroupAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AdminGroup) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range adminGroupAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AdminGroup) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range adminGroupAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAdminGroupHook registers your hook function for all future operations.
func AddAdminGroupHook(hookPoint boil.HookPoint, adminGroupHook AdminGroupHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		adminGroupBeforeInsertHooks = append(adminGroupBeforeInsertHooks, adminGroupHook)
	case boil.BeforeUpdateHook:
		adminGroupBeforeUpdateHooks = append(adminGroupBeforeUpdateHooks, adminGroupHook)
	case boil.BeforeDeleteHook:
		adminGroupBeforeDeleteHooks = append(adminGroupBeforeDeleteHooks, adminGroupHook)
	case boil.BeforeUpsertHook:
		adminGroupBeforeUpsertHooks = append(adminGroupBeforeUpsertHooks, adminGroupHook)
	case boil.AfterInsertHook:
		adminGroupAfterInsertHooks = append(adminGroupAfterInsertHooks, adminGroupHook)
	case boil.AfterSelectHook:
		adminGroupAfterSelectHooks = append(adminGroupAfterSelectHooks, adminGroupHook)
	case boil.AfterUpdateHook:
		adminGroupAfterUpdateHooks = append(adminGroupAfterUpdateHooks, adminGroupHook)
	case boil.AfterDeleteHook:
		adminGroupAfterDeleteHooks = append(adminGroupAfterDeleteHooks, adminGroupHook)
	case boil.AfterUpsertHook:
		adminGroupAfterUpsertHooks = append(adminGroupAfterUpsertHooks, adminGroupHook)
	}
}

// OneP returns a single adminGroup record from the query, and panics on error.
func (q adminGroupQuery) OneP() *AdminGroup {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single adminGroup record from the query.
func (q adminGroupQuery) One() (*AdminGroup, error) {
	o := &AdminGroup{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for admin_group")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all AdminGroup records from the query, and panics on error.
func (q adminGroupQuery) AllP() AdminGroupSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all AdminGroup records from the query.
func (q adminGroupQuery) All() (AdminGroupSlice, error) {
	var o []*AdminGroup

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AdminGroup slice")
	}

	if len(adminGroupAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all AdminGroup records in the query, and panics on error.
func (q adminGroupQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all AdminGroup records in the query.
func (q adminGroupQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count admin_group rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q adminGroupQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q adminGroupQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if admin_group exists")
	}

	return count > 0, nil
}

// GroupAdminUsergroupsG retrieves all the admin_usergroup's admin usergroup via group_id column.
func (o *AdminGroup) GroupAdminUsergroupsG(mods ...qm.QueryMod) adminUsergroupQuery {
	return o.GroupAdminUsergroups(boil.GetDB(), mods...)
}

// GroupAdminUsergroups retrieves all the admin_usergroup's admin usergroup with an executor via group_id column.
func (o *AdminGroup) GroupAdminUsergroups(exec boil.Executor, mods ...qm.QueryMod) adminUsergroupQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"admin_usergroup\".\"group_id\"=?", o.ID),
	)

	query := AdminUsergroups(exec, queryMods...)
	queries.SetFrom(query.Query, "\"admin_usergroup\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"admin_usergroup\".*"})
	}

	return query
}

// LoadGroupAdminUsergroups allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (adminGroupL) LoadGroupAdminUsergroups(e boil.Executor, singular bool, maybeAdminGroup interface{}) error {
	var slice []*AdminGroup
	var object *AdminGroup

	count := 1
	if singular {
		object = maybeAdminGroup.(*AdminGroup)
	} else {
		slice = *maybeAdminGroup.(*[]*AdminGroup)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &adminGroupR{}
		}
		args[0] = object.ID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &adminGroupR{}
			}
			args[i] = obj.ID
		}
	}

	query := fmt.Sprintf(
		"select * from \"admin_usergroup\" where \"group_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load admin_usergroup")
	}
	defer results.Close()

	var resultSlice []*AdminUsergroup
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice admin_usergroup")
	}

	if len(adminUsergroupAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.GroupAdminUsergroups = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.GroupID {
				local.R.GroupAdminUsergroups = append(local.R.GroupAdminUsergroups, foreign)
				break
			}
		}
	}

	return nil
}

// AddGroupAdminUsergroupsG adds the given related objects to the existing relationships
// of the admin_group, optionally inserting them as new records.
// Appends related to o.R.GroupAdminUsergroups.
// Sets related.R.Group appropriately.
// Uses the global database handle.
func (o *AdminGroup) AddGroupAdminUsergroupsG(insert bool, related ...*AdminUsergroup) error {
	return o.AddGroupAdminUsergroups(boil.GetDB(), insert, related...)
}

// AddGroupAdminUsergroupsP adds the given related objects to the existing relationships
// of the admin_group, optionally inserting them as new records.
// Appends related to o.R.GroupAdminUsergroups.
// Sets related.R.Group appropriately.
// Panics on error.
func (o *AdminGroup) AddGroupAdminUsergroupsP(exec boil.Executor, insert bool, related ...*AdminUsergroup) {
	if err := o.AddGroupAdminUsergroups(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddGroupAdminUsergroupsGP adds the given related objects to the existing relationships
// of the admin_group, optionally inserting them as new records.
// Appends related to o.R.GroupAdminUsergroups.
// Sets related.R.Group appropriately.
// Uses the global database handle and panics on error.
func (o *AdminGroup) AddGroupAdminUsergroupsGP(insert bool, related ...*AdminUsergroup) {
	if err := o.AddGroupAdminUsergroups(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddGroupAdminUsergroups adds the given related objects to the existing relationships
// of the admin_group, optionally inserting them as new records.
// Appends related to o.R.GroupAdminUsergroups.
// Sets related.R.Group appropriately.
func (o *AdminGroup) AddGroupAdminUsergroups(exec boil.Executor, insert bool, related ...*AdminUsergroup) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.GroupID = o.ID
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"admin_usergroup\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"group_id"}),
				strmangle.WhereClause("\"", "\"", 2, adminUsergroupPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.GroupID = o.ID
		}
	}

	if o.R == nil {
		o.R = &adminGroupR{
			GroupAdminUsergroups: related,
		}
	} else {
		o.R.GroupAdminUsergroups = append(o.R.GroupAdminUsergroups, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &adminUsergroupR{
				Group: o,
			}
		} else {
			rel.R.Group = o
		}
	}
	return nil
}

// AdminGroupsG retrieves all records.
func AdminGroupsG(mods ...qm.QueryMod) adminGroupQuery {
	return AdminGroups(boil.GetDB(), mods...)
}

// AdminGroups retrieves all the records using an executor.
func AdminGroups(exec boil.Executor, mods ...qm.QueryMod) adminGroupQuery {
	mods = append(mods, qm.From("\"admin_group\""))
	return adminGroupQuery{NewQuery(exec, mods...)}
}

// FindAdminGroupG retrieves a single record by ID.
func FindAdminGroupG(id int, selectCols ...string) (*AdminGroup, error) {
	return FindAdminGroup(boil.GetDB(), id, selectCols...)
}

// FindAdminGroupGP retrieves a single record by ID, and panics on error.
func FindAdminGroupGP(id int, selectCols ...string) *AdminGroup {
	retobj, err := FindAdminGroup(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindAdminGroup retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAdminGroup(exec boil.Executor, id int, selectCols ...string) (*AdminGroup, error) {
	adminGroupObj := &AdminGroup{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"admin_group\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(adminGroupObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from admin_group")
	}

	return adminGroupObj, nil
}

// FindAdminGroupP retrieves a single record by ID with an executor, and panics on error.
func FindAdminGroupP(exec boil.Executor, id int, selectCols ...string) *AdminGroup {
	retobj, err := FindAdminGroup(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *AdminGroup) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *AdminGroup) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *AdminGroup) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *AdminGroup) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no admin_group provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(adminGroupColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	adminGroupInsertCacheMut.RLock()
	cache, cached := adminGroupInsertCache[key]
	adminGroupInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			adminGroupColumns,
			adminGroupColumnsWithDefault,
			adminGroupColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(adminGroupType, adminGroupMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(adminGroupType, adminGroupMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"admin_group\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"admin_group\" DEFAULT VALUES"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
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
		return errors.Wrap(err, "models: unable to insert into admin_group")
	}

	if !cached {
		adminGroupInsertCacheMut.Lock()
		adminGroupInsertCache[key] = cache
		adminGroupInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single AdminGroup record. See Update for
// whitelist behavior description.
func (o *AdminGroup) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single AdminGroup record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *AdminGroup) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the AdminGroup, and panics on error.
// See Update for whitelist behavior description.
func (o *AdminGroup) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the AdminGroup.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *AdminGroup) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	adminGroupUpdateCacheMut.RLock()
	cache, cached := adminGroupUpdateCache[key]
	adminGroupUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			adminGroupColumns,
			adminGroupPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update admin_group, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"admin_group\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, adminGroupPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(adminGroupType, adminGroupMapping, append(wl, adminGroupPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update admin_group row")
	}

	if !cached {
		adminGroupUpdateCacheMut.Lock()
		adminGroupUpdateCache[key] = cache
		adminGroupUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q adminGroupQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q adminGroupQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for admin_group")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AdminGroupSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o AdminGroupSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o AdminGroupSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AdminGroupSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), adminGroupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"admin_group\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, adminGroupPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in adminGroup slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *AdminGroup) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *AdminGroup) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *AdminGroup) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *AdminGroup) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no admin_group provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(adminGroupColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
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
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	adminGroupUpsertCacheMut.RLock()
	cache, cached := adminGroupUpsertCache[key]
	adminGroupUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			adminGroupColumns,
			adminGroupColumnsWithDefault,
			adminGroupColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			adminGroupColumns,
			adminGroupPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert admin_group, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(adminGroupPrimaryKeyColumns))
			copy(conflict, adminGroupPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"admin_group\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(adminGroupType, adminGroupMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(adminGroupType, adminGroupMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert admin_group")
	}

	if !cached {
		adminGroupUpsertCacheMut.Lock()
		adminGroupUpsertCache[key] = cache
		adminGroupUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single AdminGroup record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *AdminGroup) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single AdminGroup record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *AdminGroup) DeleteG() error {
	if o == nil {
		return errors.New("models: no AdminGroup provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single AdminGroup record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *AdminGroup) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single AdminGroup record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AdminGroup) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no AdminGroup provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), adminGroupPrimaryKeyMapping)
	sql := "DELETE FROM \"admin_group\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from admin_group")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q adminGroupQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q adminGroupQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no adminGroupQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from admin_group")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o AdminGroupSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o AdminGroupSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no AdminGroup slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o AdminGroupSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AdminGroupSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no AdminGroup slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(adminGroupBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), adminGroupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"admin_group\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, adminGroupPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from adminGroup slice")
	}

	if len(adminGroupAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *AdminGroup) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *AdminGroup) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *AdminGroup) ReloadG() error {
	if o == nil {
		return errors.New("models: no AdminGroup provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AdminGroup) Reload(exec boil.Executor) error {
	ret, err := FindAdminGroup(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AdminGroupSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AdminGroupSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AdminGroupSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty AdminGroupSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AdminGroupSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	adminGroups := AdminGroupSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), adminGroupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"admin_group\".* FROM \"admin_group\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, adminGroupPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&adminGroups)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AdminGroupSlice")
	}

	*o = adminGroups

	return nil
}

// AdminGroupExists checks if the AdminGroup row exists.
func AdminGroupExists(exec boil.Executor, id int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"admin_group\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if admin_group exists")
	}

	return exists, nil
}

// AdminGroupExistsG checks if the AdminGroup row exists.
func AdminGroupExistsG(id int) (bool, error) {
	return AdminGroupExists(boil.GetDB(), id)
}

// AdminGroupExistsGP checks if the AdminGroup row exists. Panics on error.
func AdminGroupExistsGP(id int) bool {
	e, err := AdminGroupExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// AdminGroupExistsP checks if the AdminGroup row exists. Panics on error.
func AdminGroupExistsP(exec boil.Executor, id int) bool {
	e, err := AdminGroupExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

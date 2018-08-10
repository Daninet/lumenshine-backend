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

// Country is an object representing the database table.
type Country struct {
	ID          int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	LangCode    string `boil:"lang_code" json:"lang_code" toml:"lang_code" yaml:"lang_code"`
	CountryName string `boil:"country_name" json:"country_name" toml:"country_name" yaml:"country_name"`

	R *countryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L countryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CountryColumns = struct {
	ID          string
	LangCode    string
	CountryName string
}{
	ID:          "id",
	LangCode:    "lang_code",
	CountryName: "country_name",
}

// countryR is where relationships are stored.
type countryR struct {
}

// countryL is where Load methods for each relationship are stored.
type countryL struct{}

var (
	countryColumns               = []string{"id", "lang_code", "country_name"}
	countryColumnsWithoutDefault = []string{"lang_code", "country_name"}
	countryColumnsWithDefault    = []string{"id"}
	countryPrimaryKeyColumns     = []string{"id"}
)

type (
	// CountrySlice is an alias for a slice of pointers to Country.
	// This should generally be used opposed to []Country.
	CountrySlice []*Country
	// CountryHook is the signature for custom Country hook methods
	CountryHook func(boil.Executor, *Country) error

	countryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	countryType                 = reflect.TypeOf(&Country{})
	countryMapping              = queries.MakeStructMapping(countryType)
	countryPrimaryKeyMapping, _ = queries.BindMapping(countryType, countryMapping, countryPrimaryKeyColumns)
	countryInsertCacheMut       sync.RWMutex
	countryInsertCache          = make(map[string]insertCache)
	countryUpdateCacheMut       sync.RWMutex
	countryUpdateCache          = make(map[string]updateCache)
	countryUpsertCacheMut       sync.RWMutex
	countryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var countryBeforeInsertHooks []CountryHook
var countryBeforeUpdateHooks []CountryHook
var countryBeforeDeleteHooks []CountryHook
var countryBeforeUpsertHooks []CountryHook

var countryAfterInsertHooks []CountryHook
var countryAfterSelectHooks []CountryHook
var countryAfterUpdateHooks []CountryHook
var countryAfterDeleteHooks []CountryHook
var countryAfterUpsertHooks []CountryHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Country) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range countryBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Country) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range countryBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Country) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range countryBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Country) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range countryBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Country) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range countryAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Country) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range countryAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Country) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range countryAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Country) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range countryAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Country) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range countryAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCountryHook registers your hook function for all future operations.
func AddCountryHook(hookPoint boil.HookPoint, countryHook CountryHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		countryBeforeInsertHooks = append(countryBeforeInsertHooks, countryHook)
	case boil.BeforeUpdateHook:
		countryBeforeUpdateHooks = append(countryBeforeUpdateHooks, countryHook)
	case boil.BeforeDeleteHook:
		countryBeforeDeleteHooks = append(countryBeforeDeleteHooks, countryHook)
	case boil.BeforeUpsertHook:
		countryBeforeUpsertHooks = append(countryBeforeUpsertHooks, countryHook)
	case boil.AfterInsertHook:
		countryAfterInsertHooks = append(countryAfterInsertHooks, countryHook)
	case boil.AfterSelectHook:
		countryAfterSelectHooks = append(countryAfterSelectHooks, countryHook)
	case boil.AfterUpdateHook:
		countryAfterUpdateHooks = append(countryAfterUpdateHooks, countryHook)
	case boil.AfterDeleteHook:
		countryAfterDeleteHooks = append(countryAfterDeleteHooks, countryHook)
	case boil.AfterUpsertHook:
		countryAfterUpsertHooks = append(countryAfterUpsertHooks, countryHook)
	}
}

// OneP returns a single country record from the query, and panics on error.
func (q countryQuery) OneP() *Country {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single country record from the query.
func (q countryQuery) One() (*Country, error) {
	o := &Country{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for countries")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Country records from the query, and panics on error.
func (q countryQuery) AllP() CountrySlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Country records from the query.
func (q countryQuery) All() (CountrySlice, error) {
	var o []*Country

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Country slice")
	}

	if len(countryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Country records in the query, and panics on error.
func (q countryQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Country records in the query.
func (q countryQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count countries rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q countryQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q countryQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if countries exists")
	}

	return count > 0, nil
}

// CountriesG retrieves all records.
func CountriesG(mods ...qm.QueryMod) countryQuery {
	return Countries(boil.GetDB(), mods...)
}

// Countries retrieves all the records using an executor.
func Countries(exec boil.Executor, mods ...qm.QueryMod) countryQuery {
	mods = append(mods, qm.From("\"countries\""))
	return countryQuery{NewQuery(exec, mods...)}
}

// FindCountryG retrieves a single record by ID.
func FindCountryG(id int, selectCols ...string) (*Country, error) {
	return FindCountry(boil.GetDB(), id, selectCols...)
}

// FindCountryGP retrieves a single record by ID, and panics on error.
func FindCountryGP(id int, selectCols ...string) *Country {
	retobj, err := FindCountry(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindCountry retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCountry(exec boil.Executor, id int, selectCols ...string) (*Country, error) {
	countryObj := &Country{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"countries\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(countryObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from countries")
	}

	return countryObj, nil
}

// FindCountryP retrieves a single record by ID with an executor, and panics on error.
func FindCountryP(exec boil.Executor, id int, selectCols ...string) *Country {
	retobj, err := FindCountry(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Country) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Country) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Country) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Country) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no countries provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(countryColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	countryInsertCacheMut.RLock()
	cache, cached := countryInsertCache[key]
	countryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			countryColumns,
			countryColumnsWithDefault,
			countryColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(countryType, countryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(countryType, countryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"countries\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"countries\" DEFAULT VALUES"
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
		return errors.Wrap(err, "models: unable to insert into countries")
	}

	if !cached {
		countryInsertCacheMut.Lock()
		countryInsertCache[key] = cache
		countryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Country record. See Update for
// whitelist behavior description.
func (o *Country) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Country record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Country) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Country, and panics on error.
// See Update for whitelist behavior description.
func (o *Country) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Country.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Country) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	countryUpdateCacheMut.RLock()
	cache, cached := countryUpdateCache[key]
	countryUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			countryColumns,
			countryPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update countries, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"countries\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, countryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(countryType, countryMapping, append(wl, countryPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update countries row")
	}

	if !cached {
		countryUpdateCacheMut.Lock()
		countryUpdateCache[key] = cache
		countryUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q countryQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q countryQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for countries")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o CountrySlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o CountrySlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o CountrySlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CountrySlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), countryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"countries\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, countryPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in country slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Country) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Country) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Country) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Country) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no countries provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(countryColumnsWithDefault, o)

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

	countryUpsertCacheMut.RLock()
	cache, cached := countryUpsertCache[key]
	countryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			countryColumns,
			countryColumnsWithDefault,
			countryColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			countryColumns,
			countryPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert countries, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(countryPrimaryKeyColumns))
			copy(conflict, countryPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"countries\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(countryType, countryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(countryType, countryMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert countries")
	}

	if !cached {
		countryUpsertCacheMut.Lock()
		countryUpsertCache[key] = cache
		countryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Country record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Country) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Country record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Country) DeleteG() error {
	if o == nil {
		return errors.New("models: no Country provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Country record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Country) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Country record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Country) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Country provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), countryPrimaryKeyMapping)
	sql := "DELETE FROM \"countries\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from countries")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q countryQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q countryQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no countryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from countries")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o CountrySlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o CountrySlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Country slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o CountrySlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CountrySlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Country slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(countryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), countryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"countries\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, countryPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from country slice")
	}

	if len(countryAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Country) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Country) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Country) ReloadG() error {
	if o == nil {
		return errors.New("models: no Country provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Country) Reload(exec boil.Executor) error {
	ret, err := FindCountry(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CountrySlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CountrySlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CountrySlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty CountrySlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CountrySlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	countries := CountrySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), countryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"countries\".* FROM \"countries\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, countryPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&countries)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CountrySlice")
	}

	*o = countries

	return nil
}

// CountryExists checks if the Country row exists.
func CountryExists(exec boil.Executor, id int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"countries\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if countries exists")
	}

	return exists, nil
}

// CountryExistsG checks if the Country row exists.
func CountryExistsG(id int) (bool, error) {
	return CountryExists(boil.GetDB(), id)
}

// CountryExistsGP checks if the Country row exists. Panics on error.
func CountryExistsGP(id int) bool {
	e, err := CountryExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// CountryExistsP checks if the Country row exists. Panics on error.
func CountryExistsP(exec boil.Executor, id int) bool {
	e, err := CountryExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

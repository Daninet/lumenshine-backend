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

// Ico is an object representing the database table.
type Ico struct {
	ID         int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	IcoName    string    `boil:"ico_name" json:"ico_name" toml:"ico_name" yaml:"ico_name"`
	IcoStatus  string    `boil:"ico_status" json:"ico_status" toml:"ico_status" yaml:"ico_status"`
	Kyc        bool      `boil:"kyc" json:"kyc" toml:"kyc" yaml:"kyc"`
	SalesModel string    `boil:"sales_model" json:"sales_model" toml:"sales_model" yaml:"sales_model"`
	IssuerPK   string    `boil:"issuer_pk" json:"issuer_pk" toml:"issuer_pk" yaml:"issuer_pk"`
	AssetCode  string    `boil:"asset_code" json:"asset_code" toml:"asset_code" yaml:"asset_code"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	UpdatedBy  string    `boil:"updated_by" json:"updated_by" toml:"updated_by" yaml:"updated_by"`

	R *icoR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L icoL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var IcoColumns = struct {
	ID         string
	IcoName    string
	IcoStatus  string
	Kyc        string
	SalesModel string
	IssuerPK   string
	AssetCode  string
	CreatedAt  string
	UpdatedAt  string
	UpdatedBy  string
}{
	ID:         "id",
	IcoName:    "ico_name",
	IcoStatus:  "ico_status",
	Kyc:        "kyc",
	SalesModel: "sales_model",
	IssuerPK:   "issuer_pk",
	AssetCode:  "asset_code",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	UpdatedBy:  "updated_by",
}

// IcoRels is where relationship names are stored.
var IcoRels = struct {
	IcoPhases                      string
	IcoSupportedExchangeCurrencies string
}{
	IcoPhases:                      "IcoPhases",
	IcoSupportedExchangeCurrencies: "IcoSupportedExchangeCurrencies",
}

// icoR is where relationships are stored.
type icoR struct {
	IcoPhases                      IcoPhaseSlice
	IcoSupportedExchangeCurrencies IcoSupportedExchangeCurrencySlice
}

// NewStruct creates a new relationship struct
func (*icoR) NewStruct() *icoR {
	return &icoR{}
}

// icoL is where Load methods for each relationship are stored.
type icoL struct{}

var (
	icoColumns               = []string{"id", "ico_name", "ico_status", "kyc", "sales_model", "issuer_pk", "asset_code", "created_at", "updated_at", "updated_by"}
	icoColumnsWithoutDefault = []string{"ico_name", "ico_status", "kyc", "sales_model", "issuer_pk", "asset_code", "updated_by"}
	icoColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	icoPrimaryKeyColumns     = []string{"id"}
)

type (
	// IcoSlice is an alias for a slice of pointers to Ico.
	// This should generally be used opposed to []Ico.
	IcoSlice []*Ico
	// IcoHook is the signature for custom Ico hook methods
	IcoHook func(boil.Executor, *Ico) error

	icoQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	icoType                 = reflect.TypeOf(&Ico{})
	icoMapping              = queries.MakeStructMapping(icoType)
	icoPrimaryKeyMapping, _ = queries.BindMapping(icoType, icoMapping, icoPrimaryKeyColumns)
	icoInsertCacheMut       sync.RWMutex
	icoInsertCache          = make(map[string]insertCache)
	icoUpdateCacheMut       sync.RWMutex
	icoUpdateCache          = make(map[string]updateCache)
	icoUpsertCacheMut       sync.RWMutex
	icoUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var icoBeforeInsertHooks []IcoHook
var icoBeforeUpdateHooks []IcoHook
var icoBeforeDeleteHooks []IcoHook
var icoBeforeUpsertHooks []IcoHook

var icoAfterInsertHooks []IcoHook
var icoAfterSelectHooks []IcoHook
var icoAfterUpdateHooks []IcoHook
var icoAfterDeleteHooks []IcoHook
var icoAfterUpsertHooks []IcoHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Ico) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range icoBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Ico) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range icoBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Ico) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range icoBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Ico) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range icoBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Ico) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range icoAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Ico) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range icoAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Ico) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range icoAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Ico) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range icoAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Ico) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range icoAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddIcoHook registers your hook function for all future operations.
func AddIcoHook(hookPoint boil.HookPoint, icoHook IcoHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		icoBeforeInsertHooks = append(icoBeforeInsertHooks, icoHook)
	case boil.BeforeUpdateHook:
		icoBeforeUpdateHooks = append(icoBeforeUpdateHooks, icoHook)
	case boil.BeforeDeleteHook:
		icoBeforeDeleteHooks = append(icoBeforeDeleteHooks, icoHook)
	case boil.BeforeUpsertHook:
		icoBeforeUpsertHooks = append(icoBeforeUpsertHooks, icoHook)
	case boil.AfterInsertHook:
		icoAfterInsertHooks = append(icoAfterInsertHooks, icoHook)
	case boil.AfterSelectHook:
		icoAfterSelectHooks = append(icoAfterSelectHooks, icoHook)
	case boil.AfterUpdateHook:
		icoAfterUpdateHooks = append(icoAfterUpdateHooks, icoHook)
	case boil.AfterDeleteHook:
		icoAfterDeleteHooks = append(icoAfterDeleteHooks, icoHook)
	case boil.AfterUpsertHook:
		icoAfterUpsertHooks = append(icoAfterUpsertHooks, icoHook)
	}
}

// OneG returns a single ico record from the query using the global executor.
func (q icoQuery) OneG() (*Ico, error) {
	return q.One(boil.GetDB())
}

// One returns a single ico record from the query.
func (q icoQuery) One(exec boil.Executor) (*Ico, error) {
	o := &Ico{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for ico")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Ico records from the query using the global executor.
func (q icoQuery) AllG() (IcoSlice, error) {
	return q.All(boil.GetDB())
}

// All returns all Ico records from the query.
func (q icoQuery) All(exec boil.Executor) (IcoSlice, error) {
	var o []*Ico

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Ico slice")
	}

	if len(icoAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Ico records in the query, and panics on error.
func (q icoQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all Ico records in the query.
func (q icoQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count ico rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q icoQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q icoQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if ico exists")
	}

	return count > 0, nil
}

// IcoPhases retrieves all the ico_phase's IcoPhases with an executor.
func (o *Ico) IcoPhases(mods ...qm.QueryMod) icoPhaseQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"ico_phase\".\"ico_id\"=?", o.ID),
	)

	query := IcoPhases(queryMods...)
	queries.SetFrom(query.Query, "\"ico_phase\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"ico_phase\".*"})
	}

	return query
}

// IcoSupportedExchangeCurrencies retrieves all the ico_supported_exchange_currency's IcoSupportedExchangeCurrencies with an executor.
func (o *Ico) IcoSupportedExchangeCurrencies(mods ...qm.QueryMod) icoSupportedExchangeCurrencyQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"ico_supported_exchange_currency\".\"ico_id\"=?", o.ID),
	)

	query := IcoSupportedExchangeCurrencies(queryMods...)
	queries.SetFrom(query.Query, "\"ico_supported_exchange_currency\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"ico_supported_exchange_currency\".*"})
	}

	return query
}

// LoadIcoPhases allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (icoL) LoadIcoPhases(e boil.Executor, singular bool, maybeIco interface{}, mods queries.Applicator) error {
	var slice []*Ico
	var object *Ico

	if singular {
		object = maybeIco.(*Ico)
	} else {
		slice = *maybeIco.(*[]*Ico)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &icoR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &icoR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	query := NewQuery(qm.From(`ico_phase`), qm.WhereIn(`ico_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ico_phase")
	}

	var resultSlice []*IcoPhase
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ico_phase")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on ico_phase")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for ico_phase")
	}

	if len(icoPhaseAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.IcoPhases = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &icoPhaseR{}
			}
			foreign.R.Ico = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.IcoID {
				local.R.IcoPhases = append(local.R.IcoPhases, foreign)
				if foreign.R == nil {
					foreign.R = &icoPhaseR{}
				}
				foreign.R.Ico = local
				break
			}
		}
	}

	return nil
}

// LoadIcoSupportedExchangeCurrencies allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (icoL) LoadIcoSupportedExchangeCurrencies(e boil.Executor, singular bool, maybeIco interface{}, mods queries.Applicator) error {
	var slice []*Ico
	var object *Ico

	if singular {
		object = maybeIco.(*Ico)
	} else {
		slice = *maybeIco.(*[]*Ico)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &icoR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &icoR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	query := NewQuery(qm.From(`ico_supported_exchange_currency`), qm.WhereIn(`ico_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ico_supported_exchange_currency")
	}

	var resultSlice []*IcoSupportedExchangeCurrency
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ico_supported_exchange_currency")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on ico_supported_exchange_currency")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for ico_supported_exchange_currency")
	}

	if len(icoSupportedExchangeCurrencyAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.IcoSupportedExchangeCurrencies = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &icoSupportedExchangeCurrencyR{}
			}
			foreign.R.Ico = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.IcoID {
				local.R.IcoSupportedExchangeCurrencies = append(local.R.IcoSupportedExchangeCurrencies, foreign)
				if foreign.R == nil {
					foreign.R = &icoSupportedExchangeCurrencyR{}
				}
				foreign.R.Ico = local
				break
			}
		}
	}

	return nil
}

// AddIcoPhasesG adds the given related objects to the existing relationships
// of the ico, optionally inserting them as new records.
// Appends related to o.R.IcoPhases.
// Sets related.R.Ico appropriately.
// Uses the global database handle.
func (o *Ico) AddIcoPhasesG(insert bool, related ...*IcoPhase) error {
	return o.AddIcoPhases(boil.GetDB(), insert, related...)
}

// AddIcoPhases adds the given related objects to the existing relationships
// of the ico, optionally inserting them as new records.
// Appends related to o.R.IcoPhases.
// Sets related.R.Ico appropriately.
func (o *Ico) AddIcoPhases(exec boil.Executor, insert bool, related ...*IcoPhase) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.IcoID = o.ID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"ico_phase\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"ico_id"}),
				strmangle.WhereClause("\"", "\"", 2, icoPhasePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.IcoID = o.ID
		}
	}

	if o.R == nil {
		o.R = &icoR{
			IcoPhases: related,
		}
	} else {
		o.R.IcoPhases = append(o.R.IcoPhases, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &icoPhaseR{
				Ico: o,
			}
		} else {
			rel.R.Ico = o
		}
	}
	return nil
}

// AddIcoSupportedExchangeCurrenciesG adds the given related objects to the existing relationships
// of the ico, optionally inserting them as new records.
// Appends related to o.R.IcoSupportedExchangeCurrencies.
// Sets related.R.Ico appropriately.
// Uses the global database handle.
func (o *Ico) AddIcoSupportedExchangeCurrenciesG(insert bool, related ...*IcoSupportedExchangeCurrency) error {
	return o.AddIcoSupportedExchangeCurrencies(boil.GetDB(), insert, related...)
}

// AddIcoSupportedExchangeCurrencies adds the given related objects to the existing relationships
// of the ico, optionally inserting them as new records.
// Appends related to o.R.IcoSupportedExchangeCurrencies.
// Sets related.R.Ico appropriately.
func (o *Ico) AddIcoSupportedExchangeCurrencies(exec boil.Executor, insert bool, related ...*IcoSupportedExchangeCurrency) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.IcoID = o.ID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"ico_supported_exchange_currency\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"ico_id"}),
				strmangle.WhereClause("\"", "\"", 2, icoSupportedExchangeCurrencyPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.IcoID = o.ID
		}
	}

	if o.R == nil {
		o.R = &icoR{
			IcoSupportedExchangeCurrencies: related,
		}
	} else {
		o.R.IcoSupportedExchangeCurrencies = append(o.R.IcoSupportedExchangeCurrencies, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &icoSupportedExchangeCurrencyR{
				Ico: o,
			}
		} else {
			rel.R.Ico = o
		}
	}
	return nil
}

// Icos retrieves all the records using an executor.
func Icos(mods ...qm.QueryMod) icoQuery {
	mods = append(mods, qm.From("\"ico\""))
	return icoQuery{NewQuery(mods...)}
}

// FindIcoG retrieves a single record by ID.
func FindIcoG(iD int, selectCols ...string) (*Ico, error) {
	return FindIco(boil.GetDB(), iD, selectCols...)
}

// FindIco retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindIco(exec boil.Executor, iD int, selectCols ...string) (*Ico, error) {
	icoObj := &Ico{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"ico\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, icoObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from ico")
	}

	return icoObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Ico) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Ico) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no ico provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(icoColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	icoInsertCacheMut.RLock()
	cache, cached := icoInsertCache[key]
	icoInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			icoColumns,
			icoColumnsWithDefault,
			icoColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(icoType, icoMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(icoType, icoMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"ico\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"ico\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into ico")
	}

	if !cached {
		icoInsertCacheMut.Lock()
		icoInsertCache[key] = cache
		icoInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Ico record using the global executor.
// See Update for more documentation.
func (o *Ico) UpdateG(columns boil.Columns) (int64, error) {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the Ico.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Ico) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	icoUpdateCacheMut.RLock()
	cache, cached := icoUpdateCache[key]
	icoUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			icoColumns,
			icoPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update ico, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"ico\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, icoPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(icoType, icoMapping, append(wl, icoPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update ico row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for ico")
	}

	if !cached {
		icoUpdateCacheMut.Lock()
		icoUpdateCache[key] = cache
		icoUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAll updates all rows with the specified column values.
func (q icoQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for ico")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for ico")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o IcoSlice) UpdateAllG(cols M) (int64, error) {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o IcoSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), icoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"ico\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, icoPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in ico slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all ico")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Ico) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Ico) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no ico provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(icoColumnsWithDefault, o)

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

	icoUpsertCacheMut.RLock()
	cache, cached := icoUpsertCache[key]
	icoUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			icoColumns,
			icoColumnsWithDefault,
			icoColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			icoColumns,
			icoPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert ico, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(icoPrimaryKeyColumns))
			copy(conflict, icoPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"ico\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(icoType, icoMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(icoType, icoMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert ico")
	}

	if !cached {
		icoUpsertCacheMut.Lock()
		icoUpsertCache[key] = cache
		icoUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteG deletes a single Ico record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Ico) DeleteG() (int64, error) {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single Ico record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Ico) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Ico provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), icoPrimaryKeyMapping)
	sql := "DELETE FROM \"ico\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from ico")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for ico")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q icoQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no icoQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from ico")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for ico")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o IcoSlice) DeleteAllG() (int64, error) {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o IcoSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Ico slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(icoBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), icoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"ico\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, icoPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from ico slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for ico")
	}

	if len(icoAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Ico) ReloadG() error {
	if o == nil {
		return errors.New("models: no Ico provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Ico) Reload(exec boil.Executor) error {
	ret, err := FindIco(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *IcoSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty IcoSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *IcoSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := IcoSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), icoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"ico\".* FROM \"ico\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, icoPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in IcoSlice")
	}

	*o = slice

	return nil
}

// IcoExistsG checks if the Ico row exists.
func IcoExistsG(iD int) (bool, error) {
	return IcoExists(boil.GetDB(), iD)
}

// IcoExists checks if the Ico row exists.
func IcoExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"ico\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if ico exists")
	}

	return exists, nil
}

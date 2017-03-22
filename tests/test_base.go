package tests

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/three-plus-three/web_example/app"

	fixtures "github.com/AreaHQ/go-fixtures"
	"github.com/Masterminds/squirrel"
	"github.com/revel/revel"
	"github.com/revel/revel/testing"
)

// DbRunner wraps sql.DB to implement Runner.
type DbRunner struct {
	*sql.DB
}

// QueryRow wraps QueryRow to implement RowScanner.
func (r DbRunner) QueryRow(query string, args ...interface{}) squirrel.RowScanner {
	return r.DB.QueryRow(query, args...)
}

type EQU map[string]interface{}

type BaseTest struct {
	testing.TestSuite
}

func (t *BaseTest) Before() {
	println("================ Set up  =================")
	fmt.Println(app.Lifecycle.Env.Db.Models.Url())
	if !strings.Contains(app.Lifecycle.Env.Db.Models.Schema, "_test") {
		panic("runMode must is test.")
	}
}

func (t *BaseTest) After() {
	println("=============== Tear down ================")
}

func (t *BaseTest) DB() *sql.DB {
	return app.Lifecycle.DB.Engine.DB().DB
}

func (t *BaseTest) DataDB() *sql.DB {
	return app.Lifecycle.DataDB.Engine.DB().DB
}

func (t *BaseTest) DBRunable() squirrel.Runner {
	return &DbRunner{t.DB()}
}

func (t *BaseTest) DataDBRunable() squirrel.Runner {
	return &DbRunner{t.DataDB()}
}

func (t *BaseTest) ReverseUrl(args ...interface{}) string {
	s, e := revel.ReverseUrl(args...)
	if e != nil {
		t.Assertf(false, e.Error())
		return ""
	}
	return string(s)
}

func (t *BaseTest) LoadFiles(filenames ...string) {
	if err := fixtures.LoadFiles(filenames, t.DB(), "postgres"); err != nil {
		t.Assertf(false, err.Error())
		return
	}
}

func (t *BaseTest) LoadFilesToData(filenames ...string) {
	if err := fixtures.LoadFiles(filenames, t.DataDB(), "postgres"); err != nil {
		t.Assertf(false, err.Error())
		return
	}
}

func (t *BaseTest) GetCountFromTable(table string, params EQU) (count int64) {
	return t.getCountFromTable(t.DBRunable(), table, params)
}

func (t *BaseTest) GetCountFromDataTable(table string, params EQU) (count int64) {
	return t.getCountFromTable(t.DataDBRunable(), table, params)
}

func (t *BaseTest) getCountFromTable(dbRunner squirrel.Runner, table string, params EQU) (count int64) {
	builder := squirrel.Select("count(*)").From(table)
	if len(params) > 0 {
		builder = builder.Where(squirrel.Eq(params))
	}
	builder = builder.PlaceholderFormat(squirrel.Dollar)

	fmt.Println(builder.ToSql())
	rs := squirrel.QueryRowWith(dbRunner, builder)
	if err := rs.Scan(&count); nil != err {
		t.Assertf(false, err.Error())
	}
	return count
}

func (t *BaseTest) GetIDFromTable(table string, params EQU) (id int64) {
	return t.getIDFromTable(t.DBRunable(), table, params)
}

func (t *BaseTest) GetIDFromDataTable(table string, params EQU) (id int64) {
	return t.getIDFromTable(t.DataDBRunable(), table, params)
}

func (t *BaseTest) getIDFromTable(dbRunner squirrel.Runner, table string, params EQU) (id int64) {
	builder := squirrel.Select("id").From(table)
	if len(params) > 0 {
		builder = builder.Where(squirrel.Eq(params))
	}
	builder = builder.PlaceholderFormat(squirrel.Dollar)
	rs := squirrel.QueryRowWith(dbRunner, builder)
	if err := rs.Scan(&id); nil != err {
		t.Assertf(false, err.Error())
	}
	return id
}

func (t *BaseTest) ClearTable(tableName string) {
	t.clearTable(t.DBRunable(), tableName)
}

func (t *BaseTest) ClearDataTable(tableName string) {
	t.clearTable(t.DataDBRunable(), tableName)
}

func (t *BaseTest) clearTable(dbRunner squirrel.Runner, tableName string) {
	if _, err := dbRunner.Exec("truncate table " + tableName + " cascade"); err != nil {
		t.Assertf(false, err.Error())
	}
}

func (t *BaseTest) ClearDB() {
	if _, err := t.DB().Exec("select clear_data_of_all_table()"); err != nil {
		t.Assertf(false, err.Error())
	}
}

func (t *BaseTest) ResponseAsJSONObject() map[string]interface{} {
	var res map[string]interface{}
	if err := json.Unmarshal(t.ResponseBody, &res); err != nil {
		t.Assertf(false, err.Error())
	}
	return res
}

func (t *BaseTest) ResponseAsJSONArray() []map[string]interface{} {
	var res []map[string]interface{}
	if err := json.Unmarshal(t.ResponseBody, &res); err != nil {
		t.Assertf(false, err.Error())
	}
	return res
}

func (t *BaseTest) ResponseAsArray() []interface{} {
	var res []interface{}
	if err := json.Unmarshal(t.ResponseBody, &res); err != nil {
		t.Assertf(false, err.Error())
	}
	return res
}

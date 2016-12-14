package api

import (
	_ "errors"
	"fmt"
	"github.com/bigokro/gruff-server/gruff"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"os"
)

var CTX *Context
var INITDB *gorm.DB
var TESTDB *gorm.DB

var TESTTOKEN string
var READ_ONLY bool = false

func init() {
	INITDB = InitTestDB()
}

func setup() {
	TESTDB = INITDB.Begin()
	TESTDB.Callback().Create().Replace("fail_on_ro_update", failOnReadOnlyUpdate)
	TESTDB.Callback().Update().Replace("fail_on_ro_update", failOnReadOnlyUpdate)
	TESTDB.Callback().Delete().Replace("fail_on_ro_update", failOnReadOnlyUpdate)

	if CTX == nil {
		CTX = &Context{}
	}

	CTX.Database = TESTDB
}

func teardown() {
	TESTDB = TESTDB.Rollback()
}

func failOnReadOnlyUpdate(scope *gorm.Scope) {
	if READ_ONLY {
		panic("Performing an update on read-only transaction!")
	}
}

func Router() *echo.Echo {
	return SetUpRouter(true, TESTDB)
}

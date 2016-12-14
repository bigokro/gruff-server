package api

import (
	_ "errors"
	"github.com/bigokro/gruff-server/gruff"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

var CTX *Context
var INITDB *gorm.DB
var TESTDB *gorm.DB

var TESTTOKEN string
var READ_ONLY bool = false

func init() {
	INITDB = gruff.InitTestDB()
}

func setup() {
	TESTDB = INITDB.Begin()

	if CTX == nil {
		CTX = &Context{}
	}

	CTX.Database = TESTDB
}

func teardown() {
	TESTDB = TESTDB.Rollback()
}

func Router() *echo.Echo {
	return SetUpRouter(true, TESTDB)
}

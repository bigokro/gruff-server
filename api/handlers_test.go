package api

import (
	model "../model"
	_ "errors"
	"fmt"
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

func InitTestDB() *gorm.DB {
	if os.Getenv("GRUFF_DB") == "" {
		os.Setenv("GRUFF_DB", "dbname=gruff_test sslmode=disable")
	}

	var err error
	var db *gorm.DB
	if db, err = OpenTestConnection(); err != nil {
		fmt.Println("No error should happen when connecting to test database, but got", err)
	}

	db.LogMode(false)

	db.DB().SetMaxIdleConns(10)

	runMigration(db)

	return db
}

func OpenTestConnection() (db *gorm.DB, err error) {
	gruff_db := os.Getenv("GRUFF_DB")
	if gruff_db == "" {
		gruff_db = "dbname=gruff_test sslmode=disable"
	}
	db, err = gorm.Open("postgres", gruff_db)
	return
}

func runMigration(db *gorm.DB) {
	values := []interface{}{
		&mondo.User{},
		&mondo.Debate{},
		&mondo.DebateOpinion{},
		&mondo.Argument{},
		&mondo.ArgumentOpinion{},
		&mondo.Reference{},
		&mondo.Tag{},
		&mondo.Context{},
		&mondo.Value{},
	}

	for _, value := range values {
		db.DropTable(value)
	}

	if err := db.AutoMigrate(values...).Error; err != nil {
		panic(fmt.Sprintf("No error should happen when create table, but got %+v", err))
	}
}

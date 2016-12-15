package gruff

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var INITDB *gorm.DB
var TESTDB *gorm.DB

func init() {
	INITDB = InitTestDB()
}

func setupDB() {
	TESTDB = INITDB.Begin()
}

func teardownDB() {
	TESTDB = TESTDB.Rollback()
}

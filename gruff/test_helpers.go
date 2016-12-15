package gruff

import (
	_ "errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

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
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	values := []interface{}{
		&User{},
		&Debate{},
		&DebateOpinion{},
		&Argument{},
		&ArgumentOpinion{},
		&Link{},
		&Tag{},
		&Context{},
		&Value{},
	}

	for _, value := range values {
		db.DropTable(value)
	}

	if err := db.AutoMigrate(values...).Error; err != nil {
		panic(fmt.Sprintf("No error should happen when create table, but got %+v", err))
	}
}

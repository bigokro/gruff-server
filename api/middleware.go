package api

import (
	model "../model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	_ "os"
	"reflect"
	"strings"
)

var RW_DB_POOL *gorm.DB

type Context struct {
	Database   *gorm.DB
	Payload    map[string]interface{}
	Request    map[string]interface{}
	Type       reflect.Type
	ParentType reflect.Type
	Test       bool
}

func NewContext(test bool, db *gorm.DB) *Context {
	return &Context{
		Test:     test,
		Database: db,
		Payload:  make(map[string]interface{}),
	}
}

func (ctx *Context) DialDatabase(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !ctx.Test {
			db := RW_DB_POOL.Begin()
			defer db.Commit()

			ctx.Database = db
			ctx.Payload = make(map[string]interface{})
		}

		return next(c)
	}
}

func (ctx *Context) DetermineType(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		parts := PathParts(c.Path())
		var pathType string
		for i := 0; i < len(parts); i++ {
			pathType = parts[i]
			t := StringToType(pathType)
			if t != nil {
				if ctx.Type != nil {
					ctx.ParentType = ctx.Type
				}

				ctx.Type = t
			}
		}

		return next(c)
	}
}

func PathParts(path string) []string {
	parts := strings.Split(strings.Trim(path, " /"), "/")
	return parts
}

func StringToType(typeName string) (t reflect.Type) {
	switch typeName {
	case "users":
		var m model.User
		t = reflect.TypeOf(m)
	case "debates":
		var m model.Debate
		t = reflect.TypeOf(m)
	case "debate_opinions":
		var m model.DebateOpinion
		t = reflect.TypeOf(m)
	case "arguments":
		var m model.Argument
		t = reflect.TypeOf(m)
	case "argument_opinions":
		var m model.ArgumentOpinion
		t = reflect.TypeOf(m)
	case "contexts":
		var m model.Context
		t = reflect.TypeOf(m)
	case "references":
		var m model.Reference
		t = reflect.TypeOf(m)
	case "tags":
		var m model.Tag
		t = reflect.TypeOf(m)
	case "values":
		var m model.Value
		t = reflect.TypeOf(m)
	}
	return
}

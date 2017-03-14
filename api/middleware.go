package api

import (
	_ "os"
	"reflect"
	"strings"

	"github.com/bigokro/gruff-server/gruff"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
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

func DBMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("Database", db)
			next(c)
			return nil
		}
	}
}

func (ctx *Context) DialDatabase(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !ctx.Test {
			ctx.Database = c.Get("Database").(*gorm.DB)
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

func (ctx *Context) AssociationFieldNameFromPath(c echo.Context) string {
	path := c.Path()
	parts := strings.Split(path, "/")
	associationPath := ""
	for _, part := range parts {
		if StringToType(part) == ctx.Type {
			associationPath = part
		}
	}
	associationName := SnakeToCamel(associationPath)
	return associationName
}

func PathParts(path string) []string {
	parts := strings.Split(strings.Trim(path, " /"), "/")
	return parts
}

func StringToType(typeName string) (t reflect.Type) {
	switch typeName {
	case "users":
		var m gruff.User
		t = reflect.TypeOf(m)
	case "claims":
		var m gruff.Claim
		t = reflect.TypeOf(m)
	case "claim_opinions":
		var m gruff.ClaimOpinion
		t = reflect.TypeOf(m)
	case "arguments":
		var m gruff.Argument
		t = reflect.TypeOf(m)
	case "argument_opinions":
		var m gruff.ArgumentOpinion
		t = reflect.TypeOf(m)
	case "contexts":
		var m gruff.Context
		t = reflect.TypeOf(m)
	case "links":
		var m gruff.Link
		t = reflect.TypeOf(m)
	case "tags":
		var m gruff.Tag
		t = reflect.TypeOf(m)
	case "values":
		var m gruff.Value
		t = reflect.TypeOf(m)
	}
	return
}

func (ctx *Context) ServerContext() gruff.ServerContext {
	return gruff.ServerContext{
		Database: ctx.Database,
		Test:     false,
	}
}

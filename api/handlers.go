package api

import (
	model "../model"
	_ "fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func (ctx *Context) List(c echo.Context) error {
	db := ctx.Database
	db = BasicJoins(ctx, c, db)
	db = BasicPaging(ctx, c, db)

	items := reflect.New(reflect.SliceOf(ctx.Type)).Interface()
	err := db.Find(items).Error
	if err != nil {
		return model.NewServerError(err.Error())
	}

	items = itemsOrEmptySlice(ctx.Type, items)

	if ctx.Payload["ct"] != nil {
		ctx.Payload["results"] = items
		return c.JSON(http.StatusOK, ctx.Payload)
	} else {
		return c.JSON(http.StatusOK, items)
	}
}

func (ctx *Context) Create(c echo.Context) error {
	item := reflect.New(ctx.Type).Interface()
	if err := c.Bind(item); err != nil {
		return err
	}

	valerr := BasicValidationForCreate(ctx, c, item)
	if valerr != nil {
		return valerr
	}

	dberr := ctx.Database.Create(item).Error
	if dberr != nil {
		return dberr
	}

	return c.JSON(http.StatusCreated, item)
}

func (ctx *Context) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	item := reflect.New(ctx.Type).Interface()

	db := ctx.Database
	db = BasicJoins(ctx, c, db)
	db = BasicFetch(ctx, c, db, id)

	err = db.First(item, id).Error
	if err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	return c.JSON(http.StatusOK, item)
}

func (ctx *Context) GetParent(c echo.Context) error {
	parent := false
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		id, err = strconv.Atoi(c.Param("ownId"))
		parent = true
		if err != nil {
			c.String(http.StatusNotFound, "NotFound")
			return err
		}
	}

	db := ctx.Database
	db = BasicJoins(ctx, c, db)
	item := reflect.New(ctx.Type).Interface()

	if parent {
		item = reflect.New(ctx.ParentType).Interface()
		db = BasicFetch(ctx, c, db, id)
	}

	err = db.First(item, id).Error
	if err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	return c.JSON(http.StatusOK, item)
}

func (ctx *Context) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	item := reflect.New(ctx.Type).Interface()
	err = ctx.Database.First(item, id).Error
	if err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	err = BasicValidationForUpdate(ctx, c, item)
	if err != nil {
		return err
	}

	if err := c.Bind(item); err != nil {
		return err
	}

	ctx.Database.Set("gorm:save_associations", false).Save(item)
	if ctx.Database.Error != nil {
		return ctx.Database.Error
	}

	return c.JSON(http.StatusAccepted, item)
}

func (ctx *Context) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	item := reflect.New(ctx.Type).Interface()
	err = ctx.Database.First(item, id).Error
	if err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	err = ctx.Database.Delete(item).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, item)
}

func (ctx *Context) Destroy(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	item := reflect.New(ctx.Type).Interface()
	err = ctx.Database.First(item, id).Error
	if err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	err = ctx.Database.Unscoped().Delete(item).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, item)
}

func BasicJoins(ctx *Context, c echo.Context, db *gorm.DB) *gorm.DB {
	db = joinsFor(db, ctx)
	return db
}

func BasicFetch(ctx *Context, c echo.Context, db *gorm.DB, uid int) *gorm.DB {
	if uid > 0 {
		id := uint(uid)
		path := c.Path()
		db = fetchFor(db, path, id)
		return db
	}
	return db
}

func fetchFor(db *gorm.DB, path string, userId uint) *gorm.DB {
	parts := strings.Split(path, "/")
	for i := len(parts) - 1; i >= 0; i-- {
		part := parts[i]
		switch part {
		case "users":
		}
	}
	return db
}

func joinsFor(db *gorm.DB, ctx *Context) *gorm.DB {
	t := ctx.Type
	elemT := t
	if elemT.Kind() == reflect.Ptr {
		elemT = elemT.Elem()
	}
	for i := 0; i < elemT.NumField(); i++ {
		f := elemT.Field(i)
		tag := elemT.Field(i).Tag
		fetch := tag.Get("fetch")
		if fetch == "eager" {
			db = db.Preload(f.Name)
		}
	}
	return db
}

func BasicPaging(ctx *Context, c echo.Context, db *gorm.DB, opts ...bool) *gorm.DB {
	queryTC := true
	if len(opts) > 0 {
		queryTC = opts[0]
	}

	st := c.QueryParam("start")
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	if limit > 0 && queryTC {
		QueryTotalCount(ctx, c)
	}

	if st != "" {
		startIdx, _ := strconv.Atoi(st)
		if startIdx > 0 {
			db = db.Offset(startIdx)
		}
	}

	if limit > 0 {
		db = limitQueryByConfig(ctx, db, "", limit)
	}

	return db
}

func QueryTotalCount(ctx *Context, c echo.Context) {
	item := reflect.New(ctx.Type).Interface()
	var n int

	ctx.Database.Model(item).
		Select("COUNT(*)").
		Row().
		Scan(&n)

	ctx.Payload["ct"] = n
}

func limitQueryByConfig(ctx *Context, db *gorm.DB, key string, requestLimit int) *gorm.DB {
	dbLimit := requestLimit
	limitStr := os.Getenv(key)
	limit, err := strconv.Atoi(limitStr)
	if err == nil {
		if dbLimit <= 0 || (limit > 0 && limit < dbLimit) {
			dbLimit = limit
		}
	}
	if dbLimit > 0 {
		db = db.Limit(dbLimit)
	}
	return db
}

func itemsOrEmptySlice(t reflect.Type, items interface{}) interface{} {
	if reflect.ValueOf(items).IsNil() {
		items = reflect.MakeSlice(reflect.SliceOf(t), 0, 0)
	}
	return items
}

func BasicValidationForCreate(ctx *Context, c echo.Context, item interface{}) model.GruffError {
	if model.IsValidator(ctx.Type) {
		validator := item.(model.Validator)
		return validator.ValidateForCreate()
	} else {
		return nil
	}
}

func BasicValidationForUpdate(ctx *Context, c echo.Context, item interface{}) error {
	if model.IsValidator(ctx.Type) {
		return model.ValidateStructFields(item)
	} else {
		return nil
	}
}

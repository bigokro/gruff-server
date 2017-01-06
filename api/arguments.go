package api

import (
	"github.com/bigokro/gruff-server/gruff"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func Arguments(c echo.Context) error {
	return c.String(http.StatusOK, "Arguments")
}

func (ctx *Context) MoveArgument(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	newId, err := uuid.Parse(c.Param("newId"))
	if err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	t, err := strconv.Atoi(c.Param("type"))
	if err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	db := ctx.Database

	arg := gruff.Argument{}
	err = db.Where("id = ?", id).First(&arg).Error
	if err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	switch t {
	case gruff.ARGUMENT_TYPE_PRO_TRUTH, gruff.ARGUMENT_TYPE_CON_TRUTH:
		newDebate := gruff.Debate{}
		err := db.Where("id = ?", newId).First(&newDebate).Error
		if err != nil {
			c.String(http.StatusNotFound, "NotFound")
			return err
		}

		newIdN := gruff.NullableUUID(newId)
		arg.ParentID = &newIdN
		arg.Parent = &newDebate
		arg.ArgumentID = nil

	case gruff.ARGUMENT_TYPE_PRO_RELEVANCE, gruff.ARGUMENT_TYPE_CON_RELEVANCE, gruff.ARGUMENT_TYPE_PRO_IMPACT, gruff.ARGUMENT_TYPE_CON_IMPACT:
		newArg := gruff.Argument{}
		err := db.Where("id = ?", newId).First(&newArg).Error
		if err != nil {
			c.String(http.StatusNotFound, "NotFound")
			return err
		}

		newIdN := gruff.NullableUUID(newId)
		arg.ArgumentID = &newIdN
		arg.Argument = &newArg
		arg.ParentID = nil

	default:
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	db.Set("gorm:save_associations", false).Save(arg)
	if db.Error != nil {
		return db.Error
	}

	ctx.Payload["results"] = arg
	return c.JSON(http.StatusOK, ctx.Payload)
}

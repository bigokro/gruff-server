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

func (ctx *Context) GetArgument(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	argument := gruff.Argument{}

	db := ctx.Database
	db = db.Preload("Debate")
	db = db.Preload("Debate.Links")
	db = db.Preload("Debate.Contexts")
	db = db.Preload("Debate.Values")
	db = db.Preload("Debate.Tags")
	db = db.Preload("Parent")
	db = db.Preload("Parent.Links")
	db = db.Preload("Parent.Contexts")
	db = db.Preload("Parent.Values")
	db = db.Preload("Parent.Tags")
	db = db.Preload("Argument")
	db = db.Preload("Argument.Debate")
	err = db.Where("id = ?", id).First(&argument).Error
	if err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	proRel := []gruff.Argument{}
	db = ctx.Database
	db = db.Where("type = ?", gruff.ARGUMENT_TYPE_PRO_RELEVANCE)
	err = db.Where("argument_id = ?", id).Find(&proRel).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "ServerError")
		return err
	}
	argument.ProRelevance = proRel

	conRel := []gruff.Argument{}
	db = ctx.Database
	db = db.Where("type = ?", gruff.ARGUMENT_TYPE_CON_RELEVANCE)
	err = db.Where("argument_id = ?", id).Find(&conRel).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "ServerError")
		return err
	}
	argument.ConRelevance = conRel

	proImpact := []gruff.Argument{}
	db = ctx.Database
	db = db.Where("type = ?", gruff.ARGUMENT_TYPE_PRO_IMPACT)
	err = db.Where("argument_id = ?", id).Find(&proImpact).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "ServerError")
		return err
	}
	argument.ProImpact = proImpact

	conImpact := []gruff.Argument{}
	db = ctx.Database
	db = db.Where("type = ?", gruff.ARGUMENT_TYPE_CON_IMPACT)
	err = db.Where("argument_id = ?", id).Find(&conImpact).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "ServerError")
		return err
	}
	argument.ConImpact = conImpact

	return c.JSON(http.StatusOK, argument)
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

		newIdN := gruff.NullableUUID{newId}
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

		newIdN := gruff.NullableUUID{newId}
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

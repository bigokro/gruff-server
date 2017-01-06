package api

import (
	"github.com/bigokro/gruff-server/gruff"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"net/http"
)

func Debates(c echo.Context) error {
	return c.String(http.StatusOK, "Debates")
}

func (ctx *Context) GetDebate(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	debate := gruff.Debate{}

	db := ctx.Database
	db = db.Preload("Links")
	db = db.Preload("Contexts")
	db = db.Preload("Values")
	db = db.Preload("Tags")
	err = db.Where("id = ?", id).First(&debate).Error
	if err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	proArgs := []gruff.Argument{}
	db = ctx.Database
	db = db.Preload("Debate")
	db = db.Where("type = ?", gruff.ARGUMENT_TYPE_PRO_TRUTH)
	err = db.Where("parent_id = ?", id).Find(&proArgs).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "ServerError")
		return err
	}
	debate.ProTruth = proArgs

	conArgs := []gruff.Argument{}
	db = ctx.Database
	db = db.Preload("Debate")
	db = db.Where("type = ?", gruff.ARGUMENT_TYPE_CON_TRUTH)
	err = db.Where("parent_id = ?", id).Find(&conArgs).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "ServerError")
		return err
	}
	debate.ConTruth = conArgs

	return c.JSON(http.StatusOK, debate)
}

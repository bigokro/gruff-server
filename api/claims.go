package api

import (
	"errors"
	"fmt"
	"github.com/bigokro/gruff-server/gruff"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"net/http"
	"reflect"
	"strings"
)

func Claims(c echo.Context) error {
	return c.String(http.StatusOK, "Claims")
}

func (ctx *Context) GetClaim(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	claim := gruff.Claim{}

	db := ctx.Database
	db = db.Preload("Links")
	db = db.Preload("Contexts")
	db = db.Preload("Values")
	db = db.Preload("Tags")
	err = db.Where("id = ?", id).First(&claim).Error
	if err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	proArgs := []gruff.Argument{}
	db = ctx.Database
	db = db.Preload("Claim")
	db = db.Where("type = ?", gruff.ARGUMENT_TYPE_PRO_TRUTH)
	err = db.Where("target_claim_id = ?", id).Find(&proArgs).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "ServerError")
		return err
	}
	claim.ProTruth = proArgs

	conArgs := []gruff.Argument{}
	db = ctx.Database
	db = db.Preload("Claim")
	db = db.Where("type = ?", gruff.ARGUMENT_TYPE_CON_TRUTH)
	err = db.Where("target_claim_id = ?", id).Find(&conArgs).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "ServerError")
		return err
	}
	claim.ConTruth = conArgs

	return c.JSON(http.StatusOK, claim)
}

func (ctx *Context) ListTopClaims(c echo.Context) error {
	claims := []gruff.Claim{}

	db := ctx.Database
	db = BasicJoins(ctx, c, db)
	db = db.Where("0 = (SELECT COUNT(*) FROM arguments WHERE claim_id = claims.id)")
	db = BasicPaging(ctx, c, db)

	err := db.Find(&claims).Error
	if err != nil {
		return gruff.NewServerError(err.Error())
	}

	if ctx.Payload["ct"] != nil {
		ctx.Payload["results"] = claims
		return c.JSON(http.StatusOK, ctx.Payload)
	} else {
		return c.JSON(http.StatusOK, claims)
	}
}

func (ctx *Context) SetScore(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	user, err := CurrentUser(c, ctx.Database)
	if err != nil {
		c.String(http.StatusUnauthorized, "NotAuthorized")
		return err
	}

	paths := strings.Split(c.Path(), "/")
	scoreType := paths[len(paths)-1]

	var claim bool
	var target, item interface{}

	switch scoreType {
	case "truth":
		claim = true
		target = &gruff.Claim{}
		item = &gruff.ClaimOpinion{UserID: user.ID, ClaimID: id}
	case "impact":
		claim = false
		target = &gruff.Argument{}
		item = &gruff.ArgumentOpinion{UserID: user.ID, ArgumentID: id}
	case "relevance":
		claim = false
		target = &gruff.Argument{}
		item = &gruff.ArgumentOpinion{UserID: user.ID, ArgumentID: id}
	default:
		c.String(http.StatusNotFound, "NotFound")
		return errors.New("Not found")
	}

	db := ctx.Database
	err = db.Where("id = ?", id).First(target).Error
	if err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	data := map[string]interface{}{}
	if err := c.Bind(&data); err != nil {
		fmt.Println("Error:", err)
		return err
	}
	var score float64
	if val, ok := data["score"]; ok {
		score = val.(float64)
	}

	db = ctx.Database
	db = db.Where("user_id = ?", user.ID)
	if claim {
		db = db.Where("claim_id = ?", id)
	} else {
		db = db.Where("argument_id = ?", id)
	}
	if err := db.First(item).Error; err != nil {
		setScore(item, scoreType, score)
		db = ctx.Database
		err = db.Create(item).Error
		return c.JSON(http.StatusCreated, item)
	} else {
		setScore(item, scoreType, score)
		db = ctx.Database
		err = db.Save(item).Error
		return c.JSON(http.StatusAccepted, item)
	}

}

func setScore(item interface{}, field string, score float64) {
	v := reflect.ValueOf(item)
	if v.Kind() == reflect.Ptr {
		v = reflect.ValueOf(item).Elem()
	}
	f := v.FieldByName(strings.Title(field))
	f.Set(reflect.ValueOf(score))
}

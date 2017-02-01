package api

import (
	"fmt"
	"github.com/bigokro/gruff-server/gruff"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"net/http"
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

func (ctx *Context) SetTruthScore(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	claim := gruff.Claim{}

	db := ctx.Database
	err = db.Where("id = ?", id).First(&claim).Error
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

	user, err := CurrentUser(c, ctx.Database)
	if err != nil {
		c.String(http.StatusUnauthorized, "NotAuthorized")
		return err
	}

	co := gruff.ClaimOpinion{}
	db = ctx.Database
	db = db.Where("user_id = ?", user.ID)
	db = db.Where("claim_id = ?", id)
	if err := db.First(&co).Error; err != nil {
		co.Truth = score
		co.UserID = user.ID
		co.ClaimID = claim.ID
		db = ctx.Database
		err = db.Create(&co).Error
		return c.JSON(http.StatusCreated, co)
	} else {
		co.Truth = score
		db = ctx.Database
		err = db.Save(&co).Error
		return c.JSON(http.StatusAccepted, co)
	}

}

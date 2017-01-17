package api

import (
	"fmt"
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
	db = db.Preload("Claim")
	db = db.Preload("Claim.Links")
	db = db.Preload("Claim.Contexts")
	db = db.Preload("Claim.Values")
	db = db.Preload("Claim.Tags")
	db = db.Preload("TargetClaim")
	db = db.Preload("TargetClaim.Links")
	db = db.Preload("TargetClaim.Contexts")
	db = db.Preload("TargetClaim.Values")
	db = db.Preload("TargetClaim.Tags")
	db = db.Preload("TargetArgument")
	db = db.Preload("TargetArgument.Claim")
	err = db.Where("id = ?", id).First(&argument).Error
	if err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	proRel := []gruff.Argument{}
	db = ctx.Database
	db = db.Preload("Claim")
	db = db.Where("type = ?", gruff.ARGUMENT_TYPE_PRO_RELEVANCE)
	err = db.Where("target_argument_id = ?", id).Find(&proRel).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "ServerError")
		return err
	}
	argument.ProRelevance = proRel

	fmt.Println("Pro Relevance:", proRel)

	conRel := []gruff.Argument{}
	db = ctx.Database
	db = db.Preload("Claim")
	db = db.Where("type = ?", gruff.ARGUMENT_TYPE_CON_RELEVANCE)
	err = db.Where("target_argument_id = ?", id).Find(&conRel).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "ServerError")
		return err
	}
	argument.ConRelevance = conRel

	fmt.Println("Con Relevance:", conRel)

	proImpact := []gruff.Argument{}
	db = ctx.Database
	db = db.Preload("Claim")
	db = db.Where("type = ?", gruff.ARGUMENT_TYPE_PRO_IMPACT)
	err = db.Where("target_argument_id = ?", id).Find(&proImpact).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "ServerError")
		return err
	}
	argument.ProImpact = proImpact

	conImpact := []gruff.Argument{}
	db = ctx.Database
	db = db.Preload("Claim")
	db = db.Where("type = ?", gruff.ARGUMENT_TYPE_CON_IMPACT)
	err = db.Where("target_argument_id = ?", id).Find(&conImpact).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "ServerError")
		return err
	}
	argument.ConImpact = conImpact

	return c.JSON(http.StatusOK, argument)
}

func (ctx *Context) CreateArgument(c echo.Context) error {
	arg := gruff.Argument{Claim: &gruff.Claim{}}
	if err := c.Bind(&arg); err != nil {
		return err
	}

	if arg.ClaimID == uuid.Nil {
		// First create a new Claim for this argument
		// TODO: grab a title from a sub debate sent with the post data
		claim := gruff.Claim{Title: arg.Title, Description: arg.Description}
		if arg.Claim.Title != "" {
			claim.Title = arg.Claim.Title
		}
		if arg.Claim.Description != "" {
			claim.Description = arg.Claim.Description
		}
		valerr := BasicValidationForCreate(ctx, c, claim)
		if valerr != nil {
			return valerr
		}
		err := ctx.Database.Create(&claim).Error
		if err != nil {
			c.String(http.StatusInternalServerError, "ServerError")
			return err
		}
		arg.ClaimID = claim.ID
		arg.Claim = &claim
	} else {
		arg.Claim = nil
	}

	valerr := BasicValidationForCreate(ctx, c, arg)
	if valerr != nil {
		return valerr
	}

	dberr := ctx.Database.Set("gorm:save_associations", false).Create(&arg).Error
	if dberr != nil {
		return dberr
	}

	return c.JSON(http.StatusCreated, arg)
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
		newClaim := gruff.Claim{}
		err := db.Where("id = ?", newId).First(&newClaim).Error
		if err != nil {
			c.String(http.StatusNotFound, "NotFound")
			return err
		}

		newIdN := gruff.NullableUUID{newId}
		arg.TargetClaimID = &newIdN
		arg.TargetClaim = &newClaim
		arg.TargetArgumentID = nil

	case gruff.ARGUMENT_TYPE_PRO_RELEVANCE, gruff.ARGUMENT_TYPE_CON_RELEVANCE, gruff.ARGUMENT_TYPE_PRO_IMPACT, gruff.ARGUMENT_TYPE_CON_IMPACT:
		newArg := gruff.Argument{}
		err := db.Where("id = ?", newId).First(&newArg).Error
		if err != nil {
			c.String(http.StatusNotFound, "NotFound")
			return err
		}

		newIdN := gruff.NullableUUID{newId}
		arg.TargetArgumentID = &newIdN
		arg.TargetArgument = &newArg
		arg.TargetClaimID = nil

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

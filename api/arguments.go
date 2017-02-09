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
	db = db.Where("target_argument_id = ?", id)
	db = db.Scopes(gruff.OrderByBestArgument)
	err = db.Find(&proRel).Error
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
	db = db.Where("target_argument_id = ?", id)
	db = db.Scopes(gruff.OrderByBestArgument)
	err = db.Find(&conRel).Error
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
	db = db.Where("target_argument_id = ?", id)
	db = db.Scopes(gruff.OrderByBestArgument)
	err = db.Find(&proImpact).Error
	if err != nil {
		c.String(http.StatusInternalServerError, "ServerError")
		return err
	}
	argument.ProImpact = proImpact

	conImpact := []gruff.Argument{}
	db = ctx.Database
	db = db.Preload("Claim")
	db = db.Where("type = ?", gruff.ARGUMENT_TYPE_CON_IMPACT)
	db = db.Where("target_argument_id = ?", id)
	db = db.Scopes(gruff.OrderByBestArgument)
	err = db.Find(&conImpact).Error
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

	arg.CreatedByID = CurrentUserID(c)

	if arg.ClaimID == uuid.Nil {
		ctxIds := arg.Claim.ContextIDs

		// First create a new Claim for this argument
		claim := gruff.Claim{Title: arg.Title, Description: arg.Description}
		claim.CreatedByID = arg.CreatedByID
		if arg.Claim.Title != "" {
			claim.Title = arg.Claim.Title
			if arg.Title == "" {
				arg.Title = arg.Claim.Title
			}
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

		for _, ctxId := range ctxIds {
			ctx.Database.Exec("INSERT INTO claim_contexts (claim_id, context_id) VALUES (?, ?)", claim.ID, ctxId)
		}
	} else {
		arg.Claim = nil
	}

	valerr := BasicValidationForCreate(ctx, c, arg)
	if valerr != nil {
		return valerr
	}

	if dberr := ctx.Database.Set("gorm:save_associations", false).Create(&arg).Error; dberr != nil {
		c.String(http.StatusInternalServerError, "ServerError")
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
	if err := db.Where("id = ?", id).First(&arg).Error; err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return err
	}

	if err := (&arg).MoveTo(ctx.ServerContext(), newId, t); err != nil {
		c.String(http.StatusInternalServerError, "ServerError")
		return err
	}

	ctx.Payload["results"] = arg
	return c.JSON(http.StatusOK, ctx.Payload)
}

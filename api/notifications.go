package api

import (
	"errors"
	"github.com/bigokro/gruff-server/gruff"
	"github.com/labstack/echo"
	"net/http"
)

func (ctx *Context) ListNotifications(c echo.Context) error {
	userId := CurrentUserID(c)

	notifications := []gruff.Notification{}
	db := ctx.Database
	db = db.Where("user_id = ?", userId)
	db = db.Where("viewed = false")
	db = db.Order("created_at DESC")
	if err := db.Find(&notifications).Error; err != nil {
		return gruff.NewServerError(err.Error())
	}

	return c.JSON(http.StatusOK, notifications)
}

func (ctx *Context) MarkNotificationViewed(c echo.Context) error {
	db := ctx.Database

	userId := CurrentUserID(c)
	notificationId := c.Param("id")
	if notificationId == "" {
		c.String(http.StatusNotFound, "NotFound")
		return errors.New("Not found")
	}

	notification := gruff.Notification{}
	if err := db.First(&notification, notificationId).Error; err != nil {
		c.String(http.StatusNotFound, "NotFound")
		return errors.New("Not found")
	}

	if notification.UserID != userId {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return errors.New("This is not your notification")
	}

	notification.Viewed = true
	if err := db.Save(&notification).Error; err != nil {
		c.String(http.StatusInternalServerError, "ServerError")
		return gruff.NewServerError(err.Error())
	}

	return c.JSON(http.StatusOK, notification)
}

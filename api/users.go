package api

import (
	model "../model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type jwtCustomClaims struct {
	Name string `json:"name"`
	Type uint   `json:"type"`
	jwt.StandardClaims
}

type customPassword struct {
	Email       string `json:"email"`
	OldPassword string `json:"oldpassword"`
	NewPassword string `json:"newpassword"`
}

func (ctx *Context) SignUp(c echo.Context) error {
	u := new(model.User)

	if err := c.Bind(u); err != nil {
		return err
	}

	password := u.Password
	u.Password = ""
	u.HashedPassword, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err := ctx.Database.Create(u).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}

func (ctx *Context) SignIn(c echo.Context) error {
	u := model.User{}
	if err := c.Bind(&u); err != nil {
		return err
	}

	if u.Email != "" {
		user := model.User{}
		if err := ctx.Database.Where("email = ?", u.Email).Find(&user).Error; err != nil {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}

		if ok, _ := verifyPassword(user, u.Password); ok {

			claims := &jwtCustomClaims{
				u.Email,
				0,
				jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			t, err := token.SignedString([]byte("secret"))
			if err != nil {
				return c.String(http.StatusUnauthorized, "Unauthorized")
			}

			u := map[string]interface{}{"user": user, "token": t}

			return c.JSON(http.StatusOK, u)

		}
	}

	return c.String(http.StatusUnauthorized, "Unauthorized")
}

func verifyPassword(user model.User, password string) (bool, error) {
	return bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password)) == nil, nil
}

func (ctx *Context) ChangePassword(c echo.Context) error {
	u := new(customPassword)
	if err := c.Bind(&u); err != nil {
		return err
	}

	user := model.User{}
	err := ctx.Database.Where("email = ?", u.Email).Find(&user).Error
	if err != nil {
		return model.NewServerError(err.Error())
	}

	if ok, _ := verifyPassword(user, u.OldPassword); ok {
		user.HashedPassword, _ = bcrypt.GenerateFromPassword([]byte(u.NewPassword), bcrypt.DefaultCost)

		if err := ctx.Database.Save(user).Error; err != nil {
			return err
		}

		return c.JSON(http.StatusOK, user)
	}

	return c.String(http.StatusNotFound, "NotFound")
}

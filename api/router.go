package api

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetUpRouter(test bool, db *gorm.DB) *echo.Echo {
	root := echo.New()

	root.Use(middleware.Logger())
	root.Use(middleware.Recover())
	root.Use(middleware.CORS())
	root.Use(middleware.Gzip())
	root.Use(middleware.Secure())

	ctx := NewContext(test, db)
	root.Use(ctx.DialDatabase)
	root.Use(ctx.DetermineType)

	root.GET("/", Home)

	// Public Api
	public := root.Group("/api")

	public.POST("/auth", ctx.SignIn)
	public.POST("/users", ctx.SignUp)

	// Private Api
	private := root.Group("/api")
	private.Use(middleware.JWT([]byte("secret")))

	private.GET("/users", ctx.List)
	private.GET("/users/:id", ctx.Get)
	private.PUT("/users/:id", ctx.Update)
	private.PUT("/users/password", ctx.ChangePassword)
	public.PUT("/users/changePassword", ctx.ChangePassword)
	private.DELETE("/users/:id", ctx.Delete)

	public.GET("/arguments", ctx.List)
	public.GET("/arguments/:id", ctx.Get)
	private.POST("/arguments", ctx.Create)
	private.PUT("/arguments/:id", ctx.Update)
	private.DELETE("/arguments/:id", ctx.Delete)

	public.GET("/argument_opinions", ctx.List)
	public.GET("/argument_opinions/:id", ctx.Get)
	private.POST("/argument_opinions", ctx.Create)
	private.PUT("/argument_opinions/:id", ctx.Update)
	private.DELETE("/argument_opinions/:id", ctx.Delete)

	public.GET("/contexts", ctx.List)
	public.GET("/contexts/:id", ctx.Get)
	private.POST("/contexts", ctx.Create)
	private.PUT("/contexts/:id", ctx.Update)
	private.DELETE("/contexts/:id", ctx.Delete)

	public.GET("/debates", ctx.List)
	public.GET("/debates/:id", ctx.Get)
	private.POST("/debates", ctx.Create)
	private.PUT("/debates/:id", ctx.Update)
	private.DELETE("/debates/:id", ctx.Delete)

	public.GET("/debate_opinions", ctx.List)
	public.GET("/debate_opinions/:id", ctx.Get)
	private.POST("/debate_opinions", ctx.Create)
	private.PUT("/debate_opinions/:id", ctx.Update)
	private.DELETE("/debate_opinions/:id", ctx.Delete)

	public.GET("/references", ctx.List)
	public.GET("/references/:id", ctx.Get)
	private.POST("/references", ctx.Create)
	private.PUT("/references/:id", ctx.Update)
	private.DELETE("/references/:id", ctx.Delete)

	public.GET("/tags", ctx.List)
	public.GET("/tags/:id", ctx.Get)
	private.POST("/tags", ctx.Create)
	private.PUT("/tags/:id", ctx.Update)
	private.DELETE("/tags/:id", ctx.Delete)

	public.GET("/values", ctx.List)
	public.GET("/values/:id", ctx.Get)
	private.POST("/values", ctx.Create)
	private.PUT("/values/:id", ctx.Update)
	private.DELETE("/values/:id", ctx.Delete)

	return root
}

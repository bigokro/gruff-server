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
	public.GET("/arguments/:id", ctx.GetArgument)
	private.POST("/arguments", ctx.CreateArgument)
	private.PUT("/arguments/:id", ctx.Update)
	private.DELETE("/arguments/:id", ctx.Delete)
	private.PUT("/arguments/:id/move/:newId/type/:type", ctx.MoveArgument)

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

	public.GET("/claims", ctx.List)
	public.GET("/claims/:id", ctx.GetClaim)
	private.POST("/claims", ctx.Create)
	private.PUT("/claims/:id", ctx.Update)
	private.DELETE("/claims/:id", ctx.Delete)

	public.GET("/claim_opinions", ctx.List)
	public.GET("/claim_opinions/:id", ctx.Get)
	private.POST("/claim_opinions", ctx.Create)
	private.PUT("/claim_opinions/:id", ctx.Update)
	private.DELETE("/claim_opinions/:id", ctx.Delete)

	public.GET("/links", ctx.List)
	public.GET("/links/:id", ctx.Get)
	private.POST("/links", ctx.Create)
	private.PUT("/links/:id", ctx.Update)
	private.DELETE("/links/:id", ctx.Delete)

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

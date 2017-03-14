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
	root.Use(DBMiddleware(db))

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

	private.POST("/arguments/:id/impact", ctx.SetScore)
	private.PUT("/arguments/:id/impact", ctx.SetScore)
	private.POST("/arguments/:id/relevance", ctx.SetScore)
	private.PUT("/arguments/:id/relevance", ctx.SetScore)

	public.GET("/contexts", ctx.List)
	public.GET("/contexts/:id", ctx.Get)
	private.POST("/contexts", ctx.Create)
	private.PUT("/contexts/:id", ctx.Update)
	private.DELETE("/contexts/:id", ctx.Delete)

	private.POST("/claims/:parentId/contexts/:id", ctx.AddAssociation)
	private.DELETE("/claims/:parentId/contexts/:id", ctx.RemoveAssociation)

	public.GET("/claims", ctx.List)
	public.GET("/claims/top", ctx.ListTopClaims)
	public.GET("/claims/:id", ctx.GetClaim)
	private.POST("/claims", ctx.Create)
	private.PUT("/claims/:id", ctx.Update)
	private.DELETE("/claims/:id", ctx.Delete)
	private.POST("/claims/:id/truth", ctx.SetScore)
	private.PUT("/claims/:id/truth", ctx.SetScore)

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

	private.GET("/notifications", ctx.ListNotifications)
	private.POST("/notifications/:id", ctx.MarkNotificationViewed)
	private.PUT("/notifications/:id", ctx.MarkNotificationViewed)

	return root
}

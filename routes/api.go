package routes

import (
	"simplelink/controllers"
	"github.com/tbxark/g4vercel"
	"github.com/gin-gonic/gin"
)

func ApiRoute(route *gin.Engine) {
	v1 := route.Group("/api/v1")

	app := v1.Group("/app")
	{
		app.GET("test", controllers.RootHandler)
	}

	// auth := v1.Group("/auth")
	// {
	// 	auth.POST("/login", controllers.Login)
	// 	auth.POST("/register", controllers.SignUp)
	// }

	// app := v1.Group("/app")
	// app.Use(middleware.Auth)
	// {
	// 	// USER ROUTES
	// 	userRoutes := app.Group("/user")
	// 	{
	// 		userRoutes.GET("/", controllers.GetUser)
	// 		userRoutes.PUT("/", controllers.UpdateUser)
	// 	}
	// }
}

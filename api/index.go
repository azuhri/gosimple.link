package routes

import (
	"fmt"
	"net/http"
	"simplelink/controllers"

	. "github.com/tbxark/g4vercel"
)

func API(w http.ResponseWriter, r *http.Request) {
	server := New()
	server.Use(Recovery(func(err interface{}, c *Context) {
		if httpError, ok := err.(HttpError); ok {
			c.JSON(httpError.Status, H{
				"message": httpError.Error(),
			})
		} else {
			message := fmt.Sprintf("%s", err)
			c.JSON(500, H{
				"message": message,
			})
		}
	}))
	server.GET("/", controllers.RootHandler)
	// server.GET("/hello", func(context *Context) {
	// 	name := context.Query("name")
	// 	if name == "" {
	// 		context.JSON(400, H{
	// 			"message": "name not found",
	// 		})
	// 	} else {
	// 		context.JSON(200, H{
	// 			"data": fmt.Sprintf("Hello %s!", name),
	// 		})
	// 	}
	// })
	// server.GET("/user/:id", func(context *Context) {
	// 	context.JSON(400, H{
	// 		"data": H{
	// 			"id": context.Param("id"),
	// 		},
	// 	})
	// })
	// server.GET("/long/long/long/path/*test", func(context *Context) {
	// 	context.JSON(200, H{
	// 		"data": H{
	// 			"url": context.Path,
	// 		},
	// 	})
	// })
	// server.Handle(w, r)
}

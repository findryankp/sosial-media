package routes

import (
	"sosialMedia/controllers"
	"sosialMedia/middlewares"

	"github.com/gin-gonic/gin"
)

func userRoutes(c *gin.Engine) {
	r := c.Group("/users").Use(middlewares.Auth())
	{
		r.GET("/", controllers.GetUsers)
		r.GET("/:id", controllers.GetUser)
		r.POST("/", controllers.RegisterUser)
		r.PUT("/:id", controllers.UpdateUser)
		r.DELETE("/:id", controllers.DeleteUser)
	}
}

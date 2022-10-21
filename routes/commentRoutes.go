package routes

import (
	"sosialMedia/controllers"
	"sosialMedia/middlewares"

	"github.com/gin-gonic/gin"
)

func commentRoutes(c *gin.Engine) {
	r := c.Group("/comments").Use(middlewares.Auth())
	{
		r.GET("/", controllers.GetComments)
		r.GET("/:id", controllers.GetComment)
		r.POST("/", controllers.CreateComment)
		r.PUT("/:id", controllers.UpdateComment)
		r.DELETE("/:id", controllers.DeleteComment)
	}
}

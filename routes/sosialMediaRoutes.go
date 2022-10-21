package routes

import (
	"sosialMedia/controllers"
	"sosialMedia/middlewares"

	"github.com/gin-gonic/gin"
)

func sosialMediaRoutes(c *gin.Engine) {
	r := c.Group("/sosial-medias").Use(middlewares.Auth())
	{
		r.GET("/", controllers.GetSosialMedias)
		r.GET("/:id", controllers.GetSosialMedia)
		r.POST("/", controllers.CreateSosialMedia)
		r.PUT("/:id", controllers.UpdateSosialMedia)
		r.DELETE("/:id", controllers.DeleteSosialMedia)
	}
}

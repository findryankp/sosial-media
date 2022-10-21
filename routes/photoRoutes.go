package routes

import (
	"sosialMedia/controllers"
	"sosialMedia/middlewares"

	"github.com/gin-gonic/gin"
)

func photoRoutes(c *gin.Engine) {
	r := c.Group("/photos").Use(middlewares.Auth())
	{
		r.GET("/", controllers.GetPhotos)
		r.GET("/:id", controllers.GetPhoto)
		r.GET("/:id/user", controllers.GetPhotoByIdUser)
		r.POST("/", controllers.CreatePhoto)
		r.PUT("/:id", controllers.UpdatePhoto)
		r.DELETE("/:id", controllers.DeletePhoto)
	}
}

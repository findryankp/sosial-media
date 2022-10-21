package routes

import (
	"sosialMedia/controllers"
	"sosialMedia/middlewares"

	"github.com/gin-gonic/gin"
)

func authRoutes(r *gin.Engine) {
	r.POST("/login", controllers.LoginUser)
	r.POST("/register", controllers.RegisterUser)
	r.GET("/auth-user", controllers.ClaimToken).Use(middlewares.Auth())
}

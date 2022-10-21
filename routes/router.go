package routes

import "github.com/gin-gonic/gin"

func InitRoutes() *gin.Engine {
	r := gin.Default()
	authRoutes(r)
	userRoutes(r)
	commentRoutes(r)
	photoRoutes(r)
	sosialMediaRoutes(r)
	return r
}

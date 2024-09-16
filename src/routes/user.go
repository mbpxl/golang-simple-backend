package routes

import (
	"main/src/controllers"
	"main/src/middlewares"

	"github.com/gin-gonic/gin"
)

func startupsGroupRouter(baseRouter *gin.RouterGroup) {
	startups := baseRouter.Group("/auth")

	startups.GET("/profile",
		middlewares.AuthMiddleware(), controllers.GetProfile)
	startups.POST("/login", controllers.Login)
	startups.POST("/register", controllers.Register)
}

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	versionRouter := r.Group("/api")
	startupsGroupRouter(versionRouter)

	return r
}

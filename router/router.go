package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/AntonYaskevich/lu-server/middlewares"
)

func CreateRouterEngine() *gin.Engine {
	engine := gin.Default();

	api := engine.Group("/api/v1");
	{
		api.POST("/login")
		api.POST("/users")
	}

	auth := api.Group("/")
	auth.Use(middlewares.Auth("key"))

	BindUserRoutes(auth)

	return engine
}
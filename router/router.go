package router

import (
	"github.com/gin-gonic/gin"
	"github.com/AntonYaskevich/lu-server/handlers/users"
	"github.com/AntonYaskevich/lu-server/middlewares"
)

func CreateRouterEngine() *gin.Engine {
	engine := gin.Default();

	api := engine.Group("/api/v1");
	{
		api.POST("/login", users.Login)
		api.POST("/users", users.Create)
	}

	auth := api.Group("/")
	auth.Use(middlewares.Auth("key"))

	BindUserRoutes(&auth)

	return engine
}
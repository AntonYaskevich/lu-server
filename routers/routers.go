package routers

import (
	"github.com/gin-gonic/gin"
)

func CreateEngine() *gin.Engine {
	engine := gin.Default();

	api := engine.Group("/api/v1");
	{
		api.POST("/login")
		api.POST("/users")
	}

	auth := api.Group("/")

	users := auth.Group("/users")
	{
		users.GET("/")
		users.GET("/:id")
		users.PUT("/:id")
		users.DELETE("/:id")
	}
	return engine
}
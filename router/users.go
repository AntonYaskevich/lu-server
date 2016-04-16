package router

import (
	"github.com/gin-gonic/gin"
	"github.com/AntonYaskevich/lu-server/handlers/users"
)

func BindUserRoutes(parent *gin.RouterGroup) {
	router := parent.Group("/users")
	{
		router.GET("/", users.GetAll)
		router.GET("/:id", users.Get)
		router.PUT("/:id", users.Update)
		router.DELETE("/:id", users.Delete)
	}
}

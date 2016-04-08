package routers

import "github.com/gin-gonic/gin"

func BindUserRoutes(parent *gin.RouterGroup) {
	users := parent.Group("/users")
	{
		users.GET("/")
		users.GET("/:id")
		users.PUT("/:id")
		users.DELETE("/:id")
	}
}

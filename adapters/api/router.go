package api

import "github.com/gin-gonic/gin"

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		users := main.Group("user")
		{
			users.GET("/", nil)
			users.GET("/id", nil)
			users.POST("/", nil)
			users.PUT("/:id", nil)
			users.DELETE("/:id", nil)
		}

		Login := main.Group("login")
		{
			Login.POST("/", nil)
		}
	}

	return router
}

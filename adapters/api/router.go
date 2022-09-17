package api

import (
	"github.com/EduardoNevesRamos/frelancer.git/adapters/api/dependency"
	"github.com/EduardoNevesRamos/frelancer.git/adapters/api/middleware"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")

	{
		userControllerWithDependencies := dependency.UserDependency()
		users := main.Group("user", middleware.Auth())
		{
			users.GET("/", userControllerWithDependencies.GetUsers)
			users.GET("/:id", userControllerWithDependencies.GetUserById)
			users.POST("/", userControllerWithDependencies.CreateUser)
			users.PUT("/:id", userControllerWithDependencies.UpdateUser)
			users.DELETE("/:id", userControllerWithDependencies.DeleteUser)
		}

		loginControllerWithDependencies := dependency.LoginDependency()
		main.POST("login", loginControllerWithDependencies.Login)

	}

	return router
}

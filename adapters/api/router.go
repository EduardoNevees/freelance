package api

import (
	"github.com/EduardoNevesRamos/frelancer.git/adapters/api/dependency"
	"github.com/EduardoNevesRamos/frelancer.git/adapters/api/middleware"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	userControllerWithDependencies := dependency.UserDependency()
	loginControllerWithDependencies := dependency.LoginDependency()

	main := router.Group("api/v1")
	{
		users := main.Group("user")
		{
			users.GET("/", userControllerWithDependencies.GetUsers, middleware.Auth())
			users.GET("/:id", userControllerWithDependencies.GetUserById, middleware.Auth())
			users.PUT("/:id", userControllerWithDependencies.UpdateUser, middleware.Auth())
			users.DELETE("/:id", userControllerWithDependencies.DeleteUser, middleware.Auth())
		}
		create := main.Group("user/create")
		{
			create.POST("/", userControllerWithDependencies.CreateUser)

		}

		main.POST("login", loginControllerWithDependencies.Login)

	}

	return router
}

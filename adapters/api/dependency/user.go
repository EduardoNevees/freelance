package dependency

import (
	"github.com/EduardoNevesRamos/frelancer.git/adapters/infrastructure/db"
	"github.com/EduardoNevesRamos/frelancer.git/core/user"
)

func UserDependency() user.IUserController {
	userRepository := user.NewUserRepository(db.Get())
	userService := user.NewUserService(userRepository)
	userController := user.NewUserController(userService)

	return userController
}

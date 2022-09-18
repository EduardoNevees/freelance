package dependency

import (
	"github.com/EduardoNevesRamos/frelancer.git/adapters/infrastructure/db"
	"github.com/EduardoNevesRamos/frelancer.git/core/login"
)

func LoginDependency() login.ILoginController {
	loginRepository := login.NewLoginRepository(db.Get())
	loginService := login.NewLoginService(loginRepository)
	loginController := login.NewLoginController(loginService)
	return loginController
}

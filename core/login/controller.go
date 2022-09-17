package login

import (
	"github.com/EduardoNevesRamos/frelancer.git/model"
	"github.com/gin-gonic/gin"
)

type ILoginController interface {
	Login(context *gin.Context)
}

type LoginController struct {
	service ILoginService
}

func NewLoginController(Service ILoginService) ILoginController {
	return &LoginController{Service}
}

func (c *LoginController) Login(context *gin.Context) {
	login := model.Login{}
	err := context.ShouldBindJSON(&login)
	if err != nil {
		context.JSON(400, gin.H{
			"Error:": "Can't bind JSON" + err.Error(),
		})
	}

	token, err := c.service.Login(&login)
	if err != nil {
		context.JSON(200, gin.H{
			"error": err,
		})
	}

	context.JSON(200, gin.H{
		"token": token,
	})
}

package login

import (
	"net/http"

	"github.com/EduardoNevesRamos/frelancer.git/dto"
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
	login := &dto.Login{}

	err := context.ShouldBindJSON(login)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Error:": "Can't bind JSON" + err.Error(),
		})
		return
	}

	token, err := c.service.Login(login)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Error:": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

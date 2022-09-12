package user

import (
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	GetUsers(context *gin.Context)
	GetUserById(context *gin.Context)
	CreateUser(context *gin.Context)
	UpdateUser(context *gin.Context)
}

type UserController struct {
	service IService
}

func NewUserController(Service IService) IUserController {
	return &UserController{Service}
}

func (c *UserController) GetUsers(context *gin.Context) {
	userResponse, err := c.service.GetUsers()
	if err != nil {
		context.JSON(400, gin.H{
			"Error:": "Can't find user: " + err.Error(),
		})
		return
	}
	context.JSON(200, userResponse)
}

func (c *UserController) GetUserById(context *gin.Context) {
	id := context.Param("id")

	userResponse, err := c.service.GetUserById(id)
	if err != nil {
		context.JSON(400, gin.H{
			"Error:": "Can't find a user" + err.Error(),
		})
		return
	}

	context.JSON(200, userResponse)
}

func (c *UserController) CreateUser(context *gin.Context) {

	userResponse, err := c.service.CreateUser()
	if err != nil {
		context.JSON(400, err)
	}

	err = context.ShouldBindJSON(userResponse)
	if err != nil {
		context.JSON(400, gin.H{
			"Error:": "Can't bind JSON" + err.Error(),
		})
	}

	context.JSON(201, userResponse)
}

func (c *UserController) UpdateUser(context *gin.Context) {
	userResponse, err := c.service.UpdateUser()
	if err != nil {
		context.JSON(400, err)
	}

	err = context.ShouldBindJSON(userResponse)
	if err != nil {
		context.JSON(400, gin.H{
			"Error:": "Can't bind JSON" + err.Error(),
		})
	}

	context.JSON(200, userResponse)
}

func (c *UserController) DeleteUser(context *gin.Context) {
	id := context.Param("id")

	err := c.service.DeleteUser(id)
	if err != nil {
		context.JSON(400, err)
		return
	}

	context.JSON(200, nil)
}

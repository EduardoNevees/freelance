package user

import (
	"github.com/EduardoNevesRamos/frelancer.git/common"
	"github.com/EduardoNevesRamos/frelancer.git/dto"
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	GetUsers(context *gin.Context)
	GetUserById(context *gin.Context)
	CreateUser(context *gin.Context)
	UpdateUser(context *gin.Context)
	DeleteUser(context *gin.Context)
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
	newUser := dto.User{}.WithUserRole()

	err := context.ShouldBindJSON(&newUser)
	if err != nil {
		context.JSON(400, gin.H{
			"Error:": "Can't bind JSON" + err.Error(),
		})
		return
	}

	err = c.service.CreateUser(&newUser)
	if err != nil {
		context.JSON(400, err)
	}

	newUser.Password = common.SHA256Enconder(newUser.Password)

	context.JSON(201, nil)
}

func (c *UserController) UpdateUser(context *gin.Context) {
	id := context.Param("id")

	user := &dto.User{}
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(400, gin.H{
			"Error:": "Can't bind JSON" + err.Error(),
		})
		return
	}

	err = c.service.UpdateUser(user, id)
	if err != nil {
		context.JSON(400, err)
	}

	context.JSON(200, user)
}

func (c *UserController) DeleteUser(context *gin.Context) {
	id := context.Param("id")

	err := c.service.DeleteUser(id)
	if err != nil {
		context.JSON(500, err)
		return
	}

	context.JSON(200, nil)
}

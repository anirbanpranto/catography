package controller

import (
	"net/http"

	common "backend/common"
	model "backend/model"
	service "backend/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UserService
}

func (userController *UserController) Init() {
	//get the appropriate service
	userController.service = service.GetUserService()
}

func (userController UserController) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, userController.service.GetAll())
}

func (userController UserController) CreateUser(c *gin.Context) {
	user := model.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic("Could not bind JSON")
	}

	//check if email already exists
	ok := userController.service.CheckEmail(&user)
	if !ok {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Email already exists",
		})
		return
	}
	//hash password
	user.Password = common.HashPassword(user.Password)
	data := userController.service.Create(user)
	if data == nil {
		panic("Could not create user")
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Created user",
	})
}

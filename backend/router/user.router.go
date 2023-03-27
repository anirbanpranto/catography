package router

import (
	controller "backend/controller"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	controller *controller.UserController
}

func (routes UserRouter) Init(routerGroup *gin.RouterGroup) {
	//set prefix
	userRouter := routerGroup.Group("/users/")
	routes.controller.Init()

	//create routes
	userRouter.GET("/", routes.controller.GetUsers)
	userRouter.POST("/", routes.controller.CreateUser)
}

package router

import (
	controller "backend/controller"

	"github.com/gin-gonic/gin"
)

type ImageRouter struct {
	controller *controller.ImageController
}

func (routes ImageRouter) Init(routerGroup *gin.RouterGroup) {
	//set prefix
	imageRouter := routerGroup.Group("/images/")
	routes.controller.Init()

	//create routes
	imageRouter.GET("/", routes.controller.GetAll)
	imageRouter.POST("/", routes.controller.UploadImage)
	imageRouter.GET("/connect", routes.controller.Pinger)
}

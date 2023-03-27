package main

import (
	"github.com/gin-gonic/gin"

	database "backend/database"
	router "backend/router"
)

func main() {
	//create a default gin engine
	mainRouter := gin.Default()
	database.InitDB()

	//create your initial router group
	apiGroup := mainRouter.Group("/v1")
	//initialize routers
	routers := router.GetRouters()
	for _, routes := range routers {
		routes.Init(apiGroup)
	}
	//run
	mainRouter.Run()
}

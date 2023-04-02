package main

import (
	"github.com/gin-gonic/gin"

	router "backend/router"

	"github.com/gin-contrib/cors"
)

func main() {
	//create a default gin engine
	mainRouter := gin.Default()
	//database.InitDB()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*", "http://localhost:5173"}
	// config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}
	// config.AllowAllOrigins = true

	mainRouter.Use(cors.New(config))

	//create your initial router group
	apiGroup := mainRouter.Group("/v1")

	apiGroup.Use(cors.New(config))
	//initialize routers
	routers := router.GetRouters()
	for _, routes := range routers {
		routes.Init(apiGroup)
	}
	//run
	mainRouter.Run()
}

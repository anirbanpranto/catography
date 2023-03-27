package common

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	Init(routerGroup *gin.RouterGroup)
}

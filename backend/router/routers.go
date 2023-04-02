package router

import (
	"backend/common"
	controller "backend/controller"
)

func GetRouters() []common.Router {
	return []common.Router{
		ImageRouter{&controller.ImageController{}},
	}
}

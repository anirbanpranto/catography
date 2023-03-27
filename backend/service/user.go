package service

import (
	database "backend/database"
	model "backend/model"

	"gorm.io/gorm"
)

var userService *UserService = nil

type UserService struct {
	dbObject *gorm.DB
}

func (userService *UserService) Init() {
	//set prefix
	userService.dbObject = database.GetDB()
	userService.dbObject.AutoMigrate(&model.User{})
}

func GetUserService() *UserService {
	if userService != nil {
		return userService
	}
	userService = &UserService{}
	userService.Init()
	return userService
}

func (userService UserService) Create(u model.User) *gorm.DB {
	return userService.dbObject.Create(&u)
}

func (userService UserService) GetAll() []map[string]interface{} {
	res := []map[string]interface{}{}
	userService.dbObject.Model(&model.User{}).Find(&res)
	return res
}

func (userService UserService) CheckEmail(u *model.User) bool {
	res := model.User{}
	data := userService.dbObject.Where("Email = ?", u.Email).Find(&res)
	return data.RowsAffected <= 0
}

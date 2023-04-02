package service

import (
	common "backend/common"
	database "backend/database"
	model "backend/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

var imageService *ImageService = nil

type ImageService struct {
	dbObject     *gorm.DB
	images       *common.Queue
	Channel_pool int
}

func (imageService *ImageService) Init() {
	//set prefix
	imageService.dbObject = database.GetDB()
	imageService.dbObject.AutoMigrate(&model.Image{})
	imageService.images = &common.Queue{}
	imageService.Channel_pool = 0
}

func UTCtime() string {
	return ""
}

func (imageService *ImageService) Connect() {
	imageService.Channel_pool += 1
}

func (imageService *ImageService) Disconnect() {
	imageService.Channel_pool -= 1
}

func (imageService ImageService) UpdateCache(c chan string) {
	for tick := range time.Tick(time.Minute * 1) {
		/* do things I need done every 1 minute */
		fmt.Println(tick, UTCtime())
		//fmt.Println(imageService.images.GetLength())
		for {
			if imageService.images.IsEmpty() {
				break
			} else {
				image := imageService.images.Peek()
				diff := tick.Sub(image.Time).Minutes()
				if diff >= 10 {
					fmt.Println("Evicting image")
					imageService.images.Dequeue()
					if imageService.Channel_pool > 0 {
						c <- "update"
					}
				} else {
					break
				}
			}
		}
	}
}

func GetImagesService() *ImageService {
	if imageService != nil {
		return imageService
	}
	imageService = &ImageService{}
	imageService.Init()
	return imageService
}

func (imageService *ImageService) Create(img model.Image, c chan string) model.Image {
	imageService.images.Enqueue(img)
	if imageService.Channel_pool > 0 {
		c <- "uploaded"
	}
	return img
}

func (imageService ImageService) GetAll() []model.Image {
	//userService.dbObject.Model(&model.User{}).Find(&res)
	return imageService.images.Elements
}

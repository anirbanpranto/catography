package service

import (
	common "backend/common"
	model "backend/model"
	"fmt"
	"time"
)

var imageService *ImageService = nil

type ImageService struct {
	images       *common.Queue
	channel_pool int
}

func (imageService *ImageService) Init() {
	//set prefix
	imageService.images = &common.Queue{}
	imageService.channel_pool = 0
}

func UTCtime() string {
	return ""
}

func (imageService *ImageService) Connect() {
	imageService.channel_pool += 1
}

func (imageService *ImageService) Disconnect() {
	imageService.channel_pool -= 1
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
					if imageService.channel_pool > 0 {
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
	if imageService.channel_pool > 0 {
		c <- "uploaded"
	}
	return img
}

func (imageService ImageService) GetAll() []model.Image {
	//userService.dbObject.Model(&model.User{}).Find(&res)
	return imageService.images.Elements
}

package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"backend/model"
	service "backend/service"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Location struct {
	Lon  float64
	Lat  float64
	Note string
}

var upgrader = websocket.Upgrader{
	// Solve cross-domain problems
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ImageController struct {
	service     *service.ImageService
	aws_session *session.Session
	s3_svc      *s3.S3
	uploader    *s3manager.Uploader
	channel     chan string
}

func (imageController *ImageController) Init() {
	//get the appropriate service
	imageController.service = service.GetImagesService()
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
	})
	imageController.aws_session = sess
	if err != nil {
		println(err)
	}
	//s3 service
	svc := s3.New(sess)
	imageController.s3_svc = svc
	uploader := s3manager.NewUploaderWithClient(imageController.s3_svc)
	imageController.uploader = uploader

	imageController.channel = make(chan string)

	go imageController.service.UpdateCache(imageController.channel)
}

func (imageController ImageController) Pinger(c *gin.Context) {
	//Upgrade get request to webSocket protocol
	wshandler(c.Writer, c.Request, &imageController)
}

func wshandler(w http.ResponseWriter, r *http.Request, ic *ImageController) {
	conn, err := upgrader.Upgrade(w, r, nil)
	ic.service.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// t, msg, err := conn.ReadMessage()
		// if err != nil {
		// 	break
		// }
		// print(t, msg)
		data := <-ic.channel
		conn.WriteMessage(websocket.TextMessage, []byte(data))
		closeHandler := conn.CloseHandler()
		conn.SetCloseHandler(func(code int, text string) error {
			ic.service.Disconnect()
			err := closeHandler(code, text)
			// ... or here.
			return err
		})
	}
}

func (imageController ImageController) GetAll(c *gin.Context) {
	println(imageController.service.GetAll())
	c.JSON(http.StatusOK, imageController.service.GetAll())
}

func (imageController ImageController) UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	lon := c.PostForm("lon")
	lat := c.PostForm("lat")
	fmt.Println(lon, lat)
	lon_d, _ := strconv.ParseFloat(lon, 64)
	lat_d, _ := strconv.ParseFloat(lat, 64)

	// The file cannot be received.
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension
	bucketName := "catography"
	file_obj, err := file.Open()

	if err != nil {
		println(err)
	}

	upParams := &s3manager.UploadInput{
		Bucket:      &bucketName,
		Key:         &newFileName,
		Body:        file_obj,
		ContentType: &extension,
	}

	result, err := imageController.uploader.Upload(upParams)

	if err != nil {
		println(err)
	}

	req, _ := imageController.s3_svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &newFileName,
	})
	urlStr, err := req.Presign(15 * time.Minute)

	img := model.Image{Url: urlStr, Time: time.Now(), Unsigned: result.Location, Lon: lon_d, Lat: lat_d}
	imageController.service.Create(img, imageController.channel)
	// File saved successfully. Return proper result
	c.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
	})

	if err != nil {
		panic("Could not bind JSON")
	}
}

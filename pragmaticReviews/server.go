package main

import (
	"github.com/gin-gonic/gin"
	"github.com/santoshdhakal667/mykanban/pragma/controller"
	"github.com/santoshdhakal667/mykanban/pragma/service"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
	ImageService    service.ImageService       = service.ImageNew()
	ImageController controller.ImageController = controller.ImageNew(ImageService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.Save(ctx))
	})
	server.POST("/images", func(ctx *gin.Context) {
		ctx.JSON(200, ImageController.Save(ctx))
	})

	server.Run(":8080")
}

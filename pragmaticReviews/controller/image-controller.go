package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/santoshdhakal667/mykanban/pragma/entity"
	"github.com/santoshdhakal667/mykanban/pragma/service"
)

type ImageController interface {
	Save(ctx *gin.Context) entity.Image
}

type imageController struct {
	service service.ImageService
}

func ImageNew(service service.ImageService) ImageController {
	return &imageController{
		service: service,
	}
}

func (ic *imageController) Save(ctx *gin.Context) entity.Image {
	var image entity.Image
	ctx.BindJSON(&image)
	ic.service.Save(image)
	return image
}

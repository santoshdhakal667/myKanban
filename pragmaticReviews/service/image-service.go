package service

import "github.com/santoshdhakal667/mykanban/pragma/entity"

type ImageService interface {
	Save(entity.Image) entity.Image
}

type imageService struct {
	images []entity.Image
}

func ImageNew() ImageService {
	return &imageService{}
}

func (is *imageService) Save(image entity.Image) entity.Image {
	is.images = append(is.images, image)
	return image
}

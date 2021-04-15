package services

import "github.com/flickr/entities"

type IFlickrService interface {
	OperateImage(image *entities.FlickrImage) error
	GetImage(image *entities.FlickrImage) (*entities.FlickrImage, error)
	UpdateVote(image *entities.FlickrImage) error
	GetAllImages() ([]entities.FlickrImage, error)
}

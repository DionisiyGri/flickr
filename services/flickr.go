package services

import (
	"github.com/flickr/entities"
	"github.com/flickr/repositories"
)

type FlickrService struct {
	flickrRepo repositories.IFlickrRepository
}

func NewFlickrSvc(flickrRepo repositories.IFlickrRepository) *FlickrService {
	return &FlickrService{
		flickrRepo: flickrRepo,
	}
}

func (fs *FlickrService) OperateImage(image *entities.FlickrImage) error {
	img, err := fs.GetImage(image)
	if err != nil {
		return err
	}
	if img.ID == 0 {
		err := fs.flickrRepo.OperateImage(image)
		if err != nil {
			return err
		}
		return nil
	}
	err = fs.UpdateVote(image)
	if err != nil {
		return err
	}
	return nil
}

func (fs *FlickrService) GetImage(image *entities.FlickrImage) (*entities.FlickrImage, error) {
	img, err := fs.flickrRepo.GetImage(image)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func (fs *FlickrService) UpdateVote(image *entities.FlickrImage) error {
	err := fs.flickrRepo.UpdateVote(image)
	if err != nil {
		return err
	}
	return nil
}

func (fs *FlickrService) GetAllImages() ([]entities.FlickrImage, error) {
	images, err := fs.flickrRepo.GetAllImages()
	if err != nil {
		return nil, err
	}
	return images, nil
}

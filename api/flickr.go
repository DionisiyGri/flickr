package api

import (
	"log"
	"net/http"

	"github.com/flickr/services"

	"github.com/flickr/entities"

	"github.com/labstack/echo/v4"
)

type flickrHandler struct {
	fs services.IFlickrService
}

func (f *flickrHandler) OperateImage(c echo.Context) error {
	image := &entities.FlickrImage{}
	err := c.Bind(image)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, nil)
	}
	err = f.fs.OperateImage(image)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return nil
}

func (f *flickrHandler) GetAllImages(c echo.Context) error {
	images, err := f.fs.GetAllImages()
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, images)
}

package api

import (
	"github.com/flickr/services"

	"github.com/labstack/echo/v4"
)

func InitHandlers(e *echo.Echo, fs *services.FlickrService) {
	fh := &flickrHandler{fs: fs}

	flickrBaseURL := e.Group("/api")
	flickrBaseURL.POST("/images", fh.OperateImage)
	flickrBaseURL.GET("/images", fh.GetAllImages)
}

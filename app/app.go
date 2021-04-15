package app

import (
	"log"

	"github.com/flickr/api"
	"github.com/flickr/db"
	"github.com/flickr/repositories"
	"github.com/flickr/services"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

const (
	defaultAppPort = ":8000"
)

func StartServer() {
	dbConn, err := db.NewSQLiteConnection()
	if err != nil {
		log.Fatalf("Cannot connect to db. Err = %s", err)
	}

	migrator, err := db.RunMigrationSQLite(dbConn, "./db/migration")
	if err != nil {
		log.Fatal(err)
	}

	_, err = migrator.Up()
	if err != nil {
		log.Fatal(err)
	}

	flickrRepo := repositories.NewFlickrRepo(dbConn)
	flickrSvc := services.NewFlickrSvc(flickrRepo)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	api.InitHandlers(e, flickrSvc)
	logrus.Fatal(e.Start(defaultAppPort))
}

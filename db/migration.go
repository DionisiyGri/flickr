package db

import (
	"database/sql"
	"strings"

	migrate "github.com/golang-migrate/migrate/v4"
	_sqlite3 "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/mattes/migrate/source/file"
)

type Migration struct {
	Migrate *migrate.Migrate
}

func (m *Migration) Up() (bool, error) {
	err := m.Migrate.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			return true, nil
		}
		return false, err
	}
	return true, nil
}

func (m *Migration) Down() (bool, error) {
	err := m.Migrate.Down()
	if err != nil {
		return false, err
	}
	return true, err
}

func RunMigrationSQLite(dbConn *sql.DB, migrationsFolderLocation string) (*Migration, error) {
	dataPath := []string{}
	dataPath = append(dataPath, "file://")
	dataPath = append(dataPath, migrationsFolderLocation)

	pathToMigrate := strings.Join(dataPath, "")

	driver, err := _sqlite3.WithInstance(dbConn, &_sqlite3.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(pathToMigrate, "sqlite3", driver)
	if err != nil {
		return nil, err
	}

	return &Migration{Migrate: m}, nil
}

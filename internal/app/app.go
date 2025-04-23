package app

import (
	"log"
	"os"

	"fixable.com/fixable/internal/config"
	"gorm.io/gorm"
)

type App struct {
	Db *gorm.DB
  Logger *log.Logger
}

func NewApp() (*App, error) {
	db, err := config.ConnectDb()
  logger := log.New(os.Stdout, "",log.Ldate|log.Ltime)
	if err != nil {
		return nil, err
	}

	return &App{
		Db: db,
    Logger: logger,
	}, nil

}

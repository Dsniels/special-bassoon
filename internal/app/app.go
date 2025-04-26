package app

import (
	"log"
	"os"

	"fixable.com/fixable/internal/config"
	handler "fixable.com/fixable/internal/handlers"
	"fixable.com/fixable/internal/repositories"
	"gorm.io/gorm"
)

type App struct {
	Db *gorm.DB
  ServicioHandler *handler.ServicioHandler
  CategoriaHandler *handler.CategoriaHandler
  Logger *log.Logger
}

func NewApp() (*App, error) {
	db, err := config.ConnectDb()
  logger := log.New(os.Stdout, "",log.Ldate|log.Ltime)
  servicioRepository := repositories.NewServicioRepository(db)
  categoriaRepository := repositories.NewCategoriaRepository(db)
  servicioHandler := handler.NewServiceHandler(servicioRepository, categoriaRepository)
  categoriaHandler := handler.NewCategoriaHandler(categoriaRepository)
	if err != nil {
		return nil, err
	}

	return &App{
		Db: db,
    Logger: logger,
    ServicioHandler: servicioHandler,
    CategoriaHandler: categoriaHandler,
	}, nil

}

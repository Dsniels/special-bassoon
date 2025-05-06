package app

import (
	"log"
	"os"

	"fixable.com/fixable/internal/config"
	handler "fixable.com/fixable/internal/handlers"
	"fixable.com/fixable/internal/storage"
	"gorm.io/gorm"
)

type App struct {
	Db                *gorm.DB
	ServicioHandler   *handler.ServicioHandler
	CategoriaHandler  *handler.CategoriaHandler
	ComentarioHandler *handler.ComentariosHandler
	Logger            *log.Logger
}

func NewApp() (*App, error) {
	db, err := config.ConnectDb()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	servicioStorage := storage.NewServicioStorage(db)
	categoriaStorage := storage.NewCategoriaStorage(db)
	comentarioStorage := storage.NewComentarioStorage(db)
	comentarioHandler := handler.NewComentariosHandler(comentarioStorage, servicioStorage)
	servicioHandler := handler.NewServiceHandler(servicioStorage, comentarioStorage, categoriaStorage)
	categoriaHandler := handler.NewCategoriaHandler(categoriaStorage, servicioStorage)
	if err != nil {
		return nil, err
	}

	return &App{
		Db:                db,
		Logger:            logger,
		ServicioHandler:   servicioHandler,
		ComentarioHandler: comentarioHandler,
		CategoriaHandler:  categoriaHandler,
	}, nil

}

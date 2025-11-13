package app

import (
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
}

func NewApp() (*App, error) {
	db, err := config.ConnectDb()
	if err != nil {
		return nil, err
	}
	azClient, err := config.NewBlobClient()
	if err != nil {
		return nil, err
	}
	servicioStorage := storage.NewServicioStorage(db)
	categoriaStorage := storage.NewCategoriaStorage(db)
	comentarioStorage := storage.NewComentarioStorage(db)
	blobStorage := storage.NewBlobStorage(azClient)
	comentarioHandler := handler.NewComentariosHandler(comentarioStorage, servicioStorage)
	servicioHandler := handler.NewServiceHandler(servicioStorage, comentarioStorage, blobStorage, categoriaStorage)
	categoriaHandler := handler.NewCategoriaHandler(categoriaStorage, servicioStorage)
	return &App{
		Db:                db,
		ServicioHandler:   servicioHandler,
		ComentarioHandler: comentarioHandler,
		CategoriaHandler:  categoriaHandler,
	}, nil

}

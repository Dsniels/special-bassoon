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
	Db                *gorm.DB
	ServicioHandler   *handler.ServicioHandler
	CategoriaHandler  *handler.CategoriaHandler
	ComentarioHandler *handler.ComentariosHandler
	Logger            *log.Logger
}

func NewApp() (*App, error) {
	db, err := config.ConnectDb()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	servicioRepository := repositories.NewServicioRepository(db)
	categoriaRepository := repositories.NewCategoriaRepository(db)
	comentarioRepository := repositories.NewComentarioRepository(db)
	comentarioHandler := handler.NewComentariosHandler(comentarioRepository)
	servicioHandler := handler.NewServiceHandler(servicioRepository, comentarioRepository, categoriaRepository)
	categoriaHandler := handler.NewCategoriaHandler(categoriaRepository,servicioRepository)
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

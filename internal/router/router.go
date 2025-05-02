package router

import (
	"fixable.com/fixable/internal/app"
	"github.com/go-chi/chi/v5"
)

func InitRoutes(app *app.App) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/servicios", func(r chi.Router) {
		r.Get("/create", app.ServicioHandler.NewServicio)
		r.Post("/create", app.ServicioHandler.CreateServicioHandler)
		r.Get("/all", app.ServicioHandler.GetAllServicios)
		r.Get("/{id}", app.ServicioHandler.GetServicioById)
	})
	router.Route("/categoria", func(r chi.Router) {
		r.Post("/create", app.CategoriaHandler.CreateCategoriaHandler)
		r.Get("/{id}", app.CategoriaHandler.ServiciosByCategoriaHandler)
		r.Get("/all", app.CategoriaHandler.ShowCategoriasHandler)
	})
	router.Route("/comentarios", func(r chi.Router) {
		r.Get("/", app.ComentarioHandler.ShowComentarios)
		r.Post("/create", app.ComentarioHandler.CreateComentarioHandler)
	})

	return router
}

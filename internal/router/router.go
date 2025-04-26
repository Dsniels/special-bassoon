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
    r.Get("/create", app.CategoriaHandler.ShowCreateForm)
    r.Post("/create", app.CategoriaHandler.CreateCategoriaHandler)
		r.Get("/all", app.CategoriaHandler.ShowCategoriasHandler)
    r.Get("/{id}", app.ServicioHandler.GetServicioById)
	})

	return router
}

package router

import (
	"net/http"

	"fixable.com/fixable/internal/app"
	"github.com/go-chi/chi/v5"
)

func InitRoutes(app *app.App) *chi.Mux {
	router := chi.NewRouter()
	files := http.FileServer(http.Dir("internal/images"))
	router.Handle("/images/*", http.StripPrefix("/images/", files))
	router.Get("/", app.ServicioHandler.GetAllServicios)
	router.Get("/promocionate", app.ServicioHandler.PromocionarseHandler)
	router.Route("/servicios", func(r chi.Router) {
		r.Get("/search", app.ServicioHandler.SearchHandler)
		r.Get("/{id}", app.ServicioHandler.GetServicioById)
		r.Post("/create", app.ServicioHandler.CreateServicio)
		r.Get("/create/{id}", app.ServicioHandler.CreateForm)
		r.Post("/delete/{id}", app.ServicioHandler.DeleteServicio)
	})
	router.Route("/admin/servicios", func(r chi.Router) {
		r.Get("/", app.ServicioHandler.GetAdminAllServicios)
		r.Post("/create", app.ServicioHandler.CreateServicio)
		r.Get("/create/{id}", app.ServicioHandler.CreateForm)
		r.Post("/delete/{id}", app.ServicioHandler.DeleteServicio)
	})
	router.Route("/categoria", func(r chi.Router) {
		r.Get("/{id}", app.CategoriaHandler.ServiciosByCategoriaHandler)
	})
	router.Route("/comentarios", func(r chi.Router) {
		r.Get("/{id}", app.ComentarioHandler.ShowComentarios)
		r.Post("/create/{id}", app.ComentarioHandler.CreateComentarioHandler)
	})

	return router
}

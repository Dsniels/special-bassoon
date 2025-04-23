package router

import "github.com/go-chi/chi/v5"

func InitRoutes() *chi.Mux {
  router := chi.NewRouter()
  return router
}

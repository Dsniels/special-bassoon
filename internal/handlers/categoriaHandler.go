package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"fixable.com/fixable/internal/models"
	"fixable.com/fixable/internal/storage"
	"fixable.com/fixable/internal/utils"
)

type CategoriaHandler struct {
	_categoriaStorage storage.ICategoriaStorage
	_serviciosStorage storage.IServicioStorage
}

func NewCategoriaHandler(categoriaStorage storage.ICategoriaStorage, serviciosStorage storage.IServicioStorage) *CategoriaHandler {
	return &CategoriaHandler{
		_categoriaStorage: categoriaStorage,
		_serviciosStorage: serviciosStorage,
	}
}


func (h *CategoriaHandler) ServiciosByCategoriaHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIdFromParams(r)
	if err != nil {
		fmt.Println("%w", err)
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"error": err})
		return
	}
	categoria, err := h._categoriaStorage.GetCategoriaById(id)
	if err != nil {
		fmt.Println("%w", err)
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"error": err})
		return
	}
	servicios, err := h._serviciosStorage.GetServiceByCategoriaId(id)
	if err != nil {
		fmt.Println("%w", err)
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"error": err})
		return
	}
	t, err := template.ParseFiles(
		"internal/templates/base.template",
		"internal/templates/navbar/navbar.template",
		"internal/templates/categorias/servicios-list.template")
	if err != nil {
		fmt.Println("%w", err)
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"error": err})
		return
	}
	data := struct {
		Categoria string
		Servicios []models.Servicio
	}{
		Categoria: categoria.Nombre,
		Servicios: *servicios,
	}
	t.Execute(w, data)
}

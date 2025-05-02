package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"fixable.com/fixable/internal/models"
	"fixable.com/fixable/internal/repositories"
	"fixable.com/fixable/internal/utils"
)

type CategoriaHandler struct {
	_categoriaRepository repositories.ICategoriaRepository
	_serviciosRepository repositories.IServicioRepository
}

func NewCategoriaHandler(categoriaRepository repositories.ICategoriaRepository, serviciosRepository repositories.IServicioRepository) *CategoriaHandler {
	return &CategoriaHandler{
		_categoriaRepository: categoriaRepository,
		_serviciosRepository: serviciosRepository,
	}
}

func (h *CategoriaHandler) ShowCategoriasHandler(w http.ResponseWriter, r *http.Request) {
	_, err := h._categoriaRepository.GetCategorias()
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"error": err})
		return
	}
}

func (h *CategoriaHandler) CreateCategoriaHandler(w http.ResponseWriter, r *http.Request) {
	var categoria models.Categoria
	categoria.Nombre = r.FormValue("nombre")
	err := h._categoriaRepository.Create(categoria)
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"error": err})
		return
	}
	http.Redirect(w, r, "/servicios/all", http.StatusSeeOther)
}

func (h *CategoriaHandler) ShowCreateForm(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"internal/templates/base.template",
		"internal/templates/navbar/navbar.template",
		"internal/templates/categorias/create.template")
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"error": err})
		return
	}
	t.Execute(w, nil)
}

func (h *CategoriaHandler) ServiciosByCategoriaHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIdFromParams(r)
	if err != nil {
		fmt.Println("%w", err)
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"error": err})
		return
	}
	categoria, err := h._categoriaRepository.GetCategoriaById(id)
	if err != nil {
		fmt.Println("%w", err)
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"error": err})
		return
	}
	servicios, err := h._serviciosRepository.GetServiceByCategoriaId(id)
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

package handler

import (
	"html/template"
	"net/http"

	"fixable.com/fixable/internal/models"
	"fixable.com/fixable/internal/repositories"
	"fixable.com/fixable/internal/utils"
)

type CategoriaHandler struct {
	_categoriaRepository repositories.ICategoriaRepository
}

func NewCategoriaHandler(categoriaRepository repositories.ICategoriaRepository) *CategoriaHandler {
	return &CategoriaHandler{
		_categoriaRepository: categoriaRepository,
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
	t, err := template.ParseFiles("internal/templates/base.template", "internal/templates/categorias/create.template")
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"error": err})
		return
	}
	t.Execute(w, nil)

}

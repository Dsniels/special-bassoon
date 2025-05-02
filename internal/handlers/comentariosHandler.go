package handler

import (
	"html/template"
	"net/http"

	"fixable.com/fixable/internal/models"
	"fixable.com/fixable/internal/repositories"
	"fixable.com/fixable/internal/utils"
)

type ComentariosHandler struct {
	_comentariosRepository repositories.IComentarioRepository
}

func NewComentariosHandler(
	comentariosRepository repositories.IComentarioRepository,
) *ComentariosHandler {

	return &ComentariosHandler{ 
		_comentariosRepository: comentariosRepository,
	}
}

func (h *ComentariosHandler) prepareData() (any, error) {
	comentarios, err := h._comentariosRepository.GetAllComentarios()
	if err != nil {
		return nil, err
	}
	return struct {
		Comentarios []models.Comentario
	}{
		Comentarios: *comentarios,
	}, nil

}

func (h *ComentariosHandler) ShowComentarios(w http.ResponseWriter, r *http.Request) {
	data, err := h.prepareData()
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"message": err})
		return
	}
	t, err := template.ParseFiles(
		"internal/templates/base.template",
		"internal/templates/navbar/base.template", 
		"internal/templates/comentarios/comentarios-list.template")
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"message": err})
		return
	}
	t.Execute(w, data)
}

func (h *ComentariosHandler) CreateComentarioHandler(w http.ResponseWriter, r *http.Request) {
	var comentario models.Comentario
	comentario.Comentario = r.FormValue("comentario")
	err := h._comentariosRepository.CreateComentario(&comentario)
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"message": err})
		return
	}
	http.Redirect(w, r, "/servicios/all", http.StatusSeeOther)

}

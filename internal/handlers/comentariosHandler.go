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
	_serviciosRepository   repositories.IServicioRepository
}

func NewComentariosHandler(
	comentariosRepository repositories.IComentarioRepository,
	serviciosRepository repositories.IServicioRepository,
) *ComentariosHandler {

	return &ComentariosHandler{
		_comentariosRepository: comentariosRepository,
		_serviciosRepository:   serviciosRepository,
	}
}

func (h *ComentariosHandler) ShowComentarios(w http.ResponseWriter, r *http.Request) {
	servicioId, err := utils.GetIdFromParams(r)
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"message": err})
		return
	}
	servicio, err := h._serviciosRepository.GetServiceById(servicioId)
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"message": err})
		return
	}
	comentarios, err := h._comentariosRepository.GetComentariosPorServicioId(int(servicioId))
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"message": err})
		return
	}

	data := struct {
		Comentarios []models.Comentario
		Servicio    *models.Servicio
	}{
		Comentarios: *comentarios,
		Servicio:    servicio,
	}

	t, err := template.ParseFiles(
		"internal/templates/base.template",
		"internal/templates/navbar/navbar.template",
		"internal/templates/comentarios/comentarios-list.template")
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"message": err})
		return
	}
	t.Execute(w, data)
}

func (h *ComentariosHandler) CreateComentarioHandler(w http.ResponseWriter, r *http.Request) {
	var comentario models.Comentario
	servicioId, err := utils.GetIdFromParams(r)
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"message": err})
		return
	}
	comentario.Comentario = r.FormValue("comentario")
	comentario.ServicioId = servicioId
	err = h._comentariosRepository.CreateComentario(&comentario)
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"message": err})
		return
	}
	utils.WriteResponse(w, http.StatusOK, utils.Response{
        "message": "Comentario creado exitosamente",
        "comentario": comentario,
    })
	return

}

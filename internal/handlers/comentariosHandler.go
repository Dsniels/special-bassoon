package handler

import (
	"html/template"
	"net/http"
	"strings"

	"fixable.com/fixable/internal/models"
	"fixable.com/fixable/internal/storage"
	"fixable.com/fixable/internal/utils"
)

type ComentariosHandler struct {
	_comentariosStorage storage.IComentarioStorage
	_serviciosStorage   storage.IServicioStorage
}

func (h *ComentariosHandler) ShowComentarios(w http.ResponseWriter, r *http.Request) {
	servicioId, err := utils.GetIdFromParams(r)
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"message": err})
		return
	}
	servicio, err := h._serviciosStorage.GetServiceById(servicioId)
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"message": err})
		return
	}
	comentarios, err := h._comentariosStorage.GetComentariosPorServicioId(int(servicioId))
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"message": err})
		return
	}

	isAdmin := strings.Contains(r.URL.Path, "/admin/")
	data := struct {
		Comentarios []models.Comentario
		Servicio    *models.Servicio
		IsAdmin     bool
	}{
		Comentarios: *comentarios,
		Servicio:    servicio,
		IsAdmin:     isAdmin,
	}

	t, err := template.ParseFS(TemplatesFS,
		"templates/base.templ",
		"templates/navbar/navbar.templ",
		"templates/comentarios/comentarios-list.templ")
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
	err = h._comentariosStorage.CreateComentario(&comentario)
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"message": err})
		return
	}
	utils.WriteResponse(w, http.StatusOK, utils.Response{
		"comentario": comentario,
	})
}

func NewComentariosHandler(
	comentariosStorage storage.IComentarioStorage,
	serviciosStorage storage.IServicioStorage,
) *ComentariosHandler {

	return &ComentariosHandler{
		_comentariosStorage: comentariosStorage,
		_serviciosStorage:   serviciosStorage,
	}
}

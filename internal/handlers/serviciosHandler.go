package handler

import (
	"html/template"
	"net/http"

	"fixable.com/fixable/internal/models"
	"fixable.com/fixable/internal/storage"
	"fixable.com/fixable/internal/utils"
)

type ServicioHandler struct {
	_servicioStorage   storage.IServicioStorage
	_categoriaStorage  storage.ICategoriaStorage
	_comentarioStorage storage.IComentarioStorage
}

func NewServiceHandler(
	serviciosRepositorio storage.IServicioStorage,
	comentarioStorage storage.IComentarioStorage,
	categoriaStorage storage.ICategoriaStorage,
) *ServicioHandler {
	return &ServicioHandler{
		_categoriaStorage:  categoriaStorage,
		_comentarioStorage: comentarioStorage,
		_servicioStorage:   serviciosRepositorio,
	}
}


func (h *ServicioHandler) PromocionarseHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"internal/templates/base.template",
		"internal/templates/navbar/navbar.template",
		"internal/templates/promocionarse/promocionarse.template",
	)
  if err!= nil{
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"error": err})
		return

  }
  t.Execute(w,nil)
}

func (h *ServicioHandler) SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		utils.WriteResponse(w, http.StatusOK, utils.Response{"data": []string{}})
		return
	}

	servicios, err := h._servicioStorage.GetByQuery(query)
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"message": "algo salio mal"})
		return
	}

	utils.WriteResponse(w, http.StatusOK, utils.Response{"data": servicios})
}

func (h *ServicioHandler) GetAllServicios(w http.ResponseWriter, r *http.Request) {
	services, _ := h._servicioStorage.GetServices()
	comentarios, _ := h._comentarioStorage.GetAllComentarios()
	categorias, _ := h._categoriaStorage.GetCategorias()
	t, err := template.ParseFiles(
		"internal/templates/base.template",
		"internal/templates/navbar/navbar.template",
		"internal/templates/servicios/home.template",
		"internal/templates/categorias/list.template",
	)

	if err != nil {
		panic(err)
	}

	data := struct {
		Servicios   []models.Servicio
		Categorias  []models.Categoria
		Comentarios []models.Comentario
	}{
		Comentarios: *comentarios,
		Servicios:   *services,
		Categorias:  *categorias,
	}

	err = t.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func (h *ServicioHandler) GetServicioById(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIdFromParams(r)
	if err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, utils.Response{"error": err})
		return
	}
	servicio, err := h._servicioStorage.GetServiceById(id)
	if err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, utils.Response{"error": err})
		return
	}
	t, err := template.ParseFiles(
		"internal/templates/base.template",
		"internal/templates/navbar/navbar.template",
		"internal/templates/servicios/servicio.template",
	)
	if err != nil {
		panic(err)
	}
	data := struct {
		Servicio models.Servicio
	}{
		Servicio: *servicio,
	}
	err = t.Execute(w, data)
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"message": err})
	}
}

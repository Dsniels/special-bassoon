package handler

import (
	"embed"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"path"
	"strings"

	"fixable.com/fixable/internal/models"
	"fixable.com/fixable/internal/storage"
	"fixable.com/fixable/internal/utils"
)

//go:embed templates/*
var TemplatesFS embed.FS

type ServicioHandler struct {
	_servicioStorage   storage.IServicioStorage
	_categoriaStorage  storage.ICategoriaStorage
	_blobStorage       storage.IBlobStorage
	_comentarioStorage storage.IComentarioStorage
}

func (h *ServicioHandler) PromocionarseHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(TemplatesFS,
		"templates/base.templ",
		"templates/navbar/navbar.templ",
		"templates/promocionarse/promocionarse.templ",
	)
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"error": err})
		return

	}
	t.Execute(w, nil)
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
	t, err := template.ParseFS(TemplatesFS,
		"templates/base.templ",
		"templates/navbar/navbar.templ",
		"templates/servicios/home.templ",
		"templates/categorias/list.templ",
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

func (h *ServicioHandler) GetAdminAllServicios(w http.ResponseWriter, r *http.Request) {
	services, _ := h._servicioStorage.GetServices()
	comentarios, _ := h._comentarioStorage.GetAllComentarios()
	categorias, _ := h._categoriaStorage.GetCategorias()
	t, err := template.ParseFS(TemplatesFS,
		"templates/base.templ",
		"templates/navbar/navbar.templ",
		"templates/admin/home.templ")

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
	slog.Info("Servicio", slog.AnyValue(servicio))
	t, err := template.ParseFS(TemplatesFS,
		"templates/base.templ",
		"templates/navbar/navbar.templ",
		"templates/servicios/servicio.templ",
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

func (h *ServicioHandler) CreateServicio(w http.ResponseWriter, r *http.Request) {
	var err error
	var imageUrl string
	servicio := new(models.Servicio)
	utils.ToStruct(r, servicio)
	file, headers, _ := r.FormFile("file")
	if urlId, err := utils.GetIdFromParams(r); err == nil {
		servicio.ID = urlId
	}
	if file != nil {
		defer file.Close()
		headers.Filename = strings.ReplaceAll(headers.Filename, " ", "")
		imageExtension := path.Ext(headers.Filename)
		url, err := h._blobStorage.UploadImage(r.Context(), file, headers.Filename, fmt.Sprintf("image/%s", imageExtension))
		if err != nil {
			utils.WriteResponse(w, 400, utils.Response{"error": err.Error()})
		}
		slog.Info("image uploaded", slog.String("url", url))
		imageUrl = url
	}
	if servicio.ID == 0 {
		servicio.Imagen = imageUrl
		err = h._servicioStorage.CreateService(servicio)
		if err != nil {
			utils.WriteResponse(w, http.StatusBadRequest, utils.Response{"error": err.Error()})
			return
		}
	} else {
		prev, err := h._servicioStorage.GetServiceById(servicio.ID)
		if err != nil {
			utils.WriteResponse(w, http.StatusNotFound, utils.Response{"error": "Servicio no encontrado"})
			return
		}
		if imageUrl != "" {
			servicio.Imagen = imageUrl
		} else {
			servicio.Imagen = prev.Imagen
		}
		err = h._servicioStorage.UpdateServicio(servicio)
		if err != nil {
			utils.WriteResponse(w, http.StatusBadRequest, utils.Response{"error": err.Error()})
			return
		}
	}
	slog.Info("CreateServicioHandler Done!")
	http.Redirect(w, r, "/admin/servicios", http.StatusSeeOther)
}

func (h *ServicioHandler) CreateForm(w http.ResponseWriter, r *http.Request) {
	id, _ := utils.GetIdFromParams(r)
	var servicio *models.Servicio
	if id != 0 {
		servicio, _ = h._servicioStorage.GetServiceById(id)
	}
	if servicio == nil {
		servicio = new(models.Servicio)
	}
	t, err := template.ParseFS(TemplatesFS,
		"templates/base.templ",
		"templates/navbar/navbar.templ",
		"templates/admin/create.templ",
	)
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"message": err})
	}
	categorias, err := h._categoriaStorage.GetCategorias()
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"message": err})
	}
	data := struct {
		Servicio   *models.Servicio
		Categorias *[]models.Categoria
	}{
		Servicio:   servicio,
		Categorias: categorias,
	}

	err = t.Execute(w, data)
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"message": err})
	}

}

func (h *ServicioHandler) DeleteServicio(w http.ResponseWriter, r *http.Request) {
	id, _ := utils.GetIdFromParams(r)

	slog.Info("Delete servicio", slog.Any("id", id))
	servicio, err := h._servicioStorage.GetServiceById(id)
	if err != nil {
		utils.WriteResponse(w, 404, utils.Response{"error": err})
		return
	}
	err = h._servicioStorage.Delete(servicio)
	if err != nil {
		utils.WriteResponse(w, 400, utils.Response{"error": err})
		return
	}
	http.Redirect(w, r, "/admin/servicios/", http.StatusSeeOther)
}

func NewServiceHandler(
	serviciosRepositorio storage.IServicioStorage,
	comentarioStorage storage.IComentarioStorage,
	blobStorage storage.IBlobStorage,
	categoriaStorage storage.ICategoriaStorage,
) *ServicioHandler {
	return &ServicioHandler{
		_categoriaStorage:  categoriaStorage,
		_blobStorage:       blobStorage,
		_comentarioStorage: comentarioStorage,
		_servicioStorage:   serviciosRepositorio,
	}
}

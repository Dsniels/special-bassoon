package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"fixable.com/fixable/internal/models"
	"fixable.com/fixable/internal/repositories"
	"fixable.com/fixable/internal/utils"
)

type ServicioHandler struct {
	_servicioRepository  repositories.IServicioRepository
	_categoriaRepository repositories.ICategoriaRepository
}

func NewServiceHandler(serviciosRepositorio repositories.IServicioRepository, categoriaRepository repositories.ICategoriaRepository) *ServicioHandler {
	return &ServicioHandler{
		_categoriaRepository: categoriaRepository,
		_servicioRepository:  serviciosRepositorio,
	}
}

func (h *ServicioHandler) CreateServicioHandler(w http.ResponseWriter, r *http.Request) {
	var servicio models.Servicio
	servicio.Nombre = r.FormValue("nombre")
	servicio.Direccion = r.FormValue("direccion")
	categoriastr := r.FormValue("categoria")
	numstr := r.FormValue("telefono")
	telefonoInt, err := strconv.ParseInt(numstr, 10, 64)
  fmt.Println("Telefono: ",numstr)
	if err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, utils.Response{"error": err})
		return
	}

	categoriaInt, err := strconv.ParseUint(categoriastr, 10, 32)
	if err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, utils.Response{"error": err})
		return
	}
	telefonoInt32 := int(telefonoInt)
	servicio.Telefono = &telefonoInt32
	servicio.CategoriaId = uint(categoriaInt)
  fmt.Println("Categoria: ",servicio.CategoriaId)
  fmt.Println("Telefono: ",servicio.Telefono)
	err = servicio.Validate()
	if err != nil {
		fmt.Println("%w", err)
		utils.WriteResponse(w, http.StatusBadRequest, utils.Response{"error": err})
		return
	}
	err = h._servicioRepository.CreateService(&servicio)
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"error": err})
		return
	}
	fmt.Printf("%v", servicio)
}

func (h *ServicioHandler) NewServicio(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("internal/templates/base.template", "internal/templates/servicios/create.template")
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"error": err})
		return
	}
	categorias, err := h._categoriaRepository.GetCategorias()
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"error": err})
		return
	}
	data := struct {
		Options []models.Categoria
	}{
		Options: *categorias,
	}
	t.Execute(w, data)
}

func (h *ServicioHandler) GetAllServicios(w http.ResponseWriter, r *http.Request) {
	services, err := h._servicioRepository.GetServices()
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"error": err})
		return
	}

	t, err := template.ParseFiles("internal/templates/base.template", "internal/templates/servicios/home.template")
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, services)
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
	servicio, err := h._servicioRepository.GetServiceById(id)

	utils.WriteResponse(w, http.StatusOK, utils.Response{"data": servicio})

	return

}

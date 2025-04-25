package handler

import (
	"encoding/json"
	"html/template"
	"net/http"

	"fixable.com/fixable/internal/models"
	"fixable.com/fixable/internal/repositories"
	"fixable.com/fixable/internal/utils"
)

type ServicioHandler struct {
	_servicioRepository repositories.IServicioRepository
}

func NewServiceHandler(serviciosRepositorio repositories.IServicioRepository) *ServicioHandler {

	return &ServicioHandler{
		_servicioRepository: serviciosRepositorio,
	}
}

func (h *ServicioHandler) CreateServicioHandler(w http.ResponseWriter, r *http.Request) {
	var servicio models.Servicio
	err := json.NewDecoder(r.Body).Decode(&servicio)
	if err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, utils.Response{"error": err})
		return
	}

	err = h._servicioRepository.CreateService(&servicio)
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"error": err})
		return
	}
	utils.WriteResponse(w, http.StatusOK, utils.Response{"data": servicio})
	return

}

func (h *ServicioHandler) GetAllServicios(w http.ResponseWriter, r *http.Request) {
	services, err := h._servicioRepository.GetServices()
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, utils.Response{"error": err})
		return
	}

	// utils.WriteResponse(w, http.StatusOK, utils.Response{"data": services})
  t,err := template.ParseFiles("internal/templates/home.html")
  if err !=nil{
    panic(err)
    return
  }

  err =  t.Execute(w,services)
  if err !=nil{
    return
  }
  // return

}


func (h *ServicioHandler) GetServicioById(w http.ResponseWriter, r *http.Request){
  id, err := utils.GetIdFromParams(r)
  if err != nil{
		utils.WriteResponse(w, http.StatusBadRequest, utils.Response{"error": err})
		return
  }
  servicio, err:= h._servicioRepository.GetServiceById(id)

	utils.WriteResponse(w, http.StatusOK, utils.Response{"data": servicio})

  return

    


}






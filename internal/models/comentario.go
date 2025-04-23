package models

import "gorm.io/gorm"

type Comentario struct {
	gorm.Model
  ServicioId uint `json:"servicioId"`
  Comentario string `json:"comentario"`
}

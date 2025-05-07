package models

import (
	"gorm.io/gorm"
)

type Servicio struct {
	gorm.Model
	Nombre      string `json:"nombre"`
	Direccion   string `json:"direccion"`
	Email       string `json:"email"`
	Descripcion string `json:"descripcion"`
	Telefono    string `json:"telefono"`
	CategoriaId uint   `json:"categoria_id"`
	Imagen      string `json:"imagen"`
}

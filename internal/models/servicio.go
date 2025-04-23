package models

import "gorm.io/gorm"

type Servicio struct {
	gorm.Model
  Nombre string `json:"nombre"`
  Direccion string `json:"direccion"`  
  Email string `json:"email"`
  Telefono *int `json:"telefono"`
  CategoriaId uint `json:"categoriaId"`
  Categoria Categoria `json:"-" gorm:"foreignKey:CategoriaId`


}




package models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Servicio struct {
	gorm.Model
	Nombre      string `json:"nombre"`
	Direccion   string `json:"direccion"`
	Email       string `json:"email"`
	Telefono    *int   `json:"telefono"`
	CategoriaId uint   `json:"categoriaId"`
}

func (s *Servicio) Validate() error {
	if strings.TrimSpace(s.Nombre) == "" {
		return errors.New("Nombre no puede estar vario")
	}
	if strings.TrimSpace(s.Direccion) == "" {
		return errors.New("Direccion no puede estar vario")
	}
	if s.Telefono == nil {
		return errors.New("Telefono no puede estar vario")
	}
	if s.CategoriaId == 0 {
		return errors.New("Categoria no puede estar vario")
	}
	return nil
}

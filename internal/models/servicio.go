package models

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type Servicio struct {
	gorm.Model
	Nombre      string
	Direccion   string
	Email       string
	Descripcion string
	Telefono    string
	CategoriaId uint
	Imagen      string
}

func (s *Servicio) Validate() error {
	if strings.TrimSpace(s.Nombre) == "" {
		return errors.New("Nombre no puede estar vario")
	}
	if strings.TrimSpace(s.Direccion) == "" {
		return errors.New("Direccion no puede estar vario")
	}
	if s.CategoriaId == 0 {
		return errors.New("Categoria no puede estar vario")
	}
	return nil
}

func (s *Servicio) FromForm(r *http.Request) error {

	s.Nombre = r.FormValue("nombre")
	s.Direccion = r.FormValue("direccion")
	categoriastr := r.FormValue("categoria")
	categoriaInt, err := strconv.Atoi(categoriastr)
	if err != nil {
		return err
	}
	s.Imagen = r.FormValue("imagen")
	s.Telefono = r.FormValue("telefono")
	s.CategoriaId = uint(categoriaInt)
	return nil
}

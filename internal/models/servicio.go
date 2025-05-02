package models

import (
	"errors"
	"html/template"
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
	Telefono    *int
	CategoriaId uint
	Contenido   template.HTML
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

func (s *Servicio) FromForm(r *http.Request) error {

	s.Nombre = r.FormValue("nombre")
	s.Direccion = r.FormValue("direccion")
	categoriastr := r.FormValue("categoria")
	numstr := r.FormValue("telefono")
	content := r.FormValue("contenido")
	telefonoInt, err := strconv.Atoi(numstr)
	if err != nil {
		return err
	}

	categoriaInt, err := strconv.Atoi(categoriastr)
	if err != nil {
		return err
	}
	s.Contenido = template.HTML(content)
	s.Telefono = &telefonoInt
	s.CategoriaId = uint(categoriaInt)
	return nil
}

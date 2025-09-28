package models

type Servicio struct {
	ID          uint   `gorm:"primarykey"`
	Nombre      string `json:"nombre"`
	Direccion   string `json:"direccion"`
	Email       string `json:"email"`
	Descripcion string `json:"descripcion"`
	Telefono    string `json:"telefono"`
	CategoriaId uint   `json:"categoria_id"`
	Imagen      string `json:"imagen"`
}

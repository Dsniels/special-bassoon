package models

type Categoria struct {
	ID     uint   `gorm:"primarykey"`
	Nombre string `json:"nombre"`
}

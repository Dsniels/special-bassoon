package models

type Comentario struct {
	ID         uint   `gorm:"primarykey"`
	ServicioId uint   `json:"servicioId"`
	Comentario string `json:"comentario"`
}

package storage

import (
	"fixable.com/fixable/internal/models"
	"gorm.io/gorm"
)

type ComentarioStorage struct {
	db *gorm.DB
}

type IComentarioStorage interface {
	CreateComentario(*models.Comentario) error
	GetAllComentarios() (*[]models.Comentario, error)
	GetComentariosPorServicioId(servicioId int) (*[]models.Comentario, error)
}

func NewComentarioStorage(db *gorm.DB) *ComentarioStorage {
	return &ComentarioStorage{
		db: db,
	}
}

func (repo *ComentarioStorage) CreateComentario(comentario *models.Comentario) error {
	result := repo.db.Create(comentario)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *ComentarioStorage) GetComentariosPorServicioId(servicioId int) (*[]models.Comentario, error) {
	var comentarios []models.Comentario
	results := repo.db.Where("servicio_Id = ?", servicioId).Find(&comentarios)
	if results.Error != nil {
		return nil, results.Error
	}
	return &comentarios, nil

}

func (repo *ComentarioStorage) GetAllComentarios() (*[]models.Comentario, error) {
	var comentarios []models.Comentario
	result := repo.db.Find(&comentarios)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comentarios, nil
}

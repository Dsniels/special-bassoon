package repositories

import (
	"fixable.com/fixable/internal/models"
	"gorm.io/gorm"
)

type ComentarioRepository struct {
	db *gorm.DB
}

type IComentarioRepository interface {
	CreateComentario(*models.Comentario) error
	GetAllComentarios() (*[]models.Comentario, error)
}

func NewComentarioRepository(db *gorm.DB) *ComentarioRepository {
	return &ComentarioRepository{
		db: db,
	}
}

func (repo *ComentarioRepository) CreateComentario(comentario *models.Comentario) error {
	result := repo.db.Create(comentario)
  if result.Error != nil{
    return result.Error
  }
	return nil
}

func (repo *ComentarioRepository) GetAllComentarios()(*[]models.Comentario, error){
  var comentarios []models.Comentario
  result := repo.db.Find(&comentarios)
  if result.Error != nil{
    return nil, result.Error
  }
  return &comentarios, nil
}

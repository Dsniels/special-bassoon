package storage 

import (
	"fixable.com/fixable/internal/models"
	"gorm.io/gorm"
)

type CategoriaStorage struct {
	db *gorm.DB
}

type ICategoriaStorage interface {
	GetCategorias() (*[]models.Categoria, error)
	GetCategoriaById(uint) (*models.Categoria, error)
	Create(models.Categoria) error
}

func NewCategoriaStorage(db *gorm.DB) *CategoriaStorage {
	return &CategoriaStorage{
		db: db,
	}
}

func (repo *CategoriaStorage) Create(categoria models.Categoria) error {
	result := repo.db.Create(&categoria)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *CategoriaStorage) GetCategoriaById(id uint) (*models.Categoria, error) {
	var categoria models.Categoria
	result := repo.db.First(&categoria, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &categoria, nil
}

func (repo *CategoriaStorage) GetCategorias() (*[]models.Categoria, error) {
	var categoria []models.Categoria
	result := repo.db.Find(&categoria)
	if result.Error != nil {
		return nil, result.Error
	}
	return &categoria, nil
}

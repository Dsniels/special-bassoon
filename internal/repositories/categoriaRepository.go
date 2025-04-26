package repositories

import (
	"fixable.com/fixable/internal/models"
	"gorm.io/gorm"
)

type CategoriaRepository struct {
	Db *gorm.DB
}

type ICategoriaRepository interface {
	GetCategorias() (*[]models.Categoria, error)
	GetCategoriaById(uint) (*models.Categoria, error)
	Create(models.Categoria) error
}

func NewCategoriaRepository(db *gorm.DB) *CategoriaRepository {
	return &CategoriaRepository{
		Db: db,
	}
}

func (repo *CategoriaRepository) Create(categoria models.Categoria) error {
	result := repo.Db.Create(&categoria)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *CategoriaRepository) GetCategoriaById(id uint) (*models.Categoria, error) {
	var categoria models.Categoria
	result := repo.Db.First(&categoria, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &categoria, nil
}

func (repo *CategoriaRepository) GetCategorias() (*[]models.Categoria, error) {
	var categoria []models.Categoria
	result := repo.Db.Find(&categoria)
	if result.Error != nil {
		return nil, result.Error
	}
	return &categoria, nil
}

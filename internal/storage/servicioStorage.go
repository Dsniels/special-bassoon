package storage

import (
	"errors"

	"fixable.com/fixable/internal/models"
	"gorm.io/gorm"
)

type ServicioStorage struct {
	db *gorm.DB
}

type IServicioStorage interface {
	CreateService(*models.Servicio) error
	GetServiceByCategoriaId(uint) (*[]models.Servicio, error)
	GetServices() (*[]models.Servicio, error)
	GetServiceById(uint) (*models.Servicio, error)
	Delete(*models.Servicio) error
	GetByQuery(string) (*[]models.Servicio, error)
	UpdateServicio(servicio *models.Servicio) error
}

func NewServicioStorage(db *gorm.DB) *ServicioStorage {
	return &ServicioStorage{
		db: db,
	}
}

func (r *ServicioStorage) CreateService(servicio *models.Servicio) error {
	result := r.db.Create(&servicio)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ServicioStorage) Delete(servicio *models.Servicio) error {
	tx := r.db.Delete(servicio)
	return tx.Error
}

func (r *ServicioStorage) UpdateServicio(servicio *models.Servicio) error {
	return r.db.Save(servicio).Error
}

func (r *ServicioStorage) GetServiceById(Id uint) (*models.Servicio, error) {
	var servicio = models.Servicio{ID: Id}
	r.db.First(&servicio)
	return &servicio, nil

}

func (r *ServicioStorage) GetServiceByCategoriaId(id uint) (*[]models.Servicio, error) {
	var servicios []models.Servicio
	result := r.db.Where("categoria_id = ?", id).Find(&servicios)
	if result.Error != nil {
		return nil, result.Error
	}
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}

	return &servicios, nil
}

func (r *ServicioStorage) GetByQuery(query string) (*[]models.Servicio, error) {
	var servicios []models.Servicio
	result := r.db.Where("descripcion LIKE ? or nombre LIKE ?", "%"+query+"%", "%"+query+"%").Find(&servicios)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return &servicios, nil
}

func (r *ServicioStorage) GetServices() (*[]models.Servicio, error) {
	var servicios []models.Servicio
	result := r.db.Order("NEWID()").Find(&servicios)
	if result.Error != nil {
		return nil, result.Error
	}
	return &servicios, nil
}

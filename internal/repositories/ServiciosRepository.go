package repositories

import (
	"fixable.com/fixable/internal/models"
	"gorm.io/gorm"
)

type ServicioRepository struct {
	Db *gorm.DB
}

type IServicioRepository interface {
	CreateService(*models.Servicio) error
	GetServices() (*[]models.Servicio, error)
	GetServiceById(uint) (*models.Servicio, error)
}

func NewServicioRepository(db *gorm.DB) *ServicioRepository {
	return &ServicioRepository{
		Db: db,
	}
}

func (r *ServicioRepository) CreateService(servicio *models.Servicio) error {
	result := r.Db.Create(&servicio)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func (r *ServicioRepository) GetServiceById(Id uint) (*models.Servicio, error){
  var servicio = models.Servicio{Model: gorm.Model{ID: Id}}
  r.Db.First(&servicio)
  return &servicio, nil

}

func (r *ServicioRepository) GetServices() (*[]models.Servicio, error){
  var servicios []models.Servicio
  result := r.Db.Find(servicios)
  if result.Error != nil{
    return nil,result.Error
  }

  return &servicios, nil

}

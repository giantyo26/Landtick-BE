package repositories

import (
	"server/models"

	"gorm.io/gorm"
)

type StationRepository interface {
	FindAllStations() ([]models.Station, error)
	GetStation(ID int) (models.Station, error)
	CreateStation(station models.Station) (models.Station, error)
	UpdateStation(station models.Station) (models.Station, error)
	DeleteStation(station models.Station) (models.Station, error)
}

func RepositoryStation(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllStations() ([]models.Station, error) {
	var stations []models.Station

	err := r.db.Find(&stations).Error

	return stations, err
}

func (r *repository) GetStation(ID int) (models.Station, error) {
	var station models.Station
	err := r.db.First(&station, ID).Error

	return station, err
}

func (r *repository) CreateStation(station models.Station) (models.Station, error) {
	err := r.db.Create(&station).Error

	return station, err
}

func (r *repository) UpdateStation(station models.Station) (models.Station, error) {
	err := r.db.Save(&station).Error

	return station, err
}

func (r *repository) DeleteStation(station models.Station) (models.Station, error) {
	err := r.db.Delete(&station).Error

	return station, err
}

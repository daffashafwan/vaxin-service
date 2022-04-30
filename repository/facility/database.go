package facility

import (
	"context"
	"errors"
	"github.com/daffashafwan/vaxin-service/business/facilities"
	"gorm.io/gorm"
)

type FacilityRepo struct {
	DB *gorm.DB
}

func CreateFacilityRepo(conn *gorm.DB) facilities.Repository {
	return &FacilityRepo{
		DB: conn,
	}
}

func (rep *FacilityRepo) Create(ctx context.Context, faci *facilities.Domain) (facilities.Domain, error) {
	fac := Facility{
		Name:      faci.Name,
		Latitude:  faci.Latitude,
		Longitude: faci.Longitude,
		UrlMaps:   faci.UrlMaps,
	}
	err := rep.DB.Create(&fac)
	if err.Error != nil {
		return facilities.Domain{}, err.Error
	}
	return fac.ToDomain(), nil
}

func (rep *FacilityRepo) Update(ctx context.Context, faci facilities.Domain) (facilities.Domain, error) {
	data := FromDomain(faci)
	err := rep.DB.Table("facilities").First(&data)
	if err.Error != nil {
		return facilities.Domain{}, err.Error
	}
	data.Name = faci.Name
	data.Latitude = faci.Latitude
	data.Longitude = faci.Longitude
	data.UrlMaps = faci.UrlMaps

	if rep.DB.Save(&data).Error != nil {
		return facilities.Domain{}, errors.New("bad requests")
	}
	return data.ToDomain(), nil
}

func (rep *FacilityRepo) GetAll(ctx context.Context) ([]facilities.Domain, error) {
	var data []Facility
	err := rep.DB.Table("facilities").Find(&data)
	if err.Error != nil {
		return []facilities.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (rep *FacilityRepo) GetById(ctx context.Context, id int) (facilities.Domain, error) {
	var data Facility
	err := rep.DB.Table("facilities").Find(&data, "id=?", id)
	if err.Error != nil {
		return facilities.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}

func (rep *FacilityRepo) Delete(ctx context.Context, id int) error {
	fac := Facility{}
	err := rep.DB.Table("facilities").Where("id = ?", id).First(&fac).Delete(&fac)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}

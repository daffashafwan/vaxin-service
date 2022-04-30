package vaccine

import (
	"context"
	"errors"
	"github.com/daffashafwan/vaxin-service/business/vaccines"
	"gorm.io/gorm"
)

type VaccineRepo struct {
	DB *gorm.DB
}

func CreateVaccineRepo(conn *gorm.DB) vaccines.Repository {
	return &VaccineRepo{
		DB: conn,
	}
}

func (rep *VaccineRepo) Create(ctx context.Context, vacc *vaccines.Domain) (vaccines.Domain, error) {
	vac := Vaccine{
		Name: vacc.Name,
		Type:  vacc.Type,
	}
	err := rep.DB.Create(&vac)
	if err.Error != nil {
		return vaccines.Domain{}, err.Error
	}
	return vac.ToDomain(), nil
}

func (rep *VaccineRepo) Update(ctx context.Context, vacc vaccines.Domain) (vaccines.Domain, error) {
	data := FromDomain(vacc)
	err := rep.DB.Table("vaccines").First(&data)
	if err.Error != nil {
		return vaccines.Domain{}, err.Error
	}
	data.Name = vacc.Name
	data.Type = vacc.Type

	if rep.DB.Save(&data).Error != nil {
		return vaccines.Domain{}, errors.New("bad requests")
	}
	return data.ToDomain(), nil
}

func (rep *VaccineRepo) GetAll(ctx context.Context) ([]vaccines.Domain, error) {
	var data []Vaccine
	err := rep.DB.Table("vaccines").Find(&data)
	if err.Error != nil {
		return []vaccines.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (rep *VaccineRepo) GetById(ctx context.Context, id int) (vaccines.Domain, error) {
	var data Vaccine
	err := rep.DB.Table("vaccines").Find(&data, "id=?", id)
	if err.Error != nil {
		return vaccines.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}

func (rep *VaccineRepo) Delete(ctx context.Context, id int) error {
	vacc := Vaccine{}
	err := rep.DB.Table("vaccines").Where("id = ?", id).First(&vacc).Delete(&vacc)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}
package admin

import (
	"context"
	"errors"

	"github.com/daffashafwan/vaxin-service/business/admins"
	"github.com/daffashafwan/vaxin-service/helpers/encrypt"
	"gorm.io/gorm"
)

type AdminRepo struct {
	DB *gorm.DB
}

func CreateAdminRepo(gormDb *gorm.DB) admins.Repository {
	return &AdminRepo{
		DB: gormDb,
	}
}

func (repo *AdminRepo) Login(ctx context.Context, username string, password string) (admins.Domain, error) {
	var adm Admin
	result := repo.DB.Table("admins").Where("username = ?", username).First(&adm).Error

	if result != nil {
		return admins.Domain{}, result
	}
	if !(encrypt.Compare(password,adm.Password)) {
		return admins.Domain{}, errors.New("password tidak cocok")
	}
	return adm.ToDomain(), nil

}

func (repo *AdminRepo) GetById(ctx context.Context, id int) (admins.Domain, error) {
	var data Admin
	err := repo.DB.Table("admins").Find(&data, "id=?", id)
	if err.Error != nil {
		return admins.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}
package users

import (
	"context"
	"github.com/daffashafwan/vaxin-service/business/users"
	"github.com/daffashafwan/vaxin-service/helpers/encrypt"
	"gorm.io/gorm"
	"errors"
)

type UserRepo struct {
	DB *gorm.DB
}

func CreateUserRepo(conn *gorm.DB) users.Repository {
	return &UserRepo{
		DB: conn,
	}
}

func (rep *UserRepo) Login(ctx context.Context, username string, password string) (users.Domain, error) {
	var user User
	result := rep.DB.Table("users").Where("username = ?", username).Where("status = ? ", "1").First(&user).Error

	if result != nil {
		return users.Domain{}, result
	}
	if !(encrypt.Compare(password,user.Password)) {
		return users.Domain{}, errors.New("Password tidak cocok")
	}
	return user.ToDomain(), nil

}

func (rep *UserRepo) Create(ctx context.Context,userR *users.Domain) (users.Domain, error) {
	user := User{
		Name:     userR.Name,
		Email:    userR.Email,
		Username: userR.Username,
		Password: userR.Password,
		Status: userR.Status,
		Token: userR.Token,
	}
	err := rep.DB.Create(&user)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return user.ToDomain(), nil
}

func (rep *UserRepo) Update(ctx context.Context, userU users.Domain) (users.Domain, error) {
	data := FromDomain(userU)
	err := rep.DB.Table("users").First(&data)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	data.Name = userU.Name
	data.Username = userU.Username
	data.Password = userU.Password
	data.Status = userU.Status
	data.Email = userU.Email
	data.Token = userU.Token
	

	if rep.DB.Save(&data).Error != nil {
		return users.Domain{}, errors.New("bad requests")
	}
	return data.ToDomain(), nil
}

func (rep *UserRepo) GetAll(ctx context.Context) ([]users.Domain, error) {
	var data []User
	err := rep.DB.Table("users").Find(&data)
	if err.Error != nil {
		return []users.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (rep *UserRepo) GetById(ctx context.Context, id int) (users.Domain, error) {
	var data User
	err := rep.DB.Table("users").Find(&data, "id=?", id)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}


func (rep *UserRepo) GetByToken(ctx context.Context, token string) (users.Domain, error) {
	var data User
	err := rep.DB.Table("users").Find(&data, "token=?", token)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}

func (rep *UserRepo) Delete(ctx context.Context, id int) error {
	user := User{}
	err := rep.DB.Table("users").Where("id = ?", id).First(&user).Delete(&user)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}

func (rep *UserRepo) GetByEmail(ctx context.Context, email string) (users.Domain, error) {
	var data User
	err := rep.DB.Table("users").Find(&data, "email=?", email)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}

func (rep *UserRepo) GetByUsername(ctx context.Context, username string) (users.Domain, error) {
	var data User
	err := rep.DB.Table("users").Find(&data, "username=?", username)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}
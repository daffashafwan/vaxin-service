package users

import (
	"time"

	"github.com/daffashafwan/vaxin-service/business/users"
)

type User struct {
	Id        int    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string `gorm:"size:255;not null" json:"name"`
	Username  string `gorm:"size:255;not null;unique" json:"username"`
	Password  string `gorm:"size:100;not null;" json:"password"`
	Status    string `gorm:"size:100;not null;" json:"status"`
	Token     string `gorm:"size:100;unique;" json:"token"`
	Email     string `gorm:"size:100;not null;unique" json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) ToDomain() users.Domain {
	return users.Domain{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToListDomain(data []User) (result []users.Domain) {
	result = []users.Domain{}
	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return
}

func FromDomain(domain users.Domain) User {
	return User{
		Id:        domain.Id,
		Name:      domain.Name,
		Email:     domain.Email,
		Password:  domain.Password,
		Username:  domain.Username,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

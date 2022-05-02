package admin

import (
	"time"

	"github.com/daffashafwan/vaxin-service/business/admins"
)

type Admin struct {
	Id        int    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string `gorm:"size:255;not null" json:"name"`
	Username  string `gorm:"size:255;not null;unique" json:"username"`
	Password  string `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ad *Admin) ToDomain() admins.Domain {
	return admins.Domain {
		Id:        ad.Id,
		Name:      ad.Name,
		Username:  ad.Username,
		Password:  ad.Password,
		CreatedAt: ad.CreatedAt,
		UpdatedAt: ad.UpdatedAt,
	}
}

func FromDomain(domain admins.Domain) Admin {
	return Admin{
		Id:        domain.Id,
		Name:      domain.Name,
		Password:  domain.Password,
		Username:  domain.Username,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

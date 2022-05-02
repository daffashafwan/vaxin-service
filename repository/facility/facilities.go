package facility

import (
	"github.com/daffashafwan/vaxin-service/business/facilities"
	"time"
)

type Facility struct {
	Id        int    `gorm:"primaryKey;auto_increment" json:"id"`
	Name      string `gorm:"size:255;not null" json:"name"`
	Latitude  string `gorm:"size:255;not null" json:"latitude"`
	Longitude string `gorm:"size:255;not null" json:"longitude"`
	UrlMaps   string `gorm:"size:255;" json:"url_maps"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (fac *Facility) ToDomain() facilities.Domain {
	return facilities.Domain{
		Id:        fac.Id,
		Name:      fac.Name,
		Latitude:  fac.Latitude,
		Longitude: fac.Longitude,
		UrlMaps:   fac.UrlMaps,
		CreatedAt: fac.CreatedAt,
		UpdatedAt: fac.UpdatedAt,
	}
}

func ToListDomain(data []Facility) (result []facilities.Domain) {
	result = []facilities.Domain{}
	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return
}

func FromDomain(domain facilities.Domain) Facility {
	return Facility{
		Id:        domain.Id,
		Name:      domain.Name,
		Latitude:  domain.Latitude,
		Longitude: domain.Longitude,
		UrlMaps:   domain.UrlMaps,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

package vaccine

import (
	"github.com/daffashafwan/vaxin-service/business/vaccines"
	"time"
)

type Vaccine struct {
	Id        int    `gorm:"primaryKey;auto_increment" json:"id"`
	Name      string `gorm:"size:255;not null" json:"name"`
	Type      string `gorm:"size:255;not null" json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (vacc *Vaccine) ToDomain() vaccines.Domain {
	return vaccines.Domain{
		Id:        vacc.Id,
		Name:      vacc.Name,
		Type:      vacc.Type,
		CreatedAt: vacc.CreatedAt,
		UpdatedAt: vacc.UpdatedAt,
	}
}

func ToListDomain(data []Vaccine) (result []vaccines.Domain) {
	result = []vaccines.Domain{}
	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return
}

func FromDomain(domain vaccines.Domain) Vaccine {
	return Vaccine{
		Id:        domain.Id,
		Name:      domain.Name,
		Type:      domain.Type,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

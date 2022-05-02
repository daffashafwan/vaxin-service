package vaccination

import (
	"github.com/daffashafwan/vaxin-service/business/vaccinations"
	"github.com/daffashafwan/vaxin-service/repository/event"
	users "github.com/daffashafwan/vaxin-service/repository/user"
	"time"
)

type Vaccination struct {
	Id        int         `gorm:"primaryKey;auto_increment" json:"id"`
	SpecialId string   `gorm:"size:255;not null;unique" json:"special_id"`
	EventId   int         `gorm:"size:255;not null" json:"event_id"`
	Event     event.Event `gorm:"foreignKey:EventId;association_foreignkey:Id"`
	UserId    int         `gorm:"size:255;not null" json:"user_id"`
	User      users.User  `gorm:"foreignKey:UserId;association_foreignkey:Id"`
	Queue     int         `gorm:"size:255;" json:"queue"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func (item *Vaccination) ToDomain() vaccinations.Domain {
	return vaccinations.Domain{
		Id:        item.Id,
		SpecialId: item.SpecialId,
		EventId:   item.EventId,
		Event:     item.Event,
		UserId:    item.UserId,
		User:      item.User,
		Queue:     item.Queue,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}

func ToListDomain(data []Vaccination) (result []vaccinations.Domain) {
	result = []vaccinations.Domain{}
	for _, item := range data {
		result = append(result, item.ToDomain())
	}
	return
}

func FromDomain(domain vaccinations.Domain) Vaccination {
	return Vaccination{
		Id:        domain.Id,
		SpecialId: domain.SpecialId,
		EventId:   domain.EventId,
		UserId:    domain.UserId,
		Queue:     domain.Queue,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

package event

import (
	"time"
	"github.com/daffashafwan/vaxin-service/business/events"
	"github.com/daffashafwan/vaxin-service/repository/facility"
	"github.com/daffashafwan/vaxin-service/repository/vaccine"
)

type Event struct {
	Id         int               `gorm:"primaryKey;auto_increment" json:"id"`
	Name       string            `gorm:"size:255;not null" json:"name"`
	VaccineId  int               `gorm:"size:255;not null" json:"vaccine_id"`
	Vaccine    vaccine.Vaccine   `gorm:"foreignKey:VaccineId;association_foreignkey:Id"`
	FacilityId int               `gorm:"size:255;not null" json:"facility_id"`
	Facility   facility.Facility `gorm:"foreignKey:FacilityId;association_foreignkey:Id"`
	Quota      int               `gorm:"size:255;" json:"quota"`
	StartDate  time.Time         `gorm:"not null" json:"start_date"`
	EndDate    time.Time         `gorm:"not null" json:"end_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (item *Event) ToDomain() events.Domain {
	return events.Domain{
		Id:         item.Id,
		Name:       item.Name,
		FacilityId: item.FacilityId,
		Facility:   item.Facility,
		VaccineId:  item.VaccineId,
		Vaccine:    item.Vaccine,
		Quota:      item.Quota,
		StartDate:  item.StartDate.Format("2006-01-02T15:04:05.000Z"),
		EndDate:    item.EndDate.Format("2006-01-02T15:04:05.000Z"),
		CreatedAt:  item.CreatedAt,
		UpdatedAt:  item.UpdatedAt,
	}
}

func ToListDomain(data []Event) (result []events.Domain) {
	result = []events.Domain{}
	for _, item := range data {
		result = append(result, item.ToDomain())
	}
	return
}

func FromDomain(domain events.Domain) Event {
	layout := "2006-01-02T15:04:05.000Z"
	start, _ := time.Parse(layout, domain.StartDate)
	end, _ := time.Parse(layout, domain.EndDate)
	return Event{
		Id:         domain.Id,
		Name:       domain.Name,
		FacilityId: domain.FacilityId,
		VaccineId:  domain.VaccineId,
		StartDate:  start,
		EndDate:    end,
		Quota:      domain.Quota,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}

package requests

import (
	"github.com/daffashafwan/vaxin-service/business/events"
)

type EventRequest struct {
	Name       string `json:"name"`
	FacilityId int    `json:"facility_id"`
	VaccineId  int    `json:"vaccine_id"`
	Quota      int    `json:"quota"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
}

type ItemRequestQuota struct {
	Quota int `json:"quota"`
}

func (item *ItemRequestQuota) QuotaToDomain() events.Domain {
	return events.Domain{
		Quota: item.Quota,
	}
}

func (item *EventRequest) ToDomain() events.Domain {
	return events.Domain{
		FacilityId: item.FacilityId,
		VaccineId:  item.VaccineId,
		Name:       item.Name,
		StartDate:  item.StartDate,
		EndDate:    item.EndDate,
		Quota:      item.Quota,
	}
}

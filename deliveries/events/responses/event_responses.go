package responses

import (
	"github.com/daffashafwan/vaxin-service/business/events"
	"strconv"
)

type EventResponse struct {
	Id         int         `json:"id"`
	Name       string      `json:"name"`
	FacilityId int         `json:"facility_id"`
	Facility   interface{} `json:"facility"`
	VaccineId  int         `json:"vaccine_id"`
	Vaccine    interface{} `json:"vaccine"`
	Quota      string      `json:"quota"`
	StartDate  string      `json:"start_date"`
	EndDate    string      `json:"end_date"`
}

func FromDomain(domain events.Domain) EventResponse {
	quotaR := strconv.Itoa(domain.Quota)
	return EventResponse{
		Id:         domain.Id,
		FacilityId: domain.FacilityId,
		Facility:   domain.Facility,
		VaccineId:  domain.VaccineId,
		Vaccine:    domain.Vaccine,
		Name:       domain.Name,
		StartDate:  domain.StartDate,
		EndDate:    domain.EndDate,
		Quota:      quotaR,
	}
}

func FromListDomain(domain []events.Domain) []EventResponse {
	var list []EventResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}

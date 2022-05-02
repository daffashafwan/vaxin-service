package responses

import (
	"github.com/daffashafwan/vaxin-service/business/vaccinations"
)

type VaccinationResponse struct {
	Id        int         `json:"id"`
	SpecialId string      `json:"special_id"`
	EventId   int         `json:"event_id"`
	Event     interface{} `json:"event"`
	UserId    int         `json:"user_id"`
	User      interface{} `json:"user"`
	Queue     int         `json:"queue"`
}

func FromDomain(domain vaccinations.Domain) VaccinationResponse {
	return VaccinationResponse{
		Id:        domain.Id,
		EventId:   domain.EventId,
		Event:     domain.Event,
		UserId:    domain.UserId,
		User:      domain.User,
		SpecialId: domain.SpecialId,
		Queue:     domain.Queue,
	}
}

func FromListDomain(domain []vaccinations.Domain) []VaccinationResponse {
	var list []VaccinationResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}

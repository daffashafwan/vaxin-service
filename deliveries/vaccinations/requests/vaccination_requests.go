package requests

import (
	"github.com/daffashafwan/vaxin-service/business/vaccinations"
)

type VaccinationRequest struct {
	EventId int    `json:"event_id"`
	UserId  int    `json:"user_id"`
}

func (vaccination *VaccinationRequest) ToDomain() vaccinations.Domain {
	return vaccinations.Domain{
		EventId: vaccination.EventId,
		UserId:  vaccination.UserId,
	}
}

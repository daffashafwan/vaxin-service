package responses

import (
	"time"

	"github.com/daffashafwan/vaxin-service/business/vaccines"
)

type VaccineResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain vaccines.Domain) VaccineResponse {
	return VaccineResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		Type:      domain.Type,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromListDomain(domain []vaccines.Domain) []VaccineResponse {
	var list []VaccineResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}

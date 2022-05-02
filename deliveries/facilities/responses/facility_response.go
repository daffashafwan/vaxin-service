package responses

import (
	"time"

	"github.com/daffashafwan/vaxin-service/business/facilities"
)

type FacilityResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Latitude  string    `json:"latitude"`
	Longitude string    `json:"longitude"`
	UrlMaps   string    `json:"url_maps"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain facilities.Domain) FacilityResponse {
	return FacilityResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		Latitude:  domain.Latitude,
		Longitude: domain.Longitude,
		UrlMaps:   domain.UrlMaps,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromListDomain(domain []facilities.Domain) []FacilityResponse {
	var list []FacilityResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}

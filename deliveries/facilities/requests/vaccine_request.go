package requests

import "github.com/daffashafwan/vaxin-service/business/facilities"

type Facility struct {
	Name      string `json:"name"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	UrlMaps   string `json:"url_maps"`
}

func (fac *Facility) ToDomain() facilities.Domain {
	return facilities.Domain{
		Name:      fac.Name,
		Latitude:  fac.Latitude,
		Longitude: fac.Longitude,
		UrlMaps:   fac.UrlMaps,
	}
}

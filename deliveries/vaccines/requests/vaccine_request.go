package requests

import "github.com/daffashafwan/vaxin-service/business/vaccines"

type Vaccine struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (vac *Vaccine) ToDomain() vaccines.Domain {
	return vaccines.Domain{
		Name: vac.Name,
		Type: vac.Type,
	}
}

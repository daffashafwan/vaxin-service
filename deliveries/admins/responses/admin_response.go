package responses

import (
	"time"

	 "github.com/daffashafwan/vaxin-service/business/admins"
)

type AdminResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	JWTToken  string    `json:"jwtToken"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain admins.Domain) AdminResponse {
	return AdminResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		Username: domain.Username,
		JWTToken:  domain.JWTToken,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}


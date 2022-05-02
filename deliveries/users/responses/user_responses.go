package responses

import (
	"github.com/daffashafwan/vaxin-service/business/users"
	"time"
)

type UserResponse struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	JWTToken  string    `json:"jwtToken"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		Name:      domain.Name,
		Username:  domain.Username,
		Email:     domain.Email,
		JWTToken:  domain.JWTToken,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromListDomain(domain []users.Domain) []UserResponse {
	var list []UserResponse
	for _, v := range domain {
		list = append(list, FromDomain(v))
	}
	return list
}

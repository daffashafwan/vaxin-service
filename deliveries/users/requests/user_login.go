package requests

import (
	"github.com/daffashafwan/vaxin-service/business/users"
)

type UserLogin struct {
	Username    string `json:"username"`
	Password string `json:"password"`
}

func (ul *UserLogin) ToDomain() users.Domain {
	return users.Domain{
		Username:    ul.Username,
		Password: ul.Password,
	}
}

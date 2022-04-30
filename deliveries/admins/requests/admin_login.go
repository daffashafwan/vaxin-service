package requests

import "github.com/daffashafwan/vaxin-service/business/admins"

type AdminLogin struct {
	Username    string `json:"username"`
	Password 	string `json:"password"`
}

func (al *AdminLogin) ToDomain() admins.Domain {
	return admins.Domain{
		Username:    al.Username,
		Password: al.Password,
	}
}

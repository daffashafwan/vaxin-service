package requests

import "github.com/daffashafwan/vaxin-service/business/users"

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (ur *UserRegister) ToDomain() users.Domain {
	return users.Domain{
		Username: ur.Username,
		Password: ur.Password,
		Email:    ur.Email,
		Name:     ur.Name,
	}
}

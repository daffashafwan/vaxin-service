package requests

import "github.com/daffashafwan/vaxin-service/business/users"

type UserVerify struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Status   string `json:"status"`
	Password string `json:"password"`
}

func (ur *UserVerify) ToDomain() users.Domain {
	return users.Domain{
		Username: ur.Username,
		Password: ur.Password,
		Token:    ur.Token,
		Status:   ur.Status,
		Email:    ur.Email,
		Name:     ur.Name,
	}
}

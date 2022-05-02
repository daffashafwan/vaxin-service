package admins

import (
	"context"
	"time"
)

type Domain struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	JWTToken  string    `json:"jwtToken"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Usecase interface {
	Login(ctx context.Context, domain Domain) (Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	Login(ctx context.Context, username string, password string) (Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
}

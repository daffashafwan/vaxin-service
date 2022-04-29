package users

import (
	"context"
	"time"
)

type Domain struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Token     string    `json:"token"`
	Status    string    `json:"status"`
	Email     string    `json:"email"`
	JWTToken  string    `json:"jwtToken"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Usecase interface {
	Login(ctx context.Context, domain Domain) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	// Update(ctx context.Context, domain Domain, id int) (Domain, error)
	// ResetPassword(ctx context.Context, password string, retypePassword string, id int) (Domain, error)
	// GetAll(ctx context.Context) ([]Domain, error)
	// GetById(ctx context.Context, id int) (Domain, error)
	// GetByToken(ctx context.Context, token string) (Domain, error)
	// ForgotPassword(ctx context.Context, email string) (Domain, error)
	// Delete(ctx context.Context, id int) error
	// Verify(ctx context.Context, domain Domain, id int) (Domain, error)
}

type Repository interface {
	Login(ctx context.Context, username string, password string) (Domain, error)
	Create(ctx context.Context, domain *Domain) (Domain, error)
	Update(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	GetByToken(ctx context.Context, token string) (Domain, error)
	GetByUsername(ctx context.Context, username string) (Domain, error)
	GetByEmail(ctx context.Context, email string) (Domain, error)
	Delete(ctx context.Context, id int) error
}

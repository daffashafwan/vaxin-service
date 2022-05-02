package vaccinations

import (
	"context"
	"time"
)

type Domain struct {
	Id        int         `json:"id"`
	SpecialId string   `json:"special_id"`
	EventId   int         `json:"event_id"`
	Event     interface{} `json:"event"`
	UserId    int         `json:"user_id"`
	User      interface{} `json:"user"`
	Queue     int         `json:"queue"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, domain Domain, id int) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	GetByUserId(ctx context.Context, id int) ([]Domain, error)
	GetByEventId(ctx context.Context, id int) ([]Domain, error)
	Delete(ctx context.Context, id int) error
}

type Repository interface {
	Create(ctx context.Context, domain *Domain) (Domain, error)
	Update(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	GetBySpecialId(ctx context.Context, id string) (Domain, error)
	GetByUserId(ctx context.Context, id int) ([]Domain, error)
	GetByEventId(ctx context.Context, id int) ([]Domain, error)
	Delete(ctx context.Context, id int) error
}

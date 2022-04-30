package events

import (
	"context"
	"time"
)

type Domain struct {
	Id         int         `json:"id"`
	Name       string      `json:"name"`
	StartDate  string      `json:"start_date"`
	EndDate    string      `json:"end_date"`
	Quota      int         `json:"quota"`
	VaccineId  int         `json:"vaccine_id"`
	Vaccine    interface{} `json:"vaccine"`
	FacilityId int         `json:"facility_id"`
	Facility   interface{} `json:"facility"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Usecase interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, domain Domain, id int) (Domain, error)
	UpdateQuota(ctx context.Context, domain Domain, id int) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	GetByFacilityId(ctx context.Context, id int) ([]Domain, error)
	GetByVaccineId(ctx context.Context, id int) ([]Domain, error)
	Delete(ctx context.Context, id int) error
}

type Repository interface {
	Create(ctx context.Context, domain *Domain) (Domain, error)
	Update(ctx context.Context, domain Domain) (Domain, error)
	UpdateQuota(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	GetByFacilityId(ctx context.Context, id int) ([]Domain, error)
	GetByVaccineId(ctx context.Context, id int) ([]Domain, error)
	Delete(ctx context.Context, id int) error
}

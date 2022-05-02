package events

import (
	"context"
	"errors"
	"time"

	"github.com/daffashafwan/vaxin-service/app/middlewares"
)

type EventUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
	ConfigJWT		middlewares.ConfigJWT
}

func NewEventUsecase(repo Repository, timeout time.Duration, configJWT middlewares.ConfigJWT) Usecase {
	return &EventUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}


func (ic *EventUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	var err error
	events, err := ic.Repo.Create(ctx, &domain)
	if err != nil {
		return Domain{}, err
	}

	return events, nil
}

func (ic *EventUsecase) Delete(ctx context.Context, id int) ( error) {
	err := ic.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (ic *EventUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	events, err := ic.Repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return events, nil
}

func (ic *EventUsecase) GetById(ctx context.Context, id int) (Domain, error) {
	events, err := ic.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if events.Id == 0 {
		return Domain{}, errors.New("ID NOT FOUND")
	}
	return events, nil
}

func (ic *EventUsecase) GetByVaccineId(ctx context.Context, id int) ([]Domain, error) {
	events, err := ic.Repo.GetByVaccineId(ctx, id)
	if err != nil {
		return []Domain{}, err
	}
	return events, nil
}

func (ic *EventUsecase) GetByFacilityId(ctx context.Context, id int) ([]Domain, error) {
	events, err := ic.Repo.GetByFacilityId(ctx, id)
	if err != nil {
		return []Domain{}, err
	}
	return events, nil
}

func (ic *EventUsecase) Update(ctx context.Context, domain Domain, id int) (Domain, error) {
	domain.Id = id
	events, err := ic.Repo.Update(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return events, nil
}

func (ic *EventUsecase) UpdateQuota(ctx context.Context, domain Domain, id int) (Domain, error) {
	event, errs := ic.Repo.GetById(ctx, id)
	if errs != nil {
		return Domain{}, errs
	}
	event.Id = id
	event.Quota = domain.Quota
	events, err := ic.Repo.UpdateQuota(ctx, event)
	if err != nil {
		return Domain{}, err
	}
	events.Facility = event.Facility
	events.Vaccine = event.Vaccine

	return events, nil
}

func (ic *EventUsecase) UpdateQueue(ctx context.Context, domain Domain, id int) (Domain, error) {
	event, errs := ic.Repo.GetById(ctx, id)
	if errs != nil {
		return Domain{}, errs
	}
	event.Id = id
	event.Queue = domain.Queue
	events, err := ic.Repo.UpdateQueue(ctx, event)
	if err != nil {
		return Domain{}, err
	}
	events.Facility = event.Facility
	events.Vaccine = event.Vaccine

	return events, nil
}
package facilities

import (
	"context"
	"errors"
	"time"

	"github.com/daffashafwan/vaxin-service/app/middlewares"
)

type FacilityUsecase struct {
	Repo           	Repository
	contextTimeout 	time.Duration
	ConfigJWT		middlewares.ConfigJWT
}

func NewFacilityUsecase(repo Repository, timeout time.Duration, configJWT middlewares.ConfigJWT) Usecase {
	return &FacilityUsecase{
		ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (pc *FacilityUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	var err error
	facility, err := pc.Repo.Create(ctx, &domain)
	if err != nil {
		return Domain{}, err
	}

	return facility, nil
}

func (uc *FacilityUsecase) Delete(ctx context.Context, id int) ( error) {
	err := uc.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *FacilityUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	facility, err := uc.Repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return facility, nil
}

func (uc *FacilityUsecase) GetById(ctx context.Context, id int) (Domain, error) {
	facility, err := uc.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if facility.Id == 0 {
		return Domain{}, errors.New("ID NOT FOUND")
	}
	return facility, nil
}

func (uc *FacilityUsecase) Update(ctx context.Context, domain Domain, id int) (Domain, error) {
	domain.Id = id
	facility, err := uc.Repo.Update(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return facility, nil
}
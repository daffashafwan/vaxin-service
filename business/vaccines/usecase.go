package vaccines

import (
	"context"
	"errors"
	"time"

	"github.com/daffashafwan/vaxin-service/app/middlewares"
)

type VaccineUsecase struct {
	Repo           	Repository
	contextTimeout 	time.Duration
	ConfigJWT		middlewares.ConfigJWT
}

func NewVaccineUsecase(repo Repository, timeout time.Duration, configJWT middlewares.ConfigJWT) Usecase {
	return &VaccineUsecase{
		ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (pc *VaccineUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	var err error
	vaccine, err := pc.Repo.Create(ctx, &domain)
	if err != nil {
		return Domain{}, err
	}

	return vaccine, nil
}

func (uc *VaccineUsecase) Delete(ctx context.Context, id int) ( error) {
	err := uc.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *VaccineUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	vaccine, err := uc.Repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return vaccine, nil
}

func (uc *VaccineUsecase) GetById(ctx context.Context, id int) (Domain, error) {
	vaccine, err := uc.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if vaccine.Id == 0 {
		return Domain{}, errors.New("ID NOT FOUND")
	}
	return vaccine, nil
}

func (uc *VaccineUsecase) Update(ctx context.Context, domain Domain, id int) (Domain, error) {
	domain.Id = id
	vaccine, err := uc.Repo.Update(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return vaccine, nil
}

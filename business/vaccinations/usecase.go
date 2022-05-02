package vaccinations

import (
	"context"
	"errors"
	"strings"
	"time"
	"github.com/daffashafwan/vaxin-service/app/middlewares"
	"github.com/daffashafwan/vaxin-service/helpers/randomizer"
)

type VaccinationUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
	ConfigJWT		middlewares.ConfigJWT
}

func NewVaccinationUsecase(repo Repository, timeout time.Duration, configJWT middlewares.ConfigJWT) Usecase {
	return &VaccinationUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}


func (ic *VaccinationUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	var err error
	var flag bool = false
	for !flag {
		domain.SpecialId = strings.ToUpper(randomizer.Randomize(20))[:10]
		data,_ := ic.Repo.GetBySpecialId(ctx, domain.SpecialId)
		if(data.Id == 0){
			break
		}
	}
	vaccinations, err := ic.Repo.Create(ctx, &domain)
	if err != nil {
		return Domain{}, err
	}

	return vaccinations, nil
}

func (ic *VaccinationUsecase) Delete(ctx context.Context, id int) ( error) {
	err := ic.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (ic *VaccinationUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	vaccinations, err := ic.Repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return vaccinations, nil
}

func (ic *VaccinationUsecase) GetById(ctx context.Context, id int) (Domain, error) {
	vaccinations, err := ic.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if vaccinations.Id == 0 {
		return Domain{}, errors.New("ID NOT FOUND")
	}
	return vaccinations, nil
}

func (ic *VaccinationUsecase) GetByEventId(ctx context.Context, id int) ([]Domain, error) {
	vaccinations, err := ic.Repo.GetByEventId(ctx, id)
	if err != nil {
		return []Domain{}, err
	}
	return vaccinations, nil
}

func (ic *VaccinationUsecase) GetByUserId(ctx context.Context, id int) ([]Domain, error) {
	vaccinations, err := ic.Repo.GetByUserId(ctx, id)
	if err != nil {
		return []Domain{}, err
	}
	return vaccinations, nil
}

func (ic *VaccinationUsecase) Update(ctx context.Context, domain Domain, id int) (Domain, error) {
	domain.Id = id
	vaccinations, err := ic.Repo.Update(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return vaccinations, nil
}
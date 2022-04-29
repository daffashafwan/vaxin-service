package users

import (
	"context"
	"time"

	"github.com/daffashafwan/vaxin-service/app/middlewares"
	errors "github.com/daffashafwan/vaxin-service/business"
	"github.com/daffashafwan/vaxin-service/helpers/email"
	"github.com/daffashafwan/vaxin-service/helpers/encrypt"
	"github.com/daffashafwan/vaxin-service/helpers/randomizer"
)

type UserUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
	ConfigJWT		middlewares.ConfigJWT
}

func NewUserUsecase(repo Repository, timeout time.Duration, configJWT middlewares.ConfigJWT) Usecase {
	return &UserUsecase{
		ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *UserUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Username == "" {
		return Domain{}, errors.ErrUsernameRequired
	}

	if domain.Password == "" {
		return Domain{}, errors.ErrPasswordRequired
	}
	var err error

	if err != nil {
		return Domain{}, err
	}
	var flag bool = false
	var hashed string
	hashed,_ = encrypt.Encrypt(domain.Password)
	domain.Password = hashed
	for !flag {
		domain.Token = randomizer.Randomize(20)
		user,_ := uc.Repo.GetByToken(ctx, domain.Token)
		if(user.Id == 0){
			break
		}
	}
	usern, _ := uc.Repo.GetByUsername(ctx, domain.Username)
	if usern.Id != 0 {
		return Domain{}, errors.ErrUsernameAlreadyExisted
	}
	usere, _ := uc.Repo.GetByEmail(ctx, domain.Email)
	if usere.Id != 0 {
		return Domain{}, errors.ErrEmailHasBeenRegister
	}
	user, err := uc.Repo.Create(ctx, &domain)
	if err != nil {
		return Domain{}, err
	}
	email.SendEmail(ctx, domain.Email, "Verifikasi Email", "<a href=`http://localhost:1333/users/verify/"+domain.Token+"`>Link Verifikasi</a>")

	return user, nil
}

func (uc *UserUsecase) Login(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Username == "" {
		return Domain{}, errors.ErrUsernamePasswordNotFound
	}

	if domain.Password == "" {
		return Domain{}, errors.ErrUsernamePasswordNotFound
	}

	user, err := uc.Repo.Login(ctx, domain.Username, domain.Password)

	if err != nil {
		return Domain{}, err
	}
	user.JWTToken, err = uc.ConfigJWT.GenerateTokenJWT(user.Id, 0)

	if err != nil {
		return Domain{}, err
	}

	return user,  nil
}

func (uc *UserUsecase) Verify(ctx context.Context, domain Domain, id int) (Domain, error) {
	domain.Status = "1"
	user, err := uc.Repo.Update(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc *UserUsecase) GetByToken(ctx context.Context, token string) (Domain, error) {
	user, err := uc.Repo.GetByToken(ctx, token)
	if err != nil {
		return Domain{}, err
	}
	if user.Id == 0 {
		return Domain{}, errors.ErrIDNotFound
	}
	return user, nil
}
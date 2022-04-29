package users

import (
	"net/http"
	//"strconv"
	"github.com/daffashafwan/vaxin-service/business/users"
	"github.com/daffashafwan/vaxin-service/deliveries/users/requests"
	"github.com/daffashafwan/vaxin-service/deliveries/users/responses"
	"github.com/daffashafwan/vaxin-service/helpers/response"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase users.Usecase
}

func NewUserController(userUseCase users.Usecase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func (userController UserController) Login(c echo.Context) error {

	userLogin := requests.UserLogin{}
	c.Bind(&userLogin)
	ctx := c.Request().Context()
	user, err := userController.UserUseCase.Login(ctx, userLogin.ToDomain())

	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, responses.FromDomain(user))
}

func (userController UserController) Register(c echo.Context) error {

	userRegister := requests.UserRegister{}
	c.Bind(&userRegister)
	ctx := c.Request().Context()
	user, err := userController.UserUseCase.Create(ctx, userRegister.ToDomain())

	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, responses.FromDomain(user))
}

func (userController UserController) Verify(c echo.Context) error {
	ctxNative := c.Request().Context()
	token := c.Param("token")
	data, err := userController.UserUseCase.GetByToken(ctxNative, token)
	if data.Status == "1" {
		return response.SuccessResponse(c, http.StatusOK, "Anda Sudah Pernah Melakukan Verifikasi")
	}
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	userVerif := requests.UserVerify{
		Id:       data.Id,
		Name:     data.Name,
		Username: data.Username,
		Token:    data.Token,
		Status:   "1",
		Password: data.Password,
		Email:    data.Email,
	}
	c.Bind(&userVerif)
	ctx := c.Request().Context()
	data, errs := userController.UserUseCase.Verify(ctx, userVerif.ToDomain(), data.Id)
	if errs != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, errs.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, responses.FromDomain(data))
}

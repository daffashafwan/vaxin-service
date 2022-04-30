package vaccines

import (
	"net/http"
	"strconv"

 	"github.com/daffashafwan/vaxin-service/business/vaccines"
	"github.com/daffashafwan/vaxin-service/deliveries/vaccines/requests"
	"github.com/daffashafwan/vaxin-service/deliveries/vaccines/responses"
	"github.com/daffashafwan/vaxin-service/helpers/response"
	"github.com/labstack/echo/v4"
)

type VaccineController struct {
	VaccineUsecase vaccines.Usecase
}

func NewVaccineController(vaccineUseCase vaccines.Usecase) *VaccineController {
	return &VaccineController{
		VaccineUsecase: vaccineUseCase,
	}
}

func (VaccineController VaccineController) Create(c echo.Context) error {

	vaccineCreate := requests.Vaccine{}
	c.Bind(&vaccineCreate)
	ctx := c.Request().Context()
	vaccine, error := VaccineController.VaccineUsecase.Create(ctx, vaccineCreate.ToDomain())

	if error != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, error)
	}

	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(vaccine))
}

func (VaccineController VaccineController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("vid")
	convId, _ := strconv.Atoi(id)
	vacReq := requests.Vaccine{}
	var err = c.Bind(&vacReq)
	if err != nil {
		return err
	}
	data, err := VaccineController.VaccineUsecase.Update(ctx, vacReq.ToDomain(), convId)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(data))
}

func (VaccineController VaccineController) GetAll(c echo.Context) error {
	ctxNative := c.Request().Context()
	data, err := VaccineController.VaccineUsecase.GetAll(ctxNative)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromListDomain(data))
}

func (VaccineController VaccineController) GetById(c echo.Context) error {
	ctxNative := c.Request().Context()
	id := c.Param("vid")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, errConvInt.Error())
	}
	data, err := VaccineController.VaccineUsecase.GetById(ctxNative, convInt)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromDomain(data))
}

func (VaccineController VaccineController) Delete(c echo.Context) error {
	id := c.Param("vid")
	convInt, _ := strconv.Atoi(id)
	ctx := c.Request().Context()
	var err = VaccineController.VaccineUsecase.Delete(ctx, convInt)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c,http.StatusOK, convInt)
}
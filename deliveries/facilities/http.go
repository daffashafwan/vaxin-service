package facilities

import (
	"net/http"
	"strconv"

 	"github.com/daffashafwan/vaxin-service/business/facilities"
	"github.com/daffashafwan/vaxin-service/deliveries/facilities/requests"
	"github.com/daffashafwan/vaxin-service/deliveries/facilities/responses"
	"github.com/daffashafwan/vaxin-service/helpers/response"
	"github.com/labstack/echo/v4"
)

type FacilityController struct {
	FacilityUsecase facilities.Usecase
}

func NewFacilityController(facilityUseCase facilities.Usecase) *FacilityController {
	return &FacilityController{
		FacilityUsecase: facilityUseCase,
	}
}

func (FacilityController FacilityController) Create(c echo.Context) error {

	facilityCreate := requests.Facility{}
	c.Bind(&facilityCreate)
	ctx := c.Request().Context()
	facility, error := FacilityController.FacilityUsecase.Create(ctx, facilityCreate.ToDomain())

	if error != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, error)
	}

	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(facility))
}

func (FacilityController FacilityController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("fid")
	convId, _ := strconv.Atoi(id)
	facReq := requests.Facility{}
	var err = c.Bind(&facReq)
	if err != nil {
		return err
	}
	data, err := FacilityController.FacilityUsecase.Update(ctx, facReq.ToDomain(), convId)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(data))
}

func (FacilityController FacilityController) GetAll(c echo.Context) error {
	ctxNative := c.Request().Context()
	data, err := FacilityController.FacilityUsecase.GetAll(ctxNative)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromListDomain(data))
}

func (FacilityController FacilityController) GetById(c echo.Context) error {
	ctxNative := c.Request().Context()
	id := c.Param("fid")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, errConvInt.Error())
	}
	data, err := FacilityController.FacilityUsecase.GetById(ctxNative, convInt)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromDomain(data))
}

func (FacilityController FacilityController) Delete(c echo.Context) error {
	id := c.Param("fid")
	convInt, _ := strconv.Atoi(id)
	ctx := c.Request().Context()
	var err = FacilityController.FacilityUsecase.Delete(ctx, convInt)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c,http.StatusOK, convInt)
}
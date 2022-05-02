package vaccinations

import (
	"net/http"
	"strconv"

	"github.com/daffashafwan/vaxin-service/business/events"
	"github.com/daffashafwan/vaxin-service/business/vaccinations"
	"github.com/daffashafwan/vaxin-service/deliveries/vaccinations/requests"
	"github.com/daffashafwan/vaxin-service/deliveries/vaccinations/responses"
	"github.com/daffashafwan/vaxin-service/helpers/response"
	"github.com/daffashafwan/vaxin-service/helpers/errors"
	"github.com/labstack/echo/v4"
)

type VaccinationController struct {
	VaccinationUsecase vaccinations.Usecase
	EventUsecase events.Usecase
}

func NewVaccinationController(vaccinationUsecase vaccinations.Usecase, eventUsecase events.Usecase) *VaccinationController {
	return &VaccinationController{
		VaccinationUsecase: vaccinationUsecase,
		EventUsecase: eventUsecase,
	}
}

func (vaccinationsController VaccinationController) Create(c echo.Context) error {
	itemCreate := requests.VaccinationRequest{}
	c.Bind(&itemCreate)
	ctx := c.Request().Context()
	vacDomain := itemCreate.ToDomain()

	vaccinationUser,err := vaccinationsController.VaccinationUsecase.GetByUserId(ctx, vacDomain.UserId)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	if len(vaccinationUser) > 0{
		return response.ErrorResponse(c, http.StatusBadRequest, errors.ErrDataAlreadyExisted)
	}
	eventCurrent,_ := vaccinationsController.EventUsecase.GetById(ctx, vacDomain.EventId)
	eventCurrent.Queue = eventCurrent.Queue + 1
	eventUpdatedQueue,_ := vaccinationsController.EventUsecase.UpdateQueue(ctx, eventCurrent ,vacDomain.EventId)
	
	eventUpdatedQueue.Quota = eventUpdatedQueue.Quota - 1 
	eventUpdatedQuota,_ := vaccinationsController.EventUsecase.UpdateQuota(ctx, eventUpdatedQueue ,vacDomain.EventId)
	
	vacDomain.Queue = eventUpdatedQuota.Queue
	item, errors := vaccinationsController.VaccinationUsecase.Create(ctx, vacDomain)
	if errors != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, errors)
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromDomain(item))
}

func (vaccinationsController VaccinationController) GetAll(c echo.Context) error {
	ctxNative := c.Request().Context()
	data, err := vaccinationsController.VaccinationUsecase.GetAll(ctxNative)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromListDomain(data))
}

func (vaccinationsController VaccinationController) GetById(c echo.Context) error {
	ctxNative := c.Request().Context()
	id := c.Param("vcid")
	convId, _ := strconv.Atoi(id)
	data, err := vaccinationsController.VaccinationUsecase.GetById(ctxNative, convId)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromDomain(data))
}

func (vaccinationsController VaccinationController) GetByEventId(c echo.Context) error {
	ctxNative := c.Request().Context()
	id := c.Param("eid")
	convId, _ := strconv.Atoi(id)
	data, err := vaccinationsController.VaccinationUsecase.GetByEventId(ctxNative, convId)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromListDomain(data))
}

func (vaccinationsController VaccinationController) GetByUserId(c echo.Context) error {
	ctxNative := c.Request().Context()
	id := c.Param("uid")
	convId, _ := strconv.Atoi(id)
	data, err := vaccinationsController.VaccinationUsecase.GetByUserId(ctxNative, convId)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromListDomain(data))
}

func (vaccinationsController VaccinationController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("vcid")
	convId, _ := strconv.Atoi(id)
	vaccinations, _ := vaccinationsController.VaccinationUsecase.GetById(ctx, convId)
	eventsRequest := requests.VaccinationRequest{}
	var err = c.Bind(&eventsRequest)
	if err != nil {
		return err
	}
	data, err := vaccinationsController.VaccinationUsecase.Update(ctx, eventsRequest.ToDomain(), vaccinations.Id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(data))
}

func (vaccinationsController VaccinationController) Delete(c echo.Context) error {
	id := c.Param("vcid")
	convInt, _ := strconv.Atoi(id)
	ctx := c.Request().Context()
	var err = vaccinationsController.VaccinationUsecase.Delete(ctx, convInt)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c,http.StatusOK, convInt)
}
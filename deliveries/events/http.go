package events

import (
	"net/http"
	"strconv"
	"github.com/daffashafwan/vaxin-service/business/events"
	"github.com/daffashafwan/vaxin-service/deliveries/events/requests"
	"github.com/daffashafwan/vaxin-service/deliveries/events/responses"
	"github.com/daffashafwan/vaxin-service/helpers/response"
	"github.com/labstack/echo/v4"
)

type EventController struct {
	EventUsecase events.Usecase
}

func NewEventController(eventsUsecase events.Usecase) *EventController {
	return &EventController{
		EventUsecase: eventsUsecase,
	}
}

func (eventsController EventController) Create(c echo.Context) error {
	itemCreate := requests.EventRequest{}
	c.Bind(&itemCreate)
	ctx := c.Request().Context()
	item, errors := eventsController.EventUsecase.Create(ctx, itemCreate.ToDomain())
	if errors != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, errors)
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromDomain(item))
}

func (eventsController EventController) GetAll(c echo.Context) error {
	ctxNative := c.Request().Context()
	data, err := eventsController.EventUsecase.GetAll(ctxNative)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromListDomain(data))
}

func (eventsController EventController) GetById(c echo.Context) error {
	ctxNative := c.Request().Context()
	id := c.Param("eid")
	convId, _ := strconv.Atoi(id)
	data, err := eventsController.EventUsecase.GetById(ctxNative, convId)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromDomain(data))
}

func (eventsController EventController) GetByVaccineId(c echo.Context) error {
	ctxNative := c.Request().Context()
	id := c.Param("vid")
	convId, _ := strconv.Atoi(id)
	data, err := eventsController.EventUsecase.GetByVaccineId(ctxNative, convId)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromListDomain(data))
}

func (eventsController EventController) GetByFacilityId(c echo.Context) error {
	ctxNative := c.Request().Context()
	id := c.Param("fid")
	convId, _ := strconv.Atoi(id)
	data, err := eventsController.EventUsecase.GetByFacilityId(ctxNative, convId)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, responses.FromListDomain(data))
}

func (eventsController EventController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("eid")
	convId, _ := strconv.Atoi(id)
	events, _ := eventsController.EventUsecase.GetById(ctx, convId)
	eventsRequest := requests.EventRequest{}
	var err = c.Bind(&eventsRequest)
	if err != nil {
		return err
	}
	data, err := eventsController.EventUsecase.Update(ctx, eventsRequest.ToDomain(), events.Id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(data))
}

func (eventsController EventController) UpdateStock(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("eid")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	eventsRequest := requests.ItemRequestQuota{}
	err = c.Bind(&eventsRequest)
	if err != nil {
		return err
	}
	data, err := eventsController.EventUsecase.UpdateQuota(ctx, eventsRequest.QuotaToDomain(), convId)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c,http.StatusOK, responses.FromDomain(data))
}

func (eventsController EventController) Delete(c echo.Context) error {
	id := c.Param("eid")
	convInt, _ := strconv.Atoi(id)
	ctx := c.Request().Context()
	var err = eventsController.EventUsecase.Delete(ctx, convInt)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return response.SuccessResponse(c,http.StatusOK, convInt)
}
package routes

import (
	"github.com/daffashafwan/vaxin-service/app/middlewares"
	"github.com/daffashafwan/vaxin-service/deliveries/admins"
	"github.com/daffashafwan/vaxin-service/deliveries/events"
	"github.com/daffashafwan/vaxin-service/deliveries/facilities"
	"github.com/daffashafwan/vaxin-service/deliveries/users"
	"github.com/daffashafwan/vaxin-service/deliveries/vaccines"
	"github.com/daffashafwan/vaxin-service/deliveries/vaccinations"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JwtConfig          middleware.JWTConfig
	AdminController    admins.AdminController
	UserController     users.UserController
	VaccineController  vaccines.VaccineController
	FacilityController facilities.FacilityController
	EventController    events.EventController
	VaccinationsController vaccinations.VaccinationController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	//USER
	e.POST("users/login", cl.UserController.Login)
	e.POST("users/register", cl.UserController.Register)
	e.GET("users/verify/:token", cl.UserController.Verify)


	//ADMIN
	e.GET("admin/:id", cl.AdminController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	e.POST("admin/login", cl.AdminController.Login)


	//VACCINE
	//getAll
	e.GET("/:id/vaccines", cl.VaccineController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("vaccines", cl.VaccineController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	
	//getById
	e.GET("/:id/vaccines/:vid", cl.VaccineController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("vaccines/:vid", cl.VaccineController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	
	//create
	e.POST("vaccines", cl.VaccineController.Create, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	//edit
	e.PUT("vaccines/:vid", cl.VaccineController.Update, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	
	//delete
	e.DELETE("vaccines/:vid", cl.VaccineController.Delete, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)


	//FACILITIES
	//getAll
	e.GET("/:id/facilities", cl.FacilityController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("facilities", cl.FacilityController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	
	//getById
	e.GET("/:id/facilities/:fid", cl.FacilityController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("facilities/:fid", cl.FacilityController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	
	//create
	e.POST("facilities", cl.FacilityController.Create, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	//edit
	e.PUT("facilities/:fid", cl.FacilityController.Update, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	//edit
	e.DELETE("facilities/:fid", cl.FacilityController.Delete, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)


	//EVENTS
	//getAll
	e.GET("/:id/events", cl.EventController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("events", cl.EventController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	//getById
	e.GET("/:id/events/:eid", cl.EventController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("events/:eid", cl.EventController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	//getByFacilityId
	e.GET("/:id/events/facility/:fid", cl.EventController.GetByFacilityId, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("events/facility/:fid", cl.EventController.GetByFacilityId, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	//getByVaccineId
	e.GET("/:id/events/vaccine/:vid", cl.EventController.GetByVaccineId, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("events/vaccine/:vid", cl.EventController.GetByVaccineId, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	
	//create
	e.POST("events", cl.EventController.Create, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	//edit
	e.PUT("events/:eid", cl.EventController.Update, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	//edit quota
	e.PUT("events/:eid/quota", cl.EventController.UpdateStock, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	
	//delete
	e.DELETE("events/:eid", cl.EventController.Delete, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)


	//VACCINATONS
	//getAll
	e.GET("/:id/vaccinations", cl.VaccinationsController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("vaccinations", cl.VaccinationsController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	//getAll
	e.GET("/:id/vaccinations", cl.VaccinationsController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("vaccinations", cl.VaccinationsController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	//getById
	e.GET("/:id/vaccinations/:vcid", cl.VaccinationsController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("vaccinations/:vcid", cl.VaccinationsController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	//getByUserId
	e.GET("/:id/vaccinations/user/:uid", cl.VaccinationsController.GetByUserId, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("vaccinations/user/:uid", cl.VaccinationsController.GetByUserId, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	//getByEventId
	e.GET("/:id/vaccinations/event/:eid", cl.VaccinationsController.GetByEventId, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("vaccinations/event/:eid", cl.VaccinationsController.GetByEventId, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	//create
	e.POST("/:id/vaccinations", cl.VaccinationsController.Create, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
}

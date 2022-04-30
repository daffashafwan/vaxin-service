package routes

import (
	"github.com/daffashafwan/vaxin-service/app/middlewares"
	"github.com/daffashafwan/vaxin-service/deliveries/users"
	"github.com/daffashafwan/vaxin-service/deliveries/admins"
	"github.com/daffashafwan/vaxin-service/deliveries/vaccines"
	"github.com/daffashafwan/vaxin-service/deliveries/facilities"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JwtConfig      middleware.JWTConfig
	AdminController admins.AdminController
	UserController users.UserController
	VaccineController vaccines.VaccineController
	FacilityController facilities.FacilityController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	//user
	e.POST("users/login", cl.UserController.Login)
	e.POST("users/register", cl.UserController.Register)
	e.GET("users/verify/:token", cl.UserController.Verify)

	//admin
	e.GET("admin/:id", cl.AdminController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	e.POST("admin/login", cl.AdminController.Login)

	//vaccine
	e.GET("/:id/vaccines", cl.VaccineController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("vaccines", cl.VaccineController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	
	e.GET("/:id/vaccines/:vid", cl.VaccineController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("vaccines/:vid", cl.VaccineController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	
	e.POST("vaccines", cl.VaccineController.Create, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	
	e.PUT("vaccines/:vid", cl.VaccineController.Update, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	e.DELETE("vaccines/:vid", cl.VaccineController.Delete, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)

	//facilities
	e.GET("/:id/facilities", cl.FacilityController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("facilities", cl.FacilityController.GetAll, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	
	e.GET("/:id/facilities/:fid", cl.FacilityController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsUserId)
	e.GET("facilities/:fid", cl.FacilityController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	
	e.POST("facilities", cl.FacilityController.Create, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	
	e.PUT("facilities/:fid", cl.FacilityController.Update, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	e.DELETE("facilities/:fid", cl.FacilityController.Delete, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
}
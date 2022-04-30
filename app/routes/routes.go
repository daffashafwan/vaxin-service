package routes

import (
	"github.com/daffashafwan/vaxin-service/app/middlewares"
	"github.com/daffashafwan/vaxin-service/deliveries/users"
	"github.com/daffashafwan/vaxin-service/deliveries/admins"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JwtConfig      middleware.JWTConfig
	AdminController admins.AdminController
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.POST("users/login", cl.UserController.Login)
	e.POST("users/register", cl.UserController.Register)
	e.GET("users/verify/:token", cl.UserController.Verify)

	e.GET("admin/:id", cl.AdminController.GetById, middleware.JWTWithConfig(cl.JwtConfig), middlewares.IsAdmin)
	e.POST("admin/login", cl.AdminController.Login)
}
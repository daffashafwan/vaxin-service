package main

import (
	_middleware "github.com/daffashafwan/vaxin-service/app/middlewares"
	_mysqlDriver "github.com/daffashafwan/vaxin-service/config"
	"time"

	_userUsecase "github.com/daffashafwan/vaxin-service/business/users"
	_userController "github.com/daffashafwan/vaxin-service/deliveries/users"
	_userRepository "github.com/daffashafwan/vaxin-service/repository/user"
	_userdb "github.com/daffashafwan/vaxin-service/repository/user"

	_adminUsecase "github.com/daffashafwan/vaxin-service/business/admins"
	_adminController "github.com/daffashafwan/vaxin-service/deliveries/admins"
	_adminRepository "github.com/daffashafwan/vaxin-service/repository/admin"
	_admindb "github.com/daffashafwan/vaxin-service/repository/admin"

	_vaccineUsecase "github.com/daffashafwan/vaxin-service/business/vaccines"
	_vaccineController "github.com/daffashafwan/vaxin-service/deliveries/vaccines"
	_vaccineRepository "github.com/daffashafwan/vaxin-service/repository/vaccine"
	_vaccinedb "github.com/daffashafwan/vaxin-service/repository/vaccine"

	_facilityUsecase "github.com/daffashafwan/vaxin-service/business/facilities"
	_facilityController "github.com/daffashafwan/vaxin-service/deliveries/facilities"
	_facilityRepository "github.com/daffashafwan/vaxin-service/repository/facility"
	_facilitydb "github.com/daffashafwan/vaxin-service/repository/facility"

	_eventUsecase "github.com/daffashafwan/vaxin-service/business/events"
	_eventController "github.com/daffashafwan/vaxin-service/deliveries/events"
	_eventRepository "github.com/daffashafwan/vaxin-service/repository/event"
	_eventdb "github.com/daffashafwan/vaxin-service/repository/event"

	"github.com/daffashafwan/vaxin-service/app/routes"

	"log"
	//"fmt"
	//"os"
	//"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func DbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_userdb.User{},
		&_admindb.Admin{},
		&_vaccinedb.Vaccine{},
		&_facilitydb.Facility{},
		&_eventdb.Event{})
}

func main() {
	// init koneksi databse
	configDB := _mysqlDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	Conn := configDB.InitialDB()
	DbMigrate(Conn)

	e := echo.New()
	//feAddress := viper.GetString(`frontend.address`)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://pointcuan-fe.vercel.app", "http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowMethods, echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowHeaders, echo.HeaderAccessControlRequestHeaders, echo.HeaderAccessControlAllowCredentials},
	}))
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	userRepository := _userRepository.CreateUserRepo(Conn)
	userUseCase := _userUsecase.NewUserUsecase(userRepository, timeoutContext, configJWT)
	userController := _userController.NewUserController(userUseCase)

	adminRepository := _adminRepository.CreateAdminRepo(Conn)
	adminUseCase := _adminUsecase.NewUsecase(adminRepository, timeoutContext, configJWT)
	adminController := _adminController.NewAdminController(adminUseCase)

	vaccineRepository := _vaccineRepository.CreateVaccineRepo(Conn)
	vaccineUseCase := _vaccineUsecase.NewVaccineUsecase(vaccineRepository, timeoutContext, configJWT)
	vaccineController := _vaccineController.NewVaccineController(vaccineUseCase)

	facilityRepository := _facilityRepository.CreateFacilityRepo(Conn)
	facilityUseCase := _facilityUsecase.NewFacilityUsecase(facilityRepository, timeoutContext, configJWT)
	facilityController := _facilityController.NewFacilityController(facilityUseCase)

	eventRepository := _eventRepository.CreateEventRepo(Conn)
	eventUseCase := _eventUsecase.NewEventUsecase(eventRepository, timeoutContext, configJWT)
	eventController := _eventController.NewEventController(eventUseCase)

	routesInit := routes.ControllerList{
		JwtConfig:          configJWT.Init(),
		UserController:     *userController,
		AdminController:    *adminController,
		VaccineController:  *vaccineController,
		FacilityController: *facilityController,
		EventController:    *eventController,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}

package main

import(
	"time"
	_middleware "github.com/daffashafwan/vaxin-service/app/middlewares"
	_userUsecase "github.com/daffashafwan/vaxin-service/business/users"
	_mysqlDriver "github.com/daffashafwan/vaxin-service/config"
	_userController "github.com/daffashafwan/vaxin-service/deliveries/users"
	_userdb "github.com/daffashafwan/vaxin-service/repository/user"
	_userRepository "github.com/daffashafwan/vaxin-service/repository/user"

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
		&_userdb.User{})
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

	routesInit := routes.ControllerList{
		JwtConfig:             configJWT.Init(),
		UserController:        *userController,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}




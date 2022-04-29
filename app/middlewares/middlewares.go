package middlewares

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/daffashafwan/vaxin-service/helpers/response"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JWTCustomClaims struct {
	Id      int `json:"id"`
	IsAdmin int `json:"is_admin"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JWTCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
		SuccessHandler: func(context echo.Context) {
			//userId := context.Get("userId")
		},
		ErrorHandlerWithContext: func(e error, c echo.Context) error {
			return response.ErrorResponse(c, http.StatusUnauthorized, errors.New("Invalid Token Credential"))
		},
	}
}

func (jwtConf *ConfigJWT) GenerateTokenJWT(id int, isAdmin int) (string, error) {
	claims := JWTCustomClaims{
		id,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(jwtConf.SecretJWT))

	return token, nil
}

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JWTCustomClaims)
		isAdmin := claims.IsAdmin
		if isAdmin == 0 {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

func IsUserId(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		userId := c.Param("id")
		convUserId, err := strconv.Atoi(userId)
		if err != nil {
			return echo.ErrBadRequest
		}
		claims := user.Claims.(*JWTCustomClaims)
		claimUserId := claims.Id
		if claimUserId != convUserId {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

package shared

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type ServerConfig struct {
	Port       int8
	Validator  bool
	Swagger    bool
	CORSConfig *middleware.CORSConfig
}

func Server(config *ServerConfig) *echo.Echo {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Secure())
	e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	if config.CORSConfig != nil {
		e.Use(middleware.CORSWithConfig(*config.CORSConfig))
	}

	if config.Validator {
		e.Validator = NewValidator()
	}

	if config.Swagger {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	return e
}

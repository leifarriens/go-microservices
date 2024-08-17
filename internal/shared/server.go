package shared

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type ServerConfig struct {
	Port         int8
	Validator    bool
	KeyAuth      bool
	Swagger      bool
	AllowOrigins []string
}

func Server(config *ServerConfig) *echo.Echo {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Secure())
	e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	if len(config.AllowOrigins) > 0 {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: config.AllowOrigins,
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}))
	}

	if config.Validator {
		e.Validator = NewValidator()
	}

	if config.KeyAuth {
		e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
			KeyLookup:  "header:" + echo.HeaderAuthorization,
			AuthScheme: "Bearer",
			Validator: func(key string, c echo.Context) (bool, error) {
				return key == "valid-key", nil
			},
		}))
	}

	if config.Swagger {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	return e
}

package shared

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ServerConfig struct {
	Port    int8
	KeyAuth bool
}

func Server(config *ServerConfig) *echo.Echo {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Secure())
	e.Use(middleware.RequestID())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	if config.KeyAuth {
		e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
			KeyLookup:  "header:" + echo.HeaderAuthorization,
			AuthScheme: "Bearer",
			Validator: func(key string, c echo.Context) (bool, error) {
				return key == "valid-key", nil
			},
		}))
	}

	return e
}

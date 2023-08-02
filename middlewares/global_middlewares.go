package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GlobalMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.RemoveTrailingSlash())
}

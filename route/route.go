package route

import (
	"git.onespace.co.th/project-v/project-v-go-backend/controller"
	"github.com/labstack/echo/v4"
	guard "github.com/labstack/echo/v4/middleware"
)

func RegisterApiEp(e *echo.Echo) {
	e.Use(guard.CORSWithConfig(guard.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	// API Endpoint
	apiGroup := e.Group("/api")
	apiGroup.GET("/version", controller.GetAppVersionEp)
}

package middleware

import (
	"github.com/labstack/echo/v4/middleware"
)

const AuthorizedContext = "auth"

type Config struct {
	Skipper middleware.Skipper
}

var (
	DefaultConfig = Config{
		Skipper: middleware.DefaultSkipper,
	}
)

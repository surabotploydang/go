package controller

import (
	"git.onespace.co.th/project-v/project-v-go-backend/constant"
	"git.onespace.co.th/project-v/project-v-go-backend/setting"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAppVersionEp(c echo.Context) error {
	return c.JSON(http.StatusOK, constant.Json{"msg": setting.AppVersion})
}

package main

import (
	"fmt"

	"git.onespace.co.th/project-v/project-v-go-backend/route"
	"git.onespace.co.th/project-v/project-v-go-backend/service"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := service.InitServices(); err != nil {
		fmt.Println(err)
	}
	e := echo.New()
	route.RegisterApiEp(e)
	e.Logger.Fatal(e.Start(":5000"))
}

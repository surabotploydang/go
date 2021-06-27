package main

import (
	"encoding/json"
	"fmt"
	"git.onespace.co.th/golib/requests/requests"
	"github.com/labstack/echo/v4"
)

func main() {
	token := "xxxxx"
	param := requests.Params{
		URL:  "https://testoneid.inet.co.th/api/account",
		BODY: nil,
		HEADERS: map[string]string{
			echo.HeaderContentType:   "application/json",
			echo.HeaderAuthorization: "Bearer " + token,
		},
		TIMEOUT: 5,
	}
	var res requests.Response
	if err := requests.Call().Get(param, &res).Error(); err != nil {
		fmt.Println(err)
	}
	var resp map[string]interface{}
	if err := json.Unmarshal(res.Result, &resp); err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}

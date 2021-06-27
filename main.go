package main

import (
	"encoding/json"
	"fmt"
	"git.onespace.co.th/golib/requests/requests"
	"github.com/labstack/echo/v4"
)

func main() {
	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6ImVmYjYxZmY5MTlhOTg2YzdjM2IyMzE2ZmJkMTY0NjBmNDc3ZWZjMWZhZTU1N2Q1YjUxYmQ3ZDJiZTBlY2UyNjk5OGNlNTJkMDRjMjEzYjg5In0.eyJhdWQiOiI1MTQiLCJqdGkiOiJlZmI2MWZmOTE5YTk4NmM3YzNiMjMxNmZiZDE2NDYwZjQ3N2VmYzFmYWU1NTdkNWI1MWJkN2QyYmUwZWNlMjY5OThjZTUyZDA0YzIxM2I4OSIsImlhdCI6MTYyNDc3Mjg4MywibmJmIjoxNjI0NzcyODgzLCJleHAiOjE2MjUwMzIwODMsInN1YiI6IjgxMTgwNDc4OTI0OCIsInNjb3BlcyI6W119.pOzfFLLmJM__qzscx9nA__8YtNXfsrP5IVjQokkoyFt-ldnrjgq2K3hx_9Roqwg9NslWMUEBCAiSYeFVfoqniQOKo1uFjhpZ5oVaq_QB-8FyOiJ2TuuDUKgXsh7zgXZUSGGYWSiZ-dkyJRaoGLO-UGIqT4O-1pQLtjucg0rCtzmxCHjX6odoS-nN9aL3_mRyb0K1VLlOXCQMojekLM2PA7-8UhpukGOA7Gq1t1B6IQnv7_L_4FYWzfTDmarZl9HC5T0O02QUf7WIUM4kJPMKhTTUBsFx2gxrNS2Wu-tCNhFxJug8JsFzuont63jwPJlUEn_FHKqeeOndlsJvnqLmfW7-d52MvysX1wNc2s2cgkY4x3lfeSs0wBesp3RfLc2DTfnrueteCRNdMUXQ7PQU7yd6_PwOycbQlId12Ii7iQaJK2IUbwedHv6J1SyfOLnrbzDlu7dajVVuaJseBDY9c50wVgHaoPsusUj6zJGbMdZyDHWKUFWGZhIj0e-HzNJWQ3p5f3te7fXyrPgBPOUe_bn2OlgQfoUvkS-g2_iehuvQS2QCkH_mSdJxAG2Cnr5yWJ06wM6Au5J0cUYfPIcY6XHrCcxkVooEaebqy1iWwkQb6UEtNPgcrcGG2ONb8ntT8vCdIehQHLwgg_OCq5ZYKpQt3x7NM6Gv6S4MEmRHaRE"
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

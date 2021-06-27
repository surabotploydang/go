package requests

import (
	"bytes"
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"time"
)

const (
	Timeout = 5
)

type Response struct {
	Code   int
	Status string
	Header http.Header
	Result []byte
}

type Params struct {
	URL     string
	HEADERS map[string]string
	BODY    io.Reader
	TIMEOUT int
}

type (
	Requests interface {
		Error() error
		Get(Params, *Response) Requests
		Post(Params, *Response) Requests
		Put(Params, *Response) Requests
		Delete(Params, *Response) Requests
	}
	requests struct {
		err error
	}
)

func Call() Requests {
	return &requests{}
}

func (r *requests) Error() error {
	if r.err != nil {
		return r.err
	}
	return nil
}

func (r *requests) Get(param Params, response *Response) Requests {
	err := request(http.MethodGet, param.URL, param.HEADERS, param.BODY, param.TIMEOUT, response)
	if err != nil {
		r.err = err
	}
	return r
}

func (r *requests) Post(param Params, response *Response) Requests {
	err := request(http.MethodPost, param.URL, param.HEADERS, param.BODY, param.TIMEOUT, response)
	if err != nil {
		r.err = err
	}
	return r
}

func (r *requests) Put(param Params, response *Response) Requests {
	err := request(http.MethodPut, param.URL, param.HEADERS, param.BODY, param.TIMEOUT, response)
	if err != nil {
		r.err = err
	}
	return r
}

func (r *requests) Delete(param Params, response *Response) Requests {
	err := request(http.MethodDelete, param.URL, param.HEADERS, param.BODY, param.TIMEOUT, response)
	if err != nil {
		r.err = err
	}
	return r
}

func request(method string, url string, headers map[string]string, body io.Reader, timeout int, response *Response) (err error) {
	if timeout == 0 {
		timeout = Timeout
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(resp.Body)
	response.Code = resp.StatusCode
	response.Status = resp.Status
	response.Result = buf.Bytes()
	response.Header = resp.Header
	_ = resp.Body.Close()
	return nil
}

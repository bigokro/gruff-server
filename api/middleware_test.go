package api

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine"
	"github.com/labstack/echo/test"
	"log"
	"net/url"
	"strings"
)

const (
	Version         = "1.0"
	UserAgent       = "User-Agent"
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json"
	ApplicationForm = "application/x-www-form-urlencoded"
)

var Token = map[string]string{"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"}

type HTTPResponse *test.ResponseRecorder
type HTTPRequest engine.Request
type H map[string]string
type D map[string]interface{}

type RequestConfig struct {
	Method  string
	Path    string
	Body    string
	Headers H
	Cookies H
	Debug   bool
}

func New(Authorization map[string]string) *RequestConfig {
	return &RequestConfig{
		Headers: Authorization,
	}
}

func (rc *RequestConfig) SetDebug(enable bool) *RequestConfig {
	rc.Debug = enable
	return rc
}

func (rc *RequestConfig) GET(path string) *RequestConfig {
	rc.Path = path
	rc.Method = "GET"
	return rc
}

func (rc *RequestConfig) POST(path string) *RequestConfig {
	rc.Path = path
	rc.Method = "POST"
	return rc
}

func (rc *RequestConfig) PUT(path string) *RequestConfig {
	rc.Path = path
	rc.Method = "PUT"
	return rc
}

func (rc *RequestConfig) DELETE(path string) *RequestConfig {
	rc.Path = path
	rc.Method = "DELETE"
	return rc
}

func (rc *RequestConfig) PATCH(path string) *RequestConfig {
	rc.Path = path
	rc.Method = "PATCH"
	return rc
}

func (rc *RequestConfig) HEAD(path string) *RequestConfig {
	rc.Path = path
	rc.Method = "HEAD"
	return rc
}

func (rc *RequestConfig) OPTIONS(path string) *RequestConfig {
	rc.Path = path
	rc.Method = "OPTIONS"
	return rc
}

func (rc *RequestConfig) SetHeader(headers H) *RequestConfig {
	if len(headers) > 0 {
		rc.Headers = headers
	}

	return rc
}

func (rc *RequestConfig) SetJSON(body D) *RequestConfig {
	if b, err := json.Marshal(body); err == nil {
		rc.Body = string(b)
	}

	return rc
}

func (rc *RequestConfig) SetForm(body H) *RequestConfig {
	f := make(url.Values)

	for k, v := range body {
		f.Set(k, v)
	}

	rc.Body = f.Encode()

	return rc
}

func (rc *RequestConfig) SetQuery(query H) *RequestConfig {
	f := make(url.Values)

	for k, v := range query {
		f.Set(k, v)
	}

	if strings.Contains(rc.Path, "?") {
		rc.Path = rc.Path + "&" + f.Encode()
	} else {
		rc.Path = rc.Path + "?" + f.Encode()
	}

	return rc
}

func (rc *RequestConfig) SetBody(item interface{}) *RequestConfig {
	b, _ := json.Marshal(item)
	body := string(b)
	if len(body) > 0 {
		rc.Body = body
	}

	return rc
}

func (rc *RequestConfig) SetCookie(cookies H) *RequestConfig {
	if len(cookies) > 0 {
		rc.Cookies = cookies
	}

	return rc
}

func (rc *RequestConfig) initTest() (engine.Request, *test.ResponseRecorder) {
	qs := ""
	if strings.Contains(rc.Path, "?") {
		ss := strings.Split(rc.Path, "?")
		qs = ss[1]
	}

	body := bytes.NewBufferString(rc.Body)

	rq := test.NewRequest(rc.Method, rc.Path, body)

	if len(qs) > 0 {
		rq.URL().QueryParam(qs)
	}

	if rc.Method == "POST" || rc.Method == "PUT" {
		if strings.HasPrefix(rc.Body, "{") {
			rq.Header().Add(ContentType, ApplicationJSON)
		} else {
			rq.Header().Add(ContentType, ApplicationForm)
		}
	}

	if len(rc.Headers) > 0 {
		for k, v := range rc.Headers {
			rq.Header().Add(k, v)
		}
	}

	if rc.Debug {
		log.Printf("Request Method: %s", rc.Method)
		log.Printf("Request Path: %s", rc.Path)
		log.Printf("Request Body: %s", rc.Body)
		log.Printf("Request Headers: %s", rc.Headers)
		log.Printf("Request Cookies: %s", rc.Cookies)
	}

	rec := test.NewResponseRecorder()

	return rq, rec
}

func (rc *RequestConfig) Run(e *echo.Echo) (HTTPResponse *test.ResponseRecorder, HTTPRequest engine.Request) {
	rq, rec := rc.initTest()
	e.ServeHTTP(rq, rec)
	return rec, rq
}

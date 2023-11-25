package test

import (
	"github.com/kinbiko/jsonassert"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func PostJSON(target string, json string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, target, strings.NewReader(json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return c, rec
}

func JSONEq(t *testing.T, jsonA string, jsonB string) {
	ja := jsonassert.New(t)
	ja.Assertf(jsonA, jsonB)
}

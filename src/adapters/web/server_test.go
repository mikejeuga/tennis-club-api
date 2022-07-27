package web_test

import (
	"github.com/adamluzsi/testcase"
	"github.com/mikejeuga/tennis-club-api/src/adapters/web"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	tC := testcase.NewT(t, nil)
	server := web.NewServer()

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	server.Handler.ServeHTTP(res, req)

	tC.Must.Equal(http.StatusOK, res.Code)
}

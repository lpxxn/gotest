package testutils

import (
	"github.com/gin-gonic/gin"
	"github.com/jarcoal/httpmock"
	"github.com/lpxxn/gotest/app2_thirdlib/utils"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

func GetRequst(url string, engine *gin.Engine, params ...func(r *http.Request)) (*httptest.ResponseRecorder, error) {
	return newRequest(http.MethodGet, url, engine, nil, params...)
}

func PostJsonRequst(url string, engine *gin.Engine, body string, params ...func(r *http.Request)) (*httptest.ResponseRecorder, error) {
	params = append(params, func(r *http.Request) {
		r.Header.Set("Content-Type", "application/json")
	})
	return newRequest(http.MethodPost, url, engine, strings.NewReader(body), params...)
}

func PostFormRequst(url string, engine *gin.Engine, body url.Values, params ...func(r *http.Request)) (*httptest.ResponseRecorder, error) {
	params = append(params, func(r *http.Request) {
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	})
	return newRequest(http.MethodPost, url, engine, strings.NewReader(body.Encode()), params...)
}

func newRequest(httpMethod string, url string, engine *gin.Engine, body io.Reader, params ...func(r *http.Request)) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(httpMethod, url, body)
	if err != nil {
		return nil, err
	}
	for _, p := range params {
		p(req)
	}

	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)
	return w, nil
}

func HttpMockActivateNonDefault() {
	httpmock.ActivateNonDefault(utils.HttpClient)
}
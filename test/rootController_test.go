package test

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRootController(t *testing.T) {

	r := getRouter()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message" : "hello~~~"})
	})

	req, _ := http.NewRequest("GET", "/", nil)

	testHttpResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOk := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)

		pageOk := err == nil && strings.Index(string(p), "hello") > 0

		return statusOk && pageOk
	})
}
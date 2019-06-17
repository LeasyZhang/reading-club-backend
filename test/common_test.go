package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

//TestMain : set up before testing environment
func TestMain(m *testing.M) {

	gin.SetMode(gin.TestMode)
	//run other test

	os.Exit(m.Run())
}

// GetRouter : helper function to create a router during testing
func getRouter() *gin.Engine {

	r := gin.Default()

	return r
}

// helper function to process a request and return response
func testHttpResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

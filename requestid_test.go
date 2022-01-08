package wrequestid

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const testXRequestID = "test-request-id"

func emptySuccessResponse(c *gin.Context) {
	c.String(http.StatusOK, "")
}


func Test_Default_Header(t *testing.T) {
	r := gin.New()
	r.Use(New())

	w := httptest.NewRecorder()
	r.GET("/", emptySuccessResponse)


	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add(headerXRequestID,"123456789")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "123456789", w.Header().Get(headerXRequestID))
}


func Test_Custom_Header(t *testing.T) {
	r := gin.New()
	r.Use(New(WithRequestIDHeader(testXRequestID)))

	w := httptest.NewRecorder()
	r.GET("/", emptySuccessResponse)


	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add(testXRequestID,"987654321")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "987654321", w.Header().Get(testXRequestID))
}
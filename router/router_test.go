package router

// test router
import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func performRequest(r http.Handler, method, path string, requestBody string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, strings.NewReader(requestBody))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestUploadHandler(t *testing.T) {
	// Arrange
	w := gin.New()
	w.POST("/upload", UploadHandler)

	// Act
	resp := performRequest(w, "POST", "/upload", "test payload")

	// Assert
	assert.Equal(t, 200, resp.Code)
}

func TestVerifyHandler(t *testing.T) {
	// Arrange
	w := gin.New()
	w.POST("/verify", VerifyHandler)

	// Act
	resp := performRequest(w, "POST", "/verify", "test payload")

	// Assert
	assert.Equal(t, 200, resp.Code)
}

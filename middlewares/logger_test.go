package middlewares

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	// Initialize Gin router
	r := gin.Default()

	// Capture logs
	var logBuffer bytes.Buffer
	gin.DefaultWriter = &logBuffer
	gin.DefaultErrorWriter = &logBuffer
	r.Use(Logger())

	// Define a simple handler
	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// Create a test request
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	// Serve the request
	r.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the log output
	logOutput := logBuffer.String()
	expectedLogPrefix := fmt.Sprintf("%s [%s] %s %s ",
		req.RemoteAddr,
		time.Now().Format(time.RFC822),
		http.MethodGet,
		"/test")

	if !contains(logOutput, expectedLogPrefix) {
		t.Errorf("Expected log to contain prefix %s, but got %s", expectedLogPrefix, logOutput)
	}
	if !contains(logOutput, "200") {
		t.Errorf("Expected log to contain status code 200, but got %s", logOutput)
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

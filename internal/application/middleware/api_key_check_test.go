package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestApiKeyCheck(t *testing.T) {
	// Set up a valid API key in the environment
	validApiKey := "test-api-key"

	tests := []struct {
		name           string
		providedApiKey string
		definedApiKey  string
		expectedStatus int
		expectedError  string
	}{
		{
			name:           "No API key defined in environment",
			definedApiKey:  "",
			providedApiKey: validApiKey,
			expectedStatus: http.StatusUnauthorized,
			expectedError:  "no API key has been defined",
		},
		{
			name:           "No API key provided in request",
			definedApiKey:  validApiKey,
			providedApiKey: "",
			expectedStatus: http.StatusUnauthorized,
			expectedError:  "no API key has been provided",
		},
		{
			name:           "Invalid API key provided",
			providedApiKey: "invalid-api-key",
			definedApiKey:  validApiKey,
			expectedStatus: http.StatusUnauthorized,
			expectedError:  "invalid API key",
		},
		{
			name:           "Valid API key provided",
			providedApiKey: validApiKey,
			definedApiKey:  validApiKey,
			expectedStatus: http.StatusOK,
			expectedError:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the Gin context and request
			gin.SetMode(gin.TestMode)
			router := gin.New()
			router.Use(ApiKeyCheck(tt.definedApiKey))
			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "success")
			})

			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			req.Header.Set("X-Feedback-ApiKey", tt.providedApiKey)
			rec := httptest.NewRecorder()

			// Run the middleware
			router.ServeHTTP(rec, req)

			// Check the response status
			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rec.Code)
			}

			// Check the error message if applicable
			if tt.expectedError != "" {
				body := strings.TrimSpace(rec.Body.String())
				if !strings.Contains(body, tt.expectedError) {
					t.Errorf("expected error message to contain '%s', got '%s'", tt.expectedError, body)
				}
			}
		})
	}
}

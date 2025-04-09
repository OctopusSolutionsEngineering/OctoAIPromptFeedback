package middleware

import (
	"errors"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/application/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// ApiKeyCheck is a middleware function that checks for a valid API key in the request header.
func ApiKeyCheck(serviceApiKey string) func(c *gin.Context) {
	return func(c *gin.Context) {
		if len(serviceApiKey) == 0 {
			c.IndentedJSON(http.StatusUnauthorized, responses.GenerateError("Failed to process request", errors.New("no API key has been defined")))
			c.Abort()
			return
		}

		apiKey := strings.TrimSpace(c.GetHeader("X-Feedback-ApiKey"))

		if len(apiKey) == 0 {
			c.IndentedJSON(http.StatusUnauthorized, responses.GenerateError("Failed to process request", errors.New("no API key has been provided")))
			c.Abort()
			return
		}

		if apiKey != serviceApiKey {
			c.IndentedJSON(http.StatusUnauthorized, responses.GenerateError("Failed to process request", errors.New("invalid API key")))
			c.Abort()
			return
		}

		c.Next()
	}
}

package application

import (
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/application/middleware"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/domain/environment"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"strings"
)

// StartServer starts the Gin server and sets up the routes
func StartServer() error {
	apiKey := strings.TrimSpace(os.Getenv("FEEDBACK_SERVICE_API_KEY"))
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Getting resources uses a simple API key to authorize requests
	router.GET("/api/feedback", middleware.ApiKeyCheck(apiKey), GetAllFeedback)
	router.GET("/api/feedback/:id", middleware.ApiKeyCheck(apiKey), GetFeedback)

	// Creating resources uses an Octopus Server JWT to authorize requests
	router.POST("/api/feedback", middleware.JwtCheck, CreateFeedback)

	zap.L().Info("Starting server", zap.String("port", environment.GetPort()))

	return router.Run("localhost:" + environment.GetPort())
}

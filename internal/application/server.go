package application

import (
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/application/middleware"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/domain/environment"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func StartServer() error {
	apiKey := strings.TrimSpace(os.Getenv("FEEDBACK_SERVICE_API_KEY"))
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/api/feedback", middleware.ApiKeyCheck(apiKey), GetAllFeedback)
	router.GET("/api/feedback/:id", middleware.ApiKeyCheck(apiKey), GetFeedback)
	router.POST("/api/feedback", middleware.JwtCheck, CreateFeedback)

	return router.Run("localhost:" + environment.GetPort())
}

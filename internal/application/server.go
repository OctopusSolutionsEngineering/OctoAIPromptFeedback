package application

import (
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/domain/environment"
	"github.com/gin-gonic/gin"
)

func StartServer() error {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/api/feedback", GetAllFeedback)
	router.GET("/api/feedback/:id", GetFeedback)

	return router.Run("localhost:" + environment.GetPort())
}

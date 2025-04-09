package application

import (
	"github.com/DataDog/jsonapi"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/application/responses"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/domain/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// GetHealth responds to https://jsonapi.org/format/#fetching-resources and returns a health check
func GetHealth(c *gin.Context) {
	health := model.Health{
		ID:     uuid.New().String(),
		Status: "healthy",
	}

	jsonApi, err := jsonapi.Marshal(health)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, responses.GenerateError("Failed to process request", err))
		return
	}

	c.String(http.StatusCreated, string(jsonApi))
}

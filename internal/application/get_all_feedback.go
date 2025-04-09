package application

import (
	"github.com/DataDog/jsonapi"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/application/responses"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/infrastructure"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllFeedback responds to https://jsonapi.org/format/#fetching-resources
func GetAllFeedback(c *gin.Context) {
	feedback, err := infrastructure.GetFeedback()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, responses.GenerateError("Failed to process request", err))
		return
	}

	jsonApi, err := jsonapi.Marshal(feedback)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, responses.GenerateError("Failed to process request", err))
		return
	}

	c.String(http.StatusOK, string(jsonApi))
}

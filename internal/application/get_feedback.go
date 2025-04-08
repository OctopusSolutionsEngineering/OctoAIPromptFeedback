package application

import (
	"github.com/DataDog/jsonapi"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/infrastructure"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetFeedback(c *gin.Context) {
	feedback, err := infrastructure.GetFeedback()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, jsonapi.Error{
			Title:  "Failed to retrieve feedback items",
			Detail: err.Error(),
		})
	}

	jsonApi, err := jsonapi.Marshal(feedback)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, jsonapi.Error{
			Title:  "Failed to marshal feedback items",
			Detail: err.Error(),
		})
	}

	c.String(http.StatusOK, string(jsonApi))
}

package application

import (
	"github.com/DataDog/jsonapi"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/infrastructure"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetFeedback(c *gin.Context) {
	id := c.Param("id")

	feedback, found, err := infrastructure.GetFeedbackItem(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, jsonapi.Error{
			Title:  "Failed to retrieve feedback item",
			Detail: err.Error(),
		})
		return
	}

	if !found {
		c.IndentedJSON(http.StatusNotFound, jsonapi.Error{
			Title: "Failed to find feedback item",
		})
		return
	}

	jsonApi, err := jsonapi.Marshal(feedback)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, jsonapi.Error{
			Title:  "Failed to marshal feedback item",
			Detail: err.Error(),
		})
		return
	}

	c.String(http.StatusOK, string(jsonApi))
}

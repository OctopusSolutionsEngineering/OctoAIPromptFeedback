package application

import (
	"github.com/DataDog/jsonapi"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/domain/model"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"time"
)

func CreateFeedback(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, jsonapi.Error{
			Title:  "Failed to read request body",
			Detail: err.Error(),
		})
		return
	}

	var feedback model.Feedback
	err = jsonapi.Unmarshal(body, &feedback)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, jsonapi.Error{
			Title:  "Failed to unmarshal request body",
			Detail: err.Error(),
		})
		return
	}

	newFeedback := model.Feedback{
		ID:        "1",
		Timestamp: time.Now(),
		Prompt:    feedback.Prompt,
		Comment:   feedback.Comment,
		ThumbsUp:  feedback.ThumbsUp,
	}

	jsonApi, err := jsonapi.Marshal(newFeedback)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, jsonapi.Error{
			Title:  "Failed to marshal feedback item",
			Detail: err.Error(),
		})
		return
	}

	c.String(http.StatusCreated, string(jsonApi))
}

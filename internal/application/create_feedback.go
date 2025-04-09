package application

import (
	"github.com/DataDog/jsonapi"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/application/responses"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/domain/jwt"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/domain/model"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/domain/sha"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
	"strings"
	"time"
)

// CreateFeedback https://jsonapi.org/format/#crud-creating
func CreateFeedback(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.GenerateError("Failed to process request", err))
		return
	}

	var feedback model.Feedback
	err = jsonapi.Unmarshal(body, &feedback)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.GenerateError("Failed to process request", err))
		return
	}

	token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

	aud, err := jwt.ValidateJWT(token)

	newFeedback := model.Feedback{
		ID:        uuid.New().String(),
		Timestamp: time.Now(),
		Server:    sha.GetSha256Hash(aud), // We might want this for filtering, but we don't want to save PII (even if it is just a domain name)
		Prompt:    feedback.Prompt,
		Comment:   feedback.Comment,
		ThumbsUp:  feedback.ThumbsUp,
	}

	if err := infrastructure.CreateFeedbackAzureStorageTable(newFeedback); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, responses.GenerateError("Failed to process request", err))
		return
	}

	jsonApi, err := jsonapi.Marshal(newFeedback)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, responses.GenerateError("Failed to process request", err))
		return
	}

	c.String(http.StatusCreated, string(jsonApi))
}

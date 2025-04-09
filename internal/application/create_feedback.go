package application

import (
	"github.com/DataDog/jsonapi"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/domain/jwt"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/domain/model"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func CreateFeedback(c *gin.Context) {
	// At the end of the day, this service is essentially unauthenticated.
	// We accept any user with a valid JWT token that appears to authenticate with an Octopus Deploy instance.
	token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

	aud, err := jwt.ValidateJWT(token)

	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, jsonapi.Error{
			Title:  "The token failed verification",
			Detail: err.Error(),
		})
		return
	}

	apiURL, err := url.Parse(aud)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, jsonapi.Error{
			Title:  "The token failed verification",
			Detail: err.Error(),
		})
		return
	}

	// Use the token to look up the user. This is not foolproof - you could supply any valid JWT token
	// with an audience claim that points to a server that responds to this API request.
	// We can't prove that anyone submitting feedback is a genuine Octopus user.
	// But since we store the audience in the feedback items, we can filter out bad requests later.
	// It also raises the bar for anyone looking to abuse the API, as you would need to generate valid JWTs,
	// host a JWKS server, and host a server that responds the API request.
	octopusClient, err := client.NewClientWithAccessToken(nil, apiURL, token, "")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, jsonapi.Error{
			Title:  "Failed to process request",
			Detail: err.Error(),
		})
		return
	}

	if _, err := octopusClient.Users.GetMe(); err != nil {
		c.IndentedJSON(http.StatusUnauthorized, jsonapi.Error{
			Title:  "The token failed verification",
			Detail: err.Error(),
		})
		return
	}

	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, jsonapi.Error{
			Title:  "Failed to process request",
			Detail: err.Error(),
		})
		return
	}

	var feedback model.Feedback
	err = jsonapi.Unmarshal(body, &feedback)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, jsonapi.Error{
			Title:  "Failed to process request",
			Detail: err.Error(),
		})
		return
	}

	newFeedback := model.Feedback{
		ID:        "1",
		Timestamp: time.Now(),
		Server:    aud,
		Prompt:    feedback.Prompt,
		Comment:   feedback.Comment,
		ThumbsUp:  feedback.ThumbsUp,
	}

	jsonApi, err := jsonapi.Marshal(newFeedback)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, jsonapi.Error{
			Title:  "Failed to process request",
			Detail: err.Error(),
		})
		return
	}

	c.String(http.StatusCreated, string(jsonApi))
}

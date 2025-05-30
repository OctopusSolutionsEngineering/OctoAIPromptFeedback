package middleware

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/application/responses"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/domain/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strings"
)

// JwtCheck is a middleware function that checks for a valid JWT token in the request header.
func JwtCheck(c *gin.Context) {

	// At the end of the day, this service is essentially unauthenticated.
	// We accept any user with a valid JWT token that appears to authenticate with an Octopus Deploy instance.
	token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

	aud, err := jwt.ValidateJWT(token)

	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, responses.GenerateError("Failed to process request", err))
		c.Abort()
		return
	}

	apiURL, err := url.Parse(aud)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, responses.GenerateError("Failed to process request", err))
		c.Abort()
		return
	}

	// Use the token to look up the user. This is not foolproof - you could supply any valid JWT token
	// with an audience claim that points to a server that responds to this API request.
	// We can't prove that anyone submitting feedback is a genuine Octopus user.
	// But we do effectively prove that you own a DNS name, which is almost as good.
	// Since we store the audience in the feedback items, we can filter out bad requests later.
	// It also raises the bar for anyone looking to abuse the API, as you would need to generate valid JWTs,
	// host a JWKS server, and host a server that responds the API request.
	octopusClient, err := client.NewClientWithAccessToken(nil, apiURL, token, "")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, responses.GenerateError("Failed to process request", err))
		c.Abort()
		return
	}

	if _, err := octopusClient.Users.GetMe(); err != nil {
		c.IndentedJSON(http.StatusUnauthorized, responses.GenerateError("Failed to process request", err))
		c.Abort()
		return
	}

	// normal request, and the execution chain is called down
	c.Next()
}

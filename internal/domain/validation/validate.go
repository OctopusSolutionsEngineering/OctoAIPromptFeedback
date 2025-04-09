package validation

import (
	"errors"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/domain/model"
	"strings"
)

func ValidateFeedback(feedback model.Feedback) error {
	if strings.TrimSpace(feedback.ID) == "" {
		return errors.New("id is required")
	}

	if strings.TrimSpace(feedback.Server) == "" {
		return errors.New("server is required")
	}

	return nil
}

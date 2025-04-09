package validation

import (
	"errors"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/domain/model"
)

func ValidateFeedback(feedback model.Feedback) error {
	if feedback.ID == "" {
		return errors.New("id is required")
	}

	return nil
}

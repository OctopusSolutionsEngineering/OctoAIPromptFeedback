package infrastructure

import (
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/domain/model"
	"time"
)

func GetFeedback() ([]model.Feedback, error) {
	return []model.Feedback{
		{
			ID:        "0",
			Timestamp: time.Now(),
			Prompt:    "Sample prompt",
			Comment:   "This is a sample comment",
			ThumbsUp:  false,
		},
	}, nil
}

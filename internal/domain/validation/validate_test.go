package validation

import (
	"errors"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/domain/model"
	"testing"
)

func TestValidateFeedback(t *testing.T) {
	tests := []struct {
		name        string
		feedback    model.Feedback
		expectedErr error
	}{
		{
			name:        "Valid feedback",
			feedback:    model.Feedback{ID: "123", Server: "test-server"},
			expectedErr: nil,
		},
		{
			name:        "Missing ID",
			feedback:    model.Feedback{ID: "", Server: "test-server"},
			expectedErr: errors.New("id is required"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateFeedback(tt.feedback)
			if (err != nil && tt.expectedErr == nil) || (err == nil && tt.expectedErr != nil) || (err != nil && err.Error() != tt.expectedErr.Error()) {
				t.Errorf("ValidateFeedback(%v) = %v, want %v", tt.feedback, err, tt.expectedErr)
			}
		})
	}
}

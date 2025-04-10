package infrastructure

import (
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/domain/model"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"testing"
	"time"
)

// Requires Azurite
// docker run -d -p 10000:10000 -p 10001:10001 -p 10002:10002 mcr.microsoft.com/azure-storage/azurite
// export AzureWebJobsStorage="DefaultEndpointsProtocol=http;AccountName=devstoreaccount1;AccountKey=Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==;TableEndpoint=http://127.0.0.1:10002/devstoreaccount1;"
func TestInfrastructure(t *testing.T) {
	feedback := model.Feedback{
		ID:       uuid.New().String(),
		Created:  time.Now(),
		Server:   "test.com",
		Prompt:   "My Prompt",
		Comment:  "My Comment",
		ThumbsUp: true,
	}

	if err := CreateFeedbackAzureStorageTable(feedback); err != nil {
		t.Errorf("Failed to create feedback: %v", err)
	}

	getFeedback, found, err := GetFeedbackItem(feedback.ID)

	if err != nil {
		t.Errorf("Failed to get feedback: %v", err)
	}

	if !found {
		t.Errorf("Feedback not found")
	}

	if getFeedback.ID != feedback.ID {
		t.Errorf("Expected ID %s, got %s", feedback.ID, getFeedback.ID)
	}

	getAllFeedback, err := GetFeedback()

	if err != nil {
		t.Errorf("Failed to get all feedback: %v", err)
	}

	if !lo.ContainsBy(getAllFeedback, func(item model.Feedback) bool {
		return item.ID == getFeedback.ID
	}) {
		t.Errorf("Feedback not found in all feedback")
	}

}

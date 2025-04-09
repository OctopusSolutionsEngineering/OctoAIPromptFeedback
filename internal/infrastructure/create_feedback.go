package infrastructure

import (
	"context"
	"encoding/json"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/domain/model"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/domain/sha"
	"time"
)
import "github.com/Azure/azure-sdk-for-go/sdk/data/aztables"

// https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/data/aztables
func CreateFeedbackAzureStorageTable(feedback model.Feedback) error {
	service, err := aztables.NewServiceClientFromConnectionString(GetStorageConnectionString(), nil)

	if err != nil {
		return err
	}

	ctx := context.Background()

	client := service.NewClient("Feedback")

	myEntity := aztables.EDMEntity{
		Entity: aztables.Entity{
			PartitionKey: sha.GetSha256Hash(feedback.Server),
			RowKey:       feedback.ID,
		},
		Properties: map[string]any{
			"Comment":   feedback.Comment,
			"Timestamp": feedback.Timestamp.Format(time.RFC3339),
			"Prompt":    feedback.Prompt,
			"ThumbsUp":  feedback.ThumbsUp,
		},
	}
	marshalled, err := json.Marshal(myEntity)

	if _, err := client.AddEntity(ctx, marshalled, nil); err != nil {
		return err
	}

	return nil
}

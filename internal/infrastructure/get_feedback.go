package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/domain/model"
	"time"
)

func GetFeedback() ([]model.Feedback, error) {
	service, err := aztables.NewServiceClientFromConnectionString(GetStorageConnectionString(), nil)

	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	if err := CreateTable(service, ctx); err != nil {
		return nil, err
	}

	client := service.NewClient("Feedback")

	options := &aztables.ListEntitiesOptions{}
	pager := client.NewListEntitiesPager(options)

	retValue := []model.Feedback{}

	pageCount := 0
	for pager.More() {
		response, err := pager.NextPage(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Printf("There are %d entities in page #%d\n", len(response.Entities), pageCount)
		pageCount += 1

		for _, entity := range response.Entities {
			var myEntity aztables.EDMEntity
			err = json.Unmarshal(entity, &myEntity)
			if err != nil {
				return nil, err
			}

			retValue = append(retValue, model.Feedback{
				ID:        myEntity.RowKey,
				Timestamp: GetTimeProperty("Timestamp", myEntity, time.Time{}),
				Server:    GetStringProperty("Server", myEntity, ""),
				Prompt:    GetStringProperty("Prompt", myEntity, ""),
				Comment:   GetStringProperty("Comment", myEntity, ""),
				ThumbsUp:  GetBoolProperty("Server", myEntity, false),
			})
		}
	}

	return retValue, nil
}

func GetFeedbackItem(id string) (model.Feedback, bool, error) {
	return model.Feedback{
		ID:        "0",
		Timestamp: time.Now(),
		Prompt:    "Sample prompt",
		Comment:   "This is a sample comment",
		ThumbsUp:  false,
	}, true, nil
}

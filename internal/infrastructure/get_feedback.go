package infrastructure

import (
	"context"
	"encoding/json"
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

	for pager.More() {
		response, err := pager.NextPage(ctx)
		if err != nil {
			return nil, err
		}

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
	service, err := aztables.NewServiceClientFromConnectionString(GetStorageConnectionString(), nil)

	if err != nil {
		return model.Feedback{}, false, err
	}

	ctx := context.Background()

	if err := CreateTable(service, ctx); err != nil {
		return model.Feedback{}, false, err
	}

	client := service.NewClient("Feedback")

	filter := "RowKey eq '" + id + "'"
	options := &aztables.ListEntitiesOptions{
		Filter: &filter,
	}
	pager := client.NewListEntitiesPager(options)

	for pager.More() {
		response, err := pager.NextPage(ctx)
		if err != nil {
			return model.Feedback{}, false, err
		}

		for _, entity := range response.Entities {
			var myEntity aztables.EDMEntity
			err = json.Unmarshal(entity, &myEntity)
			if err != nil {
				return model.Feedback{}, false, err
			}

			return model.Feedback{
				ID:        myEntity.RowKey,
				Timestamp: GetTimeProperty("Timestamp", myEntity, time.Time{}),
				Server:    GetStringProperty("Server", myEntity, ""),
				Prompt:    GetStringProperty("Prompt", myEntity, ""),
				Comment:   GetStringProperty("Comment", myEntity, ""),
				ThumbsUp:  GetBoolProperty("Server", myEntity, false),
			}, true, nil
		}
	}

	return model.Feedback{}, false, nil
}

package main

import (
	"github.com/OctopusSolutionsEngineering/OctoAIPromptFeedback/internal/application"
	"go.uber.org/zap"
)

func main() {
	if err := application.StartServer(); err != nil {
		zap.L().Error(err.Error())
	}
}

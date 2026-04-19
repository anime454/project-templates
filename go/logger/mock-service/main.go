package main

import (
	"github.com/anime454/project-templates/go/logger/logger"
	"github.com/anime454/project-templates/go/logger/mock-service/handler"
	"github.com/anime454/project-templates/go/logger/mock-service/repo"
	"github.com/anime454/project-templates/go/logger/mock-service/service"
	"github.com/anime454/project-templates/go/logger/model"
)

func main() {

	log := logger.NewLogger(model.LoggerConfig{
		Level: model.DebugLevel,
		Masking: model.ConfigMasking{
			Enabled: true,
		},
	})

	repo := repo.NewRepo(log)
	service := service.NewService(log, repo)
	handler := handler.NewHandler(log, service)

	handler.PrintLog()
}

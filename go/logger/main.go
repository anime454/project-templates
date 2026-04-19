package main

import (
	"github.com/anime454/project-templates/go/logger/logger"
	"github.com/anime454/project-templates/go/logger/model"
)

func main() {

	log := logger.NewLogger(model.LoggerConfig{
		Level: model.DebugLevel,
	})
	log.Debug("Hello, World!")
}

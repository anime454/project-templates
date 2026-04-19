package repo

import (
	"context"

	"github.com/anime454/project-templates/go/logger/logger"
)

type RepoInterface interface {
	PrintLog(ctx context.Context)
}

type Repo struct {
	log logger.LoggerPort
}

func NewRepo(log logger.LoggerPort) *Repo {
	return &Repo{log: log}
}

func (r *Repo) PrintLog(ctx context.Context) {
	log := r.log.WithContext(ctx)
	log.Debug("This is a debug log from Repo")
}

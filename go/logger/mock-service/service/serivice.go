package service

import (
	"context"

	"github.com/anime454/project-templates/go/logger/logger"
	"github.com/anime454/project-templates/go/logger/mock-service/repo"
)

type ServiceInterface interface {
	PrintLog(ctx context.Context)
}

type Service struct {
	log  logger.LoggerPort
	repo repo.RepoInterface
}

func NewService(log logger.LoggerPort, repo repo.RepoInterface) *Service {
	return &Service{log: log, repo: repo}
}

func (s *Service) PrintLog(ctx context.Context) {
	log := s.log.WithContext(ctx)
	log.Debug("this is message from service")
	s.repo.PrintLog(ctx)
}

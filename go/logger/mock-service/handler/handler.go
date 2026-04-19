package handler

import (
	"context"

	"github.com/anime454/project-templates/go/logger/logger"
	"github.com/anime454/project-templates/go/logger/mock-service/service"
)

type Handler struct {
	log     logger.LoggerPort
	service service.ServiceInterface
}

func (h *Handler) PrintLog() {
	ctx := context.WithValue(context.Background(), logger.RequestIDKey, "mock-request-id")
	log := h.log.WithContext(ctx)
	log.Debug("this is message from handler")
	h.service.PrintLog(ctx)
}

func NewHandler(log logger.LoggerPort, service service.ServiceInterface) *Handler {
	return &Handler{
		log:     log,
		service: service,
	}
}

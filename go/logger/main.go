package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/anime454/project-templates/go/logger/logger"
	"github.com/anime454/project-templates/go/logger/model"
)

func main() {

	log := logger.NewLogger(model.LoggerConfig{
		Level: model.DebugLevel,
		Masking: model.ConfigMasking{
			Enabled: true,
			FieldMap: map[string]any{
				"password": "******",
				"token":    "redacted-token",
			},
		},
	})

	ctx := context.WithValue(context.Background(), logger.RequestIDKey, "1234567890")

	log = log.WithContext(ctx)

	type st struct {
		ID       string
		Password string
	}
	mock := st{
		ID:       "user-123",
		Password: "super-secret",
	}
	log.Debug(mock)
	log.Debugf("this is an message template id = %d, and message = %s, password=%v", 1, "template1", map[string]any{"password": "secret"})
	err := errors.New("first error")
	log.Error(fmt.Errorf("second error: %w", err))
}

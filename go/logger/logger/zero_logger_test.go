package logger

import (
	"reflect"
	"testing"
	"time"

	"github.com/anime454/project-templates/go/logger/model"
)

type nestedPayload struct {
	Password string           `json:"password"`
	Tokens   []map[string]any `json:"tokens"`
}

type requestPayload struct {
	UserID   string        `json:"user_id"`
	Password string        `json:"password"`
	Created  time.Time     `json:"created"`
	Nested   nestedPayload `json:"nested"`
}

func TestMaskValueMasksNestedStructFields(t *testing.T) {
	logger := &Logger{
		maskingEnabled: true,
		maskFields: normalizeMaskFields(map[string]any{
			"password": defaultMaskValue,
			"user_id":  "hidden-user",
			"token":    "redacted-token",
		}),
	}

	created := time.Date(2026, time.April, 19, 10, 0, 0, 0, time.UTC)
	masked := logger.maskValue(requestPayload{
		UserID:   "user-123",
		Password: "super-secret",
		Created:  created,
		Nested: nestedPayload{
			Password: "nested-secret",
			Tokens: []map[string]any{{
				"token":  "abc123",
				"public": "value",
			}},
		},
	})

	result, ok := masked.(map[string]any)
	if !ok {
		t.Fatalf("expected masked struct to become a map, got %T", masked)
	}

	if result["user_id"] != "hidden-user" {
		t.Fatalf("expected user_id to be masked, got %#v", result["user_id"])
	}

	if result["password"] != defaultMaskValue {
		t.Fatalf("expected password to be masked, got %#v", result["password"])
	}

	if !reflect.DeepEqual(result["created"], created) {
		t.Fatalf("expected created time to be preserved, got %#v", result["created"])
	}

	nested, ok := result["nested"].(map[string]any)
	if !ok {
		t.Fatalf("expected nested field to be a map, got %T", result["nested"])
	}

	if nested["password"] != defaultMaskValue {
		t.Fatalf("expected nested password to be masked, got %#v", nested["password"])
	}

	tokens, ok := nested["tokens"].([]any)
	if !ok || len(tokens) != 1 {
		t.Fatalf("expected nested tokens slice, got %#v", nested["tokens"])
	}

	tokenEntry, ok := tokens[0].(map[string]any)
	if !ok {
		t.Fatalf("expected token entry to be a map, got %T", tokens[0])
	}

	if tokenEntry["token"] != "redacted-token" {
		t.Fatalf("expected token to be masked, got %#v", tokenEntry["token"])
	}

	if tokenEntry["public"] != "value" {
		t.Fatalf("expected public value to remain unchanged, got %#v", tokenEntry["public"])
	}
}

func TestMaskValueHandlesPointersAndDisabledMasking(t *testing.T) {
	payload := &model.HTTPRequestLog{
		Meta: model.RequestMeta{
			Headers: map[string][]string{
				"Authorization": {"Bearer secret"},
			},
			UserID: "user-999",
		},
	}

	maskedLogger := &Logger{
		maskingEnabled: true,
		maskFields: normalizeMaskFields(map[string]any{
			"Authorization": "[masked]",
			"UserID":        nil,
		}),
	}

	masked := maskedLogger.maskValue(payload)
	result, ok := masked.(map[string]any)
	if !ok {
		t.Fatalf("expected masked payload to become a map, got %T", masked)
	}

	meta := result["Meta"].(map[string]any)
	if meta["UserID"] != defaultMaskValue {
		t.Fatalf("expected UserID to use default mask, got %#v", meta["UserID"])
	}

	headers := meta["Headers"].(map[string]any)
	if headers["Authorization"] != "[masked]" {
		t.Fatalf("expected Authorization header to be masked, got %#v", headers["Authorization"])
	}

	disabledLogger := &Logger{
		maskingEnabled: false,
		maskFields:     normalizeMaskFields(map[string]any{"UserID": defaultMaskValue}),
	}

	if disabledLogger.maskValue(payload) != payload {
		t.Fatal("expected disabled masking to return the original value")
	}
}

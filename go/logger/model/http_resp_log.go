package model

import "time"

type HTTPResponseLog struct {
	Info     ResponseInfo
	Meta     ResponseMeta
	Body     any
	Duration int64 // in ms
}

type ResponseInfo struct {
	Timestamp time.Time
	Status    int
	Size      int64 // bytes
	Protocol  string
}

type ResponseMeta struct {
	RequestID string
	UserID    string
	Headers   map[string][]string
}

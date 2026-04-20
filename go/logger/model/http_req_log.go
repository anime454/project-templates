package model

import "time"

type HTTPRequestLog struct {
	Info RequestInfo
	Meta RequestMeta
	Body any
}

type RequestInfo struct {
	Timestamp time.Time
	Method    string
	Path      string
	IP        string
	Protocol  string
}

type RequestMeta struct {
	RequestID string
	UserID    string
	UserAgent string
	Headers   map[string][]string
}

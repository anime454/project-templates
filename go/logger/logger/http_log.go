package logger

import "time"

type HTTPLog struct {
	Request  HTTPRequestLog
	Response HTTPResponseLog
}

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

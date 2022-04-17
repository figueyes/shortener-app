package model

import "time"

type QueueShort struct {
	ShortUrl         string    `json:"short_url"`
	StatusHttp       int       `json:"status_http"`
	MethodRunner     string    `json:"method_runner"`
	RequestCreatedAt time.Time `json:"request_created_at"`
}

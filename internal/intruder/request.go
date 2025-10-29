package intruder

import (
	"time"
)

// Request Holds the meaningful data that is passed to runner for making the quert
type Request struct {
	Method    string
	URL       string
	Host      string
	Headers   map[string]string
	Data      []byte
	Input     map[string][]byte
	Position  int
	Raw       string
	Error     string
	Timestamp time.Time
}

func NewRequest(conf *Config) Request {
	var req Request
	req.Headers = make(map[string]string)
	req.Method = conf.Method
	req.URL = conf.Url
	return req
}

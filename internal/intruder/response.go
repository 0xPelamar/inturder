package intruder

import (
	"net/http"
	"time"
)

// Response struct holds the meaningful data returned from request.
type Response struct {
	StatusCode    int64
	Headers       map[string][]string
	Data          []byte
	ContentLength int64
	ContentWords  int64
	ContentLines  int64
	ContentType   string
	Cancelled     bool
	Raw           string
	ResultFile    string
	ScraperData   map[string][]string
	Duration      time.Duration
	Timestamp     time.Time
	Request       *Request
}

func NewResponse(httpResp *http.Response, req *Request) Response {
	var resp Response
	resp.Request = req
	resp.StatusCode = int64(httpResp.StatusCode)
	resp.ContentType = httpResp.Header.Get("Content-Type")
	resp.Headers = httpResp.Header
	resp.Cancelled = false
	resp.Raw = ""
	resp.ResultFile = ""
	resp.ScraperData = make(map[string][]string)
	return resp
}

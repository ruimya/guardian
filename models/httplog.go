package models

import (
	"net/http"
	"time"
)

//HTTPLog represents http log
type HTTPLog struct {
	TargetID         string
	RequestURI       string
	StatusCode       int
	RuleCheckElapsed int64
	HTTPElapsed      int64
	RequestSize      int64
	ResponseSize     int64

	timer time.Time
}

//NewHTTPLog inits HTTP log
func NewHTTPLog() *HTTPLog {
	return &HTTPLog{"", "", 0, 0, 0, 0, 0, time.Now()}
}

//Build Fills HTTP Log
func (h HTTPLog) Build(target *Target, request *http.Request, response *http.Response) *HTTPLog {
	h.TargetID = target.ID
	h.RequestURI = request.RequestURI
	h.RequestSize = request.ContentLength

	if response == nil {
		return &h
	}

	h.ResponseSize = response.ContentLength
	h.StatusCode = response.StatusCode
	h.HTTPElapsed = CalcTime(h.timer)

	return &h
}

//NoResponse handles when no response
func (h HTTPLog) NoResponse() *HTTPLog {
	h.StatusCode = -1
	h.HTTPElapsed = CalcTime(h.timer)

	return &h
}

//RuleExecutionEnd Calculates the time for execution of rules
func (h HTTPLog) RuleExecutionEnd() *HTTPLog {
	h.RuleCheckElapsed = CalcTime(h.timer)

	return &h
}

package types

import "time"

type EntryV1 struct {
	CreatedTS time.Time `json:"created-ts"`
	Traffic   string    `json:"traffic"`
	Start     time.Time `json:"start-time"`
	Duration  int

	RequestId string

	// Request attributes
	Url         string // {scheme}://{host}/{path} No query
	Protocol    string
	Host        string
	Path        string
	Method      string
	StatusCode  int32
	StatusFlags string

	Timeout   int32
	RateLimit float64
	RateBurst int32
}

type EntryV2 struct {
	CreatedTS time.Time `json:"created-ts"`
	Traffic   string    `json:"traffic"`
	Start     time.Time `json:"start-time"`
	Duration  int       `json:"duration-ms"`

	RequestId string `json:"request-id"`

	// Request attributes
	Uri            string `json:"uri"` // {scheme}://{host}/{path} No query
	Protocol       string `json:"protocol"`
	Host           string `json:"host"`
	Path           string `json:"path"`
	Method         string `json:"method"`
	StatusCode     int32  `json:"status-code"`
	ThresholdFlags string `json:"threshold-flags"`
	Threshold      int    `json:"threshold"`
}

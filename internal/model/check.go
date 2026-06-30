package model

type Check struct {
	ID         int64
	ServiceID  int64
	StatusCode int
	LatencyMS  int
	Success    int
	Error      string
	CheckedAt  int64
}

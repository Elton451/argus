package model

type Service struct {
	ID              int64
	Name            string
	URL             string
	Method          string
	IntervalSeconds int
	TimeoutSeconds  int
	State           string
	FailureStreak   int
	AlertPolicyID   *int64 
	CreatedAt       int64
	UpdatedAt       int64
}

package model

type Incident struct {
	ID int64
	ServiceID int64
	State string // CLOSED / OPEN 
	TriggerCheckID *int64
	SummaryAI *string
	StartedAt int64
	ResolvedAt *int64
	Acknowledged bool
}

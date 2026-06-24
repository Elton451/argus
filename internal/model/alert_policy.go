package model

type AlertPolicy struct {
	ID int64
	Name string
	ChannelType string
	ChannelConfig string
	FailureThreshold int 
	CooldownMinutes int
	CreatedAt int64
}

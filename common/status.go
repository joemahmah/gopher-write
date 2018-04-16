package common

type Status int8

const (
	StatusNotStarted Status = iota
	StatusInProgress
	StatusAlmostDone
	StatusDone
	StatusUnknown
)

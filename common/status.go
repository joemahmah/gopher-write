package common

//Status is an alias for int
type Status = int

const (
	StatusNotStarted Status = iota
	StatusInProgress
	StatusAlmostDone
	StatusDone
	StatusUnknown
)

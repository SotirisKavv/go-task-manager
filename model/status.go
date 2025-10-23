package model

type Status string

const (
	StatusPending   Status = "Pending"
	StatusCompleted Status = "Completed"
	StatusCancelled Status = "Cancelled"
)

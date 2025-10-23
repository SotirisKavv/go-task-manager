package model

import "time"

type Task struct {
	Id      int       `json:"id" csv:"id"`
	Title   string    `json:"title" csv:"title"`
	Status  Status    `json:"status" csv:"status"`
	DueDate time.Time `json:"due_date" csv:"due_date"`
}

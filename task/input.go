package task

import "encoding/json"

type Task struct {
	Title string		
	Desc string
	Assignee int
	CreatedAt time.Time
	UpdatedAt time.Time
}
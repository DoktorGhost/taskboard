package structs

import "time"

type status string

const (
	StatusPending    status = "в ожидании выполнения"
	StatusInProgress status = "в процессе выполнения"
	StatusCompleted  status = "задача заверешна"
	StatusCancel     status = "задача отменена"
)

type Task struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	DueDate  time.Time `json:"due_date"`
	Priority string    `json:"priority"`
	Status   string    `json:"status"`
	Comments []string  `json:"comments"`
}

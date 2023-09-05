package structs

import "time"

type Status string

const (
	StatusPending    Status = "в ожидании выполнения"
	StatusInProgress Status = "в процессе выполнения"
	StatusCompleted  Status = "задача заверешна"
	StatusCancel     Status = "задача отменена"
)

type Prioritys string

const (
	PriorityLow    Prioritys = "низкий"
	PriorityMedium Prioritys = "средний"
	PriorityHigh   Prioritys = "высокий"
)

type User struct {
	ID         int    `json:"user_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
}

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"Description"`
	DueDate     time.Time `json:"due_date"`
	Status      Status    `json:"status"`
	Priority    Prioritys `json:"priority"`
	Comments    []string  `json:"comments"`
}

type Comment struct {
	ID          int
	Description string
	UserID      int
	TaskID      int
	DueDate     time.Time
}

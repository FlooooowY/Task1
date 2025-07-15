package taskService

type Task struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type TaskRequest struct {
	Name string `json:"name"`
}

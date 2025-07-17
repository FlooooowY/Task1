package taskService

type Task struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Name   string `json:"name"`
	UserId string `gorm:"type:uuid" json:"user_id"`
}

type TaskRequest struct {
	Name   string `json:"name"`
	UserId string `json:"user_id"`
}

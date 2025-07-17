package userService

import taskService "Tasks/internal/taskService"

type Users struct {
	ID       string             `gorm:"primaryKey" json:"id"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Task     []taskService.Task `gorm:"foreignKey:UserId"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

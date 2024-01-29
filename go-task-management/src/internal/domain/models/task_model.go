package models

type Task struct {
	Model
	Name     string `json:"name"`
	Priority int    `gorm:"index:idx_task_priority" json:"priority"`
	IsDone   bool   `gorm:"default:false" json:"isDone"`
	UserId   uint   `json:"userId"`
}

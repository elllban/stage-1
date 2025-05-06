package taskRepository

import "time"

type Tasks struct {
	ID        string `gorm:"primaryKey"`
	Task      string
	IsDone    bool
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

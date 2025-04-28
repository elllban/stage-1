package repository

import "time"

type Tasks struct {
	ID        string `gorm:"primaryKey"`
	Task      string
	IsDone    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

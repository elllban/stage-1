package model

type TaskRequest struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"isDone"`
	UserID string `json:"user_id"`
}

type TaskResponse struct {
	Task   string `json:"task"`
	IsDone bool   `json:"isDone"`
	UserID string `json:"user_id"`
}

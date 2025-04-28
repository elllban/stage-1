package internal

type TaskRequest struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"isDone"`
}

type TaskResponse struct {
	Task   string `json:"task"`
	IsDone bool   `json:"isDone"`
}

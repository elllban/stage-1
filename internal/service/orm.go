package service

type RequestBody struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"isDone"`
}
type Response struct {
	Task   string `json:"task"`
	IsDone bool   `json:"isDone"`
}

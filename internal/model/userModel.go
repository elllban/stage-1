package model

type UserRequest struct {
	ID       string        `gorm:"primaryKey" json:"id"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Tasks    []TaskRequest `json:"tasks"`
}

type UserResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

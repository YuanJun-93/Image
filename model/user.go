package model

type User struct {
	UserID       int64  `json:"user_id" db:"user_id" `
	UserName     string `json:"username" db:"username" binding:"required"`
	PassWord     string `json:"password" db:"password" binding:"required"`
	AccessToken  string
	RefreshToken string
}

package model

type UserLoginRequest struct {
	UserId   int64  `json:"user_id"`
	Password string `json:"password"`
	Type     int    `json:"type"`
}

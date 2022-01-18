package model

// User 用户
type User struct {
	Username string `json:"username"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

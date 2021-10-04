package model

type (
	User struct {
		ID       uint64    `json:"user_id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}
)

package models

type User struct {
	UserID     int    `json:"user_id" db:"user_id"`
	UserName   string `json:"user_name" db:"user_name"`
	FirstName  string `json:"first_name" db:"first_name"`
	LastName   string `json:"last_name" db:"last_name"`
	Email      string `json:"email" db:"email"`
	UserStatus string `json:"user_status" db:"user_status"`
	Department string `json:"department" db:"department"`
}

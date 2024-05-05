package models

type User struct {
	//Id       int    `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role" db:"role"`
	Number   string `json:"number" db:"number"`
}

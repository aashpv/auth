package models

type User struct {
	//Id       int    `json:"id" db:"id"`
	Email    string `json:"email" db:"email" validate:"email, required"`
	Password string `json:"password" db:"password" validate:"required, min=8"`
	//Role     string `json:"role" db:"role" validate:"required, eq=ADMIN|eq=USER"`
	//Number   string `json:"number" db:"number" validate:"required,numeric"`
}

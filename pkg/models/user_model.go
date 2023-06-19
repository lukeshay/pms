package models

type User struct {
	Timestamps
	Id            string `json:"id" binding:"required" db:"id"`
	Email         string `json:"email" binding:"required" db:"email"`
	EmailVerified bool   `json:"emailVerified" binding:"required" db:"email_verified"`
	FirstName     string `json:"firstName" binding:"required" db:"first_name"`
	LastName      string `json:"lastName" binding:"required" db:"last_name"`
	Password      string `json:"password" db:"password"`
}

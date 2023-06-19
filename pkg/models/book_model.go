package models

import (
	"time"
)

type Book struct {
	Timestamps
	Id          string     `json:"id" binding:"required" db:"id"`
	UserId      string     `json:"userId" binding:"required" db:"user_id" fk:"user.id"`
	Title       string     `json:"title" binding:"required" db:"title"`
	Author      string     `json:"author" binding:"required" db:"author"`
	Rating      int8       `json:"rating" db:"rating"`
	PurchasedAt *time.Time `json:"purchasedAt" db:"purchased_at"`
	FinishedAt  *time.Time `json:"finishedAt" db:"finished_at"`
}

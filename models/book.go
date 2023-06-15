package models

import (
	"time"
)

type Book struct {
	Id          string    `json:"id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Author      string    `json:"author" binding:"required"`
	Rating      int8      `json:"rating"`
	PurchasedAt time.Time `json:"purchasedAt"`
	FinishedAt  time.Time `json:"finishedAt"`
}

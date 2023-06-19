package models

import "time"

type Timestamps struct {
	CreatedAt time.Time  `json:"createdAt" binding:"required" db:"created_at"`
	CreatedBy string     `json:"createdBy" binding:"required" db:"created_by"`
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at"`
	DeletedBy *string    `json:"deletedBy" db:"deleted_by"`
	UpdatedAt time.Time  `json:"updatedAt" binding:"required" db:"updated_at"`
	UpdatedBy string     `json:"updatedBy" binding:"required" db:"updated_by"`
}

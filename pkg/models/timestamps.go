package models

import "time"

// Timestamps
// @Description Timestamps for a record
type Timestamps struct {
	// When the record was created
	CreatedAt time.Time `json:"createdAt" binding:"required" db:"created_at"`
	// Who created the record
	CreatedBy string `json:"createdBy" binding:"required" db:"created_by"`
	// When the record was deleted
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at"`
	// Who deleted the record
	DeletedBy *string `json:"deletedBy" db:"deleted_by"`
	// When the record was updated
	UpdatedAt time.Time `json:"updatedAt" binding:"required" db:"updated_at"`
	// Who updated the record
	UpdatedBy string `json:"updatedBy" binding:"required" db:"updated_by"`
}

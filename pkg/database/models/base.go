package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// Base is the base model for other models.
type Base struct {
	ID        uuid.UUID `gorm:"primary_key; unique; type:uuid; column:id; default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

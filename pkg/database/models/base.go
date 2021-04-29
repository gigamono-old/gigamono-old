package models

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
)

// Base is the base model for other models.
type Base struct {
	ID        uuid.UUID  `pg:",pk, unique, notnull, type:uuid, default:uuid_generate_v4()"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `pg:",soft_delete" json:"deleted_at"`
}

// BeforeInsert to update CreatedAt and UpdatedAt columns.
func (model *Base) BeforeInsert(ctx context.Context) error {
	now := time.Now()
	if model.CreatedAt.IsZero() {
		model.CreatedAt = now
	}
	if model.UpdatedAt.IsZero() {
		model.UpdatedAt = now
	}
	return nil
}

// BeforeUpdate to update UpdatedAt columns.
func (model *Base) BeforeUpdate(ctx context.Context) error {
	model.UpdatedAt = time.Now()
	return nil
}

// BaseNoID is the base model for other models.
type BaseNoID struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `pg:",soft_delete" json:"deleted_at"`
}

// BeforeInsert to update CreatedAt and UpdatedAt columns.
func (model *BaseNoID) BeforeInsert(ctx context.Context) error {
	now := time.Now()
	if model.CreatedAt.IsZero() {
		model.CreatedAt = now
	}
	if model.UpdatedAt.IsZero() {
		model.UpdatedAt = now
	}
	return nil
}

// BeforeUpdate to update UpdatedAt columns.
func (model *BaseNoID) BeforeUpdate(ctx context.Context) error {
	model.UpdatedAt = time.Now()
	return nil
}

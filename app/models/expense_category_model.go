package models

import (
	"time"

	"github.com/google/uuid"
)

type ExpenseCategory struct {
	ID          uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	Name        string    `db:"name" json:"name" validate:"required,lte=255"`
	Description string    `db:"description" json:"description" validate:"required,lte=255"`
}

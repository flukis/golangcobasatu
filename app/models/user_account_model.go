package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Name      string    `db:"name" json:"name" validate:"required,lte=255"`
	Email     string    `db:"email" json:"email" validate:"required,lte=255"`
	Password  string    `db:"password" json:"password" validate:"required,lte=255,gte=8"`
}

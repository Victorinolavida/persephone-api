package models

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID          uuid.UUID
	Name        string
	Stock       int
	Price       int // to validdate
	CreatedBy   uuid.UUID
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time //create index by name
}

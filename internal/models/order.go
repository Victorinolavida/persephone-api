package models

import (
	"github.com/google/uuid"
	"time"
)

type OrderStatus string

var (
	OrderPending   = OrderStatus("pending")
	OrderCompleted = OrderStatus("completed")
	OrderCancelled = OrderStatus("cancelled")
)

type Order struct {
	ID          uuid.UUID
	Total       int //to validate
	CustomerID  uuid.UUID
	SellerID    uuid.UUID
	State       string // PENDING,SEND , CANCELLED
	CreatedAt   time.Time
	CompletedAt time.Time
	CancelledAt time.Time
}

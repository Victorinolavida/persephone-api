package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type UserRole string

var (
	Seller = UserRole("seller")
	Client = UserRole("customer")
	Admin  = UserRole("admin")
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            uuid.UUID `json:"id" bun:"id,pk"`
	FirstName     string    `json:"first_name" bun:"first_name"`
	LastName      string    `json:"last_name" bun:"last_name"`
	Email         string    `json:"email" bun:"email"`
	Role          UserRole  `json:"rol"`
	Phone         string    `json:"phone"`
	CompanyId     uuid.UUID `json:"company_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdateAt      time.Time `json:"update_at"`
	HashPassword  byte      `json:"-"`
}

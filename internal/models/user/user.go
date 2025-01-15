package user

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type UserRol string

var (
	Seller   = UserRol("seller")
	Customer = UserRol("customer")
	Admin    = UserRol("admin")
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            uuid.UUID `json:"id" bun:"id,pk"`
	FirstName     string    `json:"first_name" bun:"first_name"`
	LastName      string    `json:"last_name" bun:"last_name"`
	Email         string    `json:"email" bun:"email"`
	Rol           UserRol   `json:"rol"`
	Phone         string    `json:"phone"`
	CompanyId     uuid.UUID `json:"company_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	HashPassword  []byte    `json:"-" bun:"password_hash"`
}

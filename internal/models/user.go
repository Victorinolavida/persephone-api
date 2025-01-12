package models

import (
	"github.com/google/uuid"
	"time"
)

type UserRole string

var (
	Seller = UserRole("seller")
	Client = UserRole("customer")
	Admin  = UserRole("admin")
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Role      UserRole
	Phone     string
	CompanyId uuid.UUID
	CreatedAt time.Time
	UpdateAt  time.Time
}

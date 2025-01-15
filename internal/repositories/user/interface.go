package user

import (
	"context"
	"github.com/Victorinolavida/persephone-api/internal/models/user"
	"github.com/google/uuid"
)

type Repository interface {
	GetByID(id uuid.UUID) (*user.User, error)
	Create(ctx context.Context, user *user.User) error
}

package user

import (
	"context"
	"github.com/Victorinolavida/persephone-api/internal/adapters/dto"
	"github.com/Victorinolavida/persephone-api/internal/models/user"
	"github.com/google/uuid"
)

type Service interface {
	GetByID(ctx context.Context, id uuid.UUID) (*user.User, error)
	Create(ctx context.Context, data *user.User) error
	ValidateUserData(ctx context.Context, user dto.User) error
}

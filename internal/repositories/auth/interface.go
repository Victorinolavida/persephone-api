package auth

import (
	"context"
	userModel "github.com/Victorinolavida/persephone-api/internal/models/user"
)

type AuthRepository interface {
	CreateToken(ctx context.Context, user userModel.User) error
	GetToken(ctx context.Context, Token string) (string, error)
	DeleteToken(ctx context.Context, id string) error
}

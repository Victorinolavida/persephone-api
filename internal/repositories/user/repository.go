package user

import (
	"context"
	"github.com/Victorinolavida/persephone-api/internal/models/user"
	"github.com/Victorinolavida/persephone-api/pkg/logger"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type UserRepo struct {
	DB *bun.DB
}

func NewUserRepository(db *bun.DB) *UserRepo {
	return &UserRepo{
		db,
	}
}

func (repo *UserRepo) GetByID(id uuid.UUID) (*user.User, error) {
	return &user.User{}, nil
}

func (repo *UserRepo) Create(ctx context.Context, user *user.User) error {
	_, err := repo.DB.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		logger.GetLogger().Error(err)
		return err
	}
	return nil
}

package user

import (
	"github.com/Victorinolavida/persephone-api/internal/models"
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

func (repo *UserRepo) GetByID(id uuid.UUID) (*models.User, error) {
	return &models.User{}, nil
}

func (repo *UserRepo) Create(data any) (*models.User, error) {
	return &models.User{}, nil
}

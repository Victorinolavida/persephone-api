package user

import (
	"context"
	"github.com/Victorinolavida/persephone-api/internal/adapters/dto"
	"github.com/Victorinolavida/persephone-api/internal/lib"
	userModel "github.com/Victorinolavida/persephone-api/internal/models/user"
	"github.com/Victorinolavida/persephone-api/internal/repositories/user"
	"github.com/Victorinolavida/persephone-api/pkg/logger"
	"github.com/google/uuid"
)

type Serv struct {
	Repo      *user.UserRepo
	validator *lib.Validator
}

func NewUserService(repo *user.UserRepo, validator *lib.Validator) *Serv {
	return &Serv{
		repo,
		validator,
	}
}

func (s *Serv) GetByID(ctx context.Context, id uuid.UUID) (*userModel.User, error) {
	return s.Repo.GetByID(id)
}

func (s *Serv) Create(ctx context.Context, data *userModel.User) error {
	return s.Repo.Create(context.TODO(), user)
}
func (s *Serv) ValidateUserData(ctx context.Context, user dto.User) error {
	err := s.validator.Validate.Struct(user)
	if err != nil {
		return err
	}
	err = s.validator.UserPasswordMatch(user)

	if err != nil {
		logger.GetLogger().Error(err)
		return err
	}
	return nil
}

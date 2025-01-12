package user

import (
	"github.com/Victorinolavida/persephone-api/internal/models"
	"github.com/Victorinolavida/persephone-api/internal/repositories/user"
	"github.com/google/uuid"
)

type Serv struct {
	Repo *user.UserRepo
}

func NewUserService(repo *user.UserRepo) *Serv {
	return &Serv{
		repo,
	}
}

func (s *Serv) GetByID(id uuid.UUID) (*models.User, error) {
	return s.Repo.GetByID(id)
}

func (s *Serv) Create(data any) (*models.User, error) {
	return s.Repo.Create(data)
}

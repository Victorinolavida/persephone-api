package user

import (
	"github.com/Victorinolavida/persephone-api/internal/models"
	"github.com/google/uuid"
)

type Repository interface {
	GetByID(id uuid.UUID) (*models.User, error)
	Create(data any) (*models.User, error)
}

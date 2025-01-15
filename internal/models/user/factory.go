package user

import (
	"github.com/Victorinolavida/persephone-api/internal/adapters/dto"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var costHash = 12

func CreateUserFromDTO(user dto.User) (*User, error) {
	u := &User{
		ID:        uuid.New(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Rol:       Customer,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	hashPassword, err := setHashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	u.HashPassword = hashPassword
	return u, nil
}

func setHashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), costHash)
	return bytes, err
}

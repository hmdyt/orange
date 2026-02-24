package usecase

import (
	"fmt"

	"github.com/hmdyt/orange/domain"
	"github.com/hmdyt/orange/usecase/port"
	"github.com/google/uuid"
)

type LoginUsecase struct {
	userRepo port.UserRepository
}

func NewLoginUsecase(userRepo port.UserRepository) *LoginUsecase {
	return &LoginUsecase{userRepo: userRepo}
}

func (u *LoginUsecase) Execute(name string) (*domain.User, string, error) {
	user := &domain.User{
		ID:   uuid.NewString(),
		Name: name,
	}
	if err := u.userRepo.Save(user); err != nil {
		return nil, "", fmt.Errorf("failed to save user: %w", err)
	}
	message := fmt.Sprintf("ようこそ、%sさん！", name)
	return user, message, nil
}

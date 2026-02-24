package server

import (
	"github.com/hmdyt/orange/domain"
	gamev1 "github.com/hmdyt/orange/gen/game/v1"
)

func UserToLoginResponse(user *domain.User, message string) *gamev1.LoginResponse {
	return &gamev1.LoginResponse{
		UserId:  user.ID,
		Name:    user.Name,
		Message: message,
	}
}

package server

import (
	"context"

	"connectrpc.com/connect"
	gamev1 "github.com/hmdyt/orange/gen/game/v1"
	"github.com/hmdyt/orange/usecase"
)

type GameServiceHandler struct {
	loginUsecase *usecase.LoginUsecase
}

func NewGameServiceHandler(loginUsecase *usecase.LoginUsecase) *GameServiceHandler {
	return &GameServiceHandler{loginUsecase: loginUsecase}
}

func (h *GameServiceHandler) Login(
	ctx context.Context,
	req *connect.Request[gamev1.LoginRequest],
) (*connect.Response[gamev1.LoginResponse], error) {
	user, message, err := h.loginUsecase.Execute(req.Msg.Name)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(UserToLoginResponse(user, message)), nil
}

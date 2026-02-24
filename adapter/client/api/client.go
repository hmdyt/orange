package api

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	gamev1 "github.com/hmdyt/orange/gen/game/v1"
	"github.com/hmdyt/orange/gen/game/v1/gamev1connect"
)

type Client struct {
	gameService gamev1connect.GameServiceClient
}

func NewClient(baseURL string) *Client {
	return &Client{
		gameService: gamev1connect.NewGameServiceClient(http.DefaultClient, baseURL),
	}
}

func (c *Client) Login(ctx context.Context, name string) (*gamev1.LoginResponse, error) {
	resp, err := c.gameService.Login(ctx, connect.NewRequest(&gamev1.LoginRequest{
		Name: name,
	}))
	if err != nil {
		return nil, err
	}
	return resp.Msg, nil
}

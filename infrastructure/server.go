package infrastructure

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hmdyt/orange/adapter/server"
	"github.com/hmdyt/orange/gen/game/v1/gamev1connect"
	"github.com/hmdyt/orange/usecase"
)

func RunServer() error {
	userRepo := NewMemoryUserRepository()
	loginUsecase := usecase.NewLoginUsecase(userRepo)
	handler := server.NewGameServiceHandler(loginUsecase)

	mux := http.NewServeMux()

	path, connectHandler := gamev1connect.NewGameServiceHandler(handler)
	mux.Handle(path, connectHandler)

	mux.Handle("/", http.FileServer(http.Dir("static")))

	addr := ":8080"
	log.Printf("server listening on %s", addr)
	return fmt.Errorf("server stopped: %w", http.ListenAndServe(addr, mux))
}

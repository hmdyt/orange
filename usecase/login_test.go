package usecase_test

import (
	"strings"
	"testing"

	"github.com/hmdyt/orange/infrastructure"
	"github.com/hmdyt/orange/usecase"
)

func TestLoginUsecase(t *testing.T) {
	repo := infrastructure.NewMemoryUserRepository()
	uc := usecase.NewLoginUsecase(repo)

	user, message, err := uc.Execute("テスト太郎")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user.ID == "" {
		t.Fatal("user ID should not be empty")
	}
	if user.Name != "テスト太郎" {
		t.Fatalf("expected name テスト太郎, got %s", user.Name)
	}
	if !strings.Contains(message, "テスト太郎") {
		t.Fatalf("expected message to contain テスト太郎, got %s", message)
	}
	if !strings.Contains(message, "ようこそ") {
		t.Fatalf("expected message to contain ようこそ, got %s", message)
	}
}

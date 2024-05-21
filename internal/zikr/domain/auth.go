package domain

import (
	"context"
)

type User struct {
	Guid     string
	Email    string
	Username string
}

type AuthUsecase interface {
	CreateUser(ctx context.Context, user *User) (string, error)
	GetUserInfo(ctx context.Context, email string) (string, error)
	GetOrCreateUser(ctx context.Context, req *User) (string, error)
}

type AuthRepository interface {
	CreateUser(ctx context.Context, user *User) (string, error)
	GetUserInfo(ctx context.Context, email string) (string, error)
}

package domain

import (
	"context"
)

type User struct {
	Guid     string
	Email    string
	Username string
}

type Guid struct {
	Guid string
}

type AuthUsecase interface {
	CreateUser(ctx context.Context, user *User) (string, error)
	GetUserInfo(ctx context.Context, email string) (string, error)
	GetOrCreateUser(ctx context.Context, req *User) (string, error)
}

type AuthRepository interface {
	CreateUser(ctx context.Context, user *User) (string, error)
	CreateZikrsForUser(ctx context.Context, userGuid, zikrGuid string) error
	GetAllZikrGuid(ctx context.Context) ([]Guid, error)
	GetUserInfo(ctx context.Context, email string) (string, error)
}

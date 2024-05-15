package domain

import (
	"context"
)

type User struct {
	Guid           string
	Email          string
	UniqueUsername string
}

type UserLoginRequest struct {
	Email          string
	UniqueUsername string
}

type AuthUsecase interface {
	CreateUser(ctx context.Context, user *User) error
	CheckUser(ctx context.Context, request UserLoginRequest) (bool, error)
}

type AuthRepository interface {
	CreateUser(ctx context.Context, user *User) error
	UserExistsByMail(ctx context.Context, mail string) (bool, error)
	UserExistsByUsername(ctx context.Context, username string) (bool, error)
}

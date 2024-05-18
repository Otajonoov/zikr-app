package domain

import (
	"context"
)

type User struct {
	Guid           string `json:"guid"`
	Email          string `json:"email"`
	UniqueUsername string `json:"unique_username"`
}

type UserLoginRequest struct {
	Email          string `json:"email"`
	UniqueUsername string `json:"unique_username"`
}

type AuthUsecase interface {
	CreateUser(ctx context.Context, user *User) error
	CheckUser(ctx context.Context, request UserLoginRequest) (string, error)
}

type AuthRepository interface {
	CreateUser(ctx context.Context, user *User) error
	UserExistsByMail(ctx context.Context, mail string) (bool, error)
	UserExistsByUsername(ctx context.Context, username string) (bool, error)
}

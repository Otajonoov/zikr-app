package domain

import (
	"context"
)

type User struct {
	FIO           string
	UniqeUsername string
	Password      string
	PhoneNumber   string
}

type AuthUsecase interface {
	CreateUser(ctx context.Context, user *User) error
	CheckUser(ctx context.Context, userName, password string) (bool, error)
	GetByUserName(ctx context.Context, userName string) (*User, error)
}

type AuthRepository interface {
	UserExists(ctx context.Context, username string) (bool, error)
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, username string) (*User, error)
	FindOneByUsername(ctx context.Context, userName string) (*User, error)
}

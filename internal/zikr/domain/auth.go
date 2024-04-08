package domain

import "context"

type User struct {
	FIO           string
	UniqeUsername string
	Password      string
	PhoneNumber   string
}

type AuthUsecase interface {
	CreateUser(ctx context.Context, user *User) error
	CheckUser(ctx context.Context, user *User) bool
}

type AuthRepository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, phone string) (*User, error)
}

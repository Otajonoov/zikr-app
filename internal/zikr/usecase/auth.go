package usecase

import (
	"context"
	"errors"
	"strings"
	"zikr-app/internal/zikr/domain"
)

type authUsecase struct {
	repo domain.AuthRepository
}

func NewAuthUsecase(repo domain.AuthRepository) *authUsecase {
	return &authUsecase{
		repo: repo,
	}
}

func (a *authUsecase) CreateUser(ctx context.Context, user *domain.User) error {

	exists, err := a.repo.UserExists(ctx, user.UniqeUsername)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("user already exists")
	} else {
		err = a.repo.CreateUser(ctx, user)
		if err != nil {
			return errors.New("failed to create user")
		}
		return nil
	}
}

func (a *authUsecase) CheckUser(ctx context.Context, userName, password string) (bool, error) {

	userData, err := a.repo.GetUser(ctx, userName)
	if err != nil {
		return false, errors.New("failed to get user")
	}
	if userData.Password != password {
		return false, errors.New("wrong password")
	}

	return true, nil
}

func (a *authUsecase) GetByUserName(ctx context.Context, userName string) (*domain.User, error) {
	if strings.TrimSpace(userName) == "" {
		return nil, errors.New("empty username")
	}

	userData, err := a.repo.FindOneByUsername(ctx, userName)
	if err != nil {
		return nil, errors.New("failed to get user")
	}
	return userData, nil
}

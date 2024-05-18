package usecase

import (
	"context"
	"errors"
	"zikr-app/internal/zikr/domain"
)

type authUsecase struct {
	repo domain.AuthRepository
	BaseUseCase
}

func NewAuthUsecase(repo domain.AuthRepository) *authUsecase {
	return &authUsecase{
		repo: repo,
	}
}

func (a *authUsecase) CreateUser(ctx context.Context, user *domain.User) error {
	a.beforeRequestForUser(user)

	err := a.repo.CreateUser(ctx, user)
	if err != nil {
		return errors.New("failed to create user")
	}
	return nil
}

func (a *authUsecase) CheckUser(ctx context.Context, request domain.UserLoginRequest) (string, error) {
	user := &domain.User{}

	existMail, _ := a.repo.UserExistsByMail(ctx, request.Email)
	//existUserName, _ := a.repo.UserExistsByUsername(ctx, request.UniqueUsername)
	if !existMail {
		a.beforeRequestForUser(user)
		user.Email = request.Email
		user.UniqueUsername = request.UniqueUsername
		err := a.repo.CreateUser(context.Background(), user)
		if err != nil {
			return "", errors.New("failed to create user")
		}
	}

	return user.Guid, nil
}

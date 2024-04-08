package usecase

import (
	"context"
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

	//u, err := a.repo.GetUser(ctx, user.PhoneNumber, user.UniqeUsername)
	//if err != nil {
	//	return err
	//}
	//
	//if u.PhoneNumber == user.PhoneNumber {
	//	return errors.New("user is already registered")
	//}

	if err := a.repo.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil

}

func (a *authUsecase) CheckUser(ctx context.Context, user *domain.User) bool {

	userData, err := a.repo.GetUser(ctx, user.PhoneNumber)
	if err != nil {
		return false
	}
	if userData.Password != user.Password {
		return false
	}

	return true
}

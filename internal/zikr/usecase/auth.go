package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/port/model"
)

type authUsecase struct {
	repo     domain.AuthRepository
	zikrRepo domain.ZikrRepo
	BaseUseCase
}

func NewAuthUsecase(repo domain.AuthRepository, zikrRepo domain.ZikrRepo) *authUsecase {
	return &authUsecase{
		repo:     repo,
		zikrRepo: zikrRepo,
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

func (a *authUsecase) CheckUser(ctx context.Context, request model.UserLoginRequest) (string, error) {
	user := &domain.User{}

	existMail, err := a.repo.UserExistsByMail(ctx, request.Email)
	if err != nil {
		return "", errors.New("error checking user by email")
	}
	existUserName, err := a.repo.UserExistsByUsername(ctx, request.UniqueUsername)
	if err != nil {
		return "", errors.New("error checking user by username")
	}

	if !existMail && !existUserName {
		a.beforeRequestForUser(user)
		user.Email = request.Email
		user.UniqueUsername = request.UniqueUsername
		err := a.repo.CreateUser(context.Background(), user)
		if err != nil {
			return "", errors.New("failed to create user")
		}
		return user.Guid, nil
	}

	zikrs, err := a.zikrRepo.GetUserZikrByMail(request.Email, request.UniqueUsername)
	if err != nil {
		return "", err
	}
	var zikrStrings []string
	for _, zikr := range zikrs {
		zikrStr := fmt.Sprintf("GUID: %s, UserGuid: %s, Arabic: %s, Uzbek: %s, Pronounce: %s, Count: %d, IsFavorite: %t, CreatedAt: %s, UpdatedAt: %s",
			zikr.GetGuid(), zikr.GetUserGUID(), zikr.GetArabic(), zikr.GetUzbek(), zikr.GetPronounce(), zikr.GetCount(), zikr.GetIsFavorite(), zikr.GetCreatedAt(), zikr.GetUpdatedAt())
		zikrStrings = append(zikrStrings, zikrStr)
	}
	result := strings.Join(zikrStrings, ", ")

	return result, nil
}

package usecase

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"log"
	"zikr-app/internal/zikr/domain"
)

type authUseCase struct {
	repo     domain.AuthRepository
	zikrRepo domain.ZikrRepo
	BaseUseCase
}

func NewAuthUsecase(repo domain.AuthRepository, zikrRepo domain.ZikrRepo) *authUseCase {
	return &authUseCase{
		repo:     repo,
		zikrRepo: zikrRepo,
	}
}

func (a *authUseCase) CreateUser(ctx context.Context, user *domain.User) (string, error) {
	a.beforeRequestForUser(user)

	guid, err := a.repo.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}
	return guid, nil
}

func (a *authUseCase) GetUserInfo(ctx context.Context, email string) (string, error) {
	guid, err := a.repo.GetUserInfo(ctx, email)
	if err != nil {
		return "", err
	}
	return guid, nil
}

func (a *authUseCase) GetOrCreateUser(ctx context.Context, req *domain.User) (string, error) {
	guid, err := a.GetUserInfo(ctx, req.Email)
	if errors.Is(err, pgx.ErrNoRows) {
		u, err := a.CreateUser(ctx, req)
		if err != nil {
			return "", err
		}

		guids, err := a.repo.GetAllZikrGuid(ctx)
		if err != nil {
			log.Println("err: ", err)
		}

		for _, g := range guids {
			if err := a.repo.CreateZikrsForUser(ctx, u, g.Guid); err != nil {
				log.Println("err: ", err)
			}
		}
		guid = u
	}

	return guid, nil
}

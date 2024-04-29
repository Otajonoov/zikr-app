package usecase

import (
	"context"
	"zikr-app/internal/zikr/domain"
)

type zikrCountUsecase struct {
	repo domain.ZikrCountRepository
}

func NewZikrCountUsecase(repo domain.ZikrCountRepository) domain.ZikrCountUsecase {
	return &zikrCountUsecase{
		repo: repo,
	}
}

func (z zikrCountUsecase) CreateCount(ctx context.Context, count *domain.ZikrCount) error {
	err := z.repo.CreateZikrCount(context.Background(), count)
	if err != nil {
		return err
	}

	return nil
}

func (z zikrCountUsecase) GetUserCounts(ctx context.Context, userId int) (map[string]int, error) {
	res, err := z.GetUserCounts(context.Background(), userId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (z zikrCountUsecase) PatchCount(ctx context.Context, userId int, countID int) error {
	err := z.repo.UpdateUserCount(context.Background(), userId, countID)
	if err != nil {
		return err
	}

	return nil
}

func (z zikrCountUsecase) ResetCount(ctx context.Context, userId int) error {
	err := z.repo.DeleteCount(context.Background(), userId)
	if err != nil {
		return err
	}

	return nil
}

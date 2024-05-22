package usecase

import (
	"log"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/domain/factory"
)

type zikrFavoritesUsecase struct {
	repo    domain.ZikrFavoritesRepository
	factory factory.Factory
}

func NewZikrFavoritesUsecase(repo domain.ZikrFavoritesRepository) domain.ZikrFavoritesUsecase {
	return &zikrFavoritesUsecase{
		repo: repo,
	}
}

func (z zikrFavoritesUsecase) Update(userId, zikrId string, isFavorite bool) error {
	if err := z.repo.Update(userId, zikrId, isFavorite); err != nil {
		log.Println("err: ", err)
		return err
	}
	return nil
}

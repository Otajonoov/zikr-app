package usecase

import (
	"zikr-app/internal/zikr/domain"
)

type zikrFavoritesUsecase struct {
	repo    domain.ZikrFavoritesRepository
	factory domain.ZikrFactory
}

func NewZikrFavoritesUsecase(repo domain.ZikrFavoritesRepository, factory domain.ZikrFactory) domain.ZikrFavoritesUsecase {
	return &zikrFavoritesUsecase{
		repo:    repo,
		factory: factory,
	}
}

func (z zikrFavoritesUsecase) FavoriteDua(userId, zikrId string) (bool, error) {
	ok, err := z.repo.FavoriteDua(userId, zikrId)
	if err != nil {
		return false, err
	}

	return ok, nil
}

func (z zikrFavoritesUsecase) UnFavoriteDua(userId, zikrId string) (bool, error) {
	ok, err := z.repo.UnFavoriteDua(userId, zikrId)
	if err != nil {
		return false, err
	}

	return ok, nil
}

func (z zikrFavoritesUsecase) GetAllFavorites(userId string) (zikrs []domain.Zikr, err error) {
	favorites, err := z.repo.GetAllFavorites(userId)
	if err != nil {
		return nil, err
	}

	return favorites, nil
}

func (z zikrFavoritesUsecase) GetAllUnFavorites(userId string) (zikrs []domain.Zikr, err error) {
	favorites, err := z.repo.GetAllUnFavorites(userId)
	if err != nil {
		return nil, err
	}

	return favorites, nil
}

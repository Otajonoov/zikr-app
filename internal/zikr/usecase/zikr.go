package usecase

import (
	"log"
	"zikr-app/internal/zikr/domain"
)

type zikrUsecase struct {
	BaseUseCase
	repo    domain.ZikrRepo
	factory domain.ZikrFactory
}

func NewZikrUsecase(repo domain.ZikrRepo, factory domain.ZikrFactory) domain.ZikrUsecase {
	return &zikrUsecase{
		repo:    repo,
		factory: factory,
	}
}

func (z zikrUsecase) Create(zikr *domain.Zikr) error {
	z.beforeRequest(zikr)
	err := z.repo.Create(zikr)
	if err != nil {
		return err
	}

	return nil
}

func (z zikrUsecase) Get(id int) (zikr *domain.Zikr, err error) {
	zikr, err = z.repo.Get(id)
	if err != nil {
		return nil, err
	}

	return zikr, nil
}

func (z zikrUsecase) GetAll() (zikrs []domain.Zikr, err error) {
	zikrs, err = z.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return zikrs, nil
}

func (z zikrUsecase) FavoritedDua(userId, zikrId int) (bool, error) {
	ok, err := z.repo.FavoriteDua(userId, zikrId)
	if err != nil {
		return false, nil
	}

	return ok, nil
}

func (z zikrUsecase) UnFavoritedDua(userId, zikrId int) (bool, error) {
	ok, err := z.repo.UnFavoriteDua(userId, zikrId)
	if err != nil {
		return false, nil
	}

	return ok, nil
}

func (z zikrUsecase) GetAllFavoriteDuas(userId int) (zikrs []domain.Zikr, err error) {
	favorites, err := z.repo.GetAllFavorites(userId)
	if err != nil {
		return nil, err
	}

	zikrs = make([]domain.Zikr, 0, len(favorites))
	for _, zikr := range favorites {
		zikrs = append(zikrs, zikr)
		log.Println(zikr.GetGUID(), zikr.GetUserId(), zikr.GetArabic(), zikr.GetUzbek(), zikr.GetPronounce(), zikr.GetIsFavourite(),
			zikr.GetCreatedAt(), zikr.GetUpdatedAt())
	}

	return zikrs, nil
}

func (z zikrUsecase) Update(zikr *domain.Zikr) error {
	err := z.repo.Update(zikr)
	if err != nil {
		return err
	}

	return nil
}

func (z zikrUsecase) Delete(id int) error {
	err := z.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

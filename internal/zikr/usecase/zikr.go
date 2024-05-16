package usecase

import (
	"time"
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

func (z *zikrUsecase) Create(zikr *domain.Zikr) error {
	z.beforeRequest(zikr)

	err := z.repo.Create(zikr)
	if err != nil {
		return err
	}

	return nil
}

func (z *zikrUsecase) Get(guid string) (zikr *domain.Zikr, err error) {

	zikr, err = z.repo.Get(guid)
	if err != nil {
		return &domain.Zikr{}, err
	}

	return zikr, nil
}

func (z *zikrUsecase) GetAll() (zikrs []domain.Zikr, err error) {
	zikrs, err = z.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return zikrs, nil
}

func (z *zikrUsecase) Update(zikr *domain.Zikr) error {
	err := z.repo.Update(zikr)
	if err != nil {
		return err
	}
	zikr.SetUpdatedAt(time.Now())
	return nil
}

func (z *zikrUsecase) UpdateZikrCount(updateZikr *domain.Zikr) error {
	err := z.repo.UpdateZikrCount(updateZikr)
	if err != nil {
		return err
	}
	return nil
}

func (z *zikrUsecase) Delete(guid string) error {
	err := z.repo.Delete(guid)
	if err != nil {
		return err
	}
	return nil
}

package usecase

import (
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

func (z zikrUsecase) Get(id string) (zikr *domain.Zikr, err error) {
	zikr, err = z.repo.Get(id)
	if err != nil {
		return nil, err
	}

	return zikr, nil
}

func (z zikrUsecase) GetAll() (zikrs []*domain.Zikr, err error) {
	zikrs, err = z.repo.GetAll()
	if err != nil {
		return nil, err
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

func (z zikrUsecase) Delete(id string) error {
	err := z.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

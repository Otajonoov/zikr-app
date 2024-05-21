package usecase

import (
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/domain/factory"
)

type zikrUsecase struct {
	BaseUseCase
	repo    domain.ZikrRepo
	factory factory.Factory
}

func NewZikrUsecase(repo domain.ZikrRepo) domain.ZikrUsecase {
	return &zikrUsecase{
		repo: repo,
	}
}

func (z *zikrUsecase) Create(zikr *domain.Zikr) error {
	z.beforeRequestForZikr(zikr)

	err := z.repo.Create(zikr)
	if err != nil {
		return err
	}

	return nil
}

//func (z *zikrUsecase) Get(guid string) (zikr *domain.Zikr, err error) {
//
//	//zikr, err = z.repo.Get(guid)
//	//if err != nil {
//	//	return &domain.Zikr{}, err
//	//}
//	//
//	return zikr, nil
//}

func (z *zikrUsecase) GetAll(guid string) (zikrs []domain.Zikr, err error) {
	zikrs, err = z.repo.GetAll(guid)
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
	//zikr.SetUpdatedAt(time.Now())
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

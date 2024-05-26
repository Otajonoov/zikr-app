package usecase

import (
	"log"
	"zikr-app/internal/zikr/domain"
)

type usersZikrUseCase struct {
	BaseUseCase
	repo domain.UsersZikrRepo
}

func NewCountUsecase(repo domain.UsersZikrRepo) *usersZikrUseCase {
	return &usersZikrUseCase{
		repo: repo,
	}
}

func (u *usersZikrUseCase) CountUpdate(count *domain.UsersZikr) error {
	if err := u.repo.CountUpdate(count); err != nil {
		log.Println("err: ", err)
		return err
	}
	return nil
}

func (u *usersZikrUseCase) Update(userId, zikrId string, isFavorite bool) error {
	if err := u.repo.Update(userId, zikrId, isFavorite); err != nil {
		log.Println("err: ", err)
		return err
	}
	return nil
}

func (u *usersZikrUseCase) Reyting(reyting *domain.Reyting) (*domain.ReytingResponse, error) {
	zikr, err := u.repo.Reyting(reyting)
	if err != nil {
		log.Println("err: ", err)
		return &domain.ReytingResponse{}, err
	}
	return zikr, nil
}

package usecase

import (
	"log"
	"zikr-app/internal/zikr/domain"
)

type countUseCase struct {
	BaseUseCase
	repo domain.CountRepo
}

func NewCountUsecase(repo domain.CountRepo) *countUseCase {
	return &countUseCase{
		repo: repo,
	}
}

func (c *countUseCase) CountUpdate(count *domain.UsersZikr) error {
	if err := c.repo.CountUpdate(count); err != nil {
		log.Println("err: ", err)
		return err
	}
	return nil
}

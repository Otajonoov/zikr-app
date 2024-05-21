package usecase

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
	"zikr-app/internal/zikr/domain"
)

type BaseUseCase struct{}

func (u *BaseUseCase) Error(msg string, err error) error {
	if len(strings.TrimSpace(msg)) != 0 {
		return fmt.Errorf("%v: %w", msg, err)
	}
	return err
}

func (u *BaseUseCase) beforeRequestForZikr(zikr *domain.Zikr) {
	if zikr.Guid == "" {
		zikr.Guid = uuid.New().String()
	}

}

func (u *BaseUseCase) beforeRequestForUser(user *domain.User) {
	if user.Guid == "" {
		user.Guid = uuid.New().String()
	}
}

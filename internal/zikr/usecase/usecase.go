package usecase

import (
	"fmt"
	"strings"
	"time"
	"zikr-app/internal/zikr/domain"
)

type BaseUseCase struct{}

func (u *BaseUseCase) Error(msg string, err error) error {
	if len(strings.TrimSpace(msg)) != 0 {
		return fmt.Errorf("%v: %w", msg, err)
	}
	return err
}

func (u *BaseUseCase) beforeRequest(zikr *domain.Zikr) {
	if zikr.GetCreatedAt().IsZero() {
		zikr.SetCreatedAt(time.Now().UTC())
	}

	if zikr.GetUpdatedAt().IsZero() {
		zikr.SetUpdatedAt(time.Now().UTC())
	}
}

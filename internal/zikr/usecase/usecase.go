package usecase

import (
	"fmt"
	"github.com/google/uuid"
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
	if zikr.GetGUID() == "" {
		zikr.SetGUID(uuid.New().String())
	}

	if zikr.GetCreatedAt().IsZero() {
		zikr.SetCreatedAt(time.Now().UTC())
	}

	if zikr.GetUpdatedAt().IsZero() {
		zikr.SetUpdatedAt(time.Now().UTC())
	}
}
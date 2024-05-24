package usecase

import "zikr-app/internal/zikr/domain"

type AppVersionUseCase struct {
	BaseUseCase
	repo domain.AppVersionRepository
}

func NewAppVersionUsecase(repo domain.AppVersionRepository) *AppVersionUseCase {
	return &AppVersionUseCase{
		repo: repo,
	}
}

func (a *AppVersionUseCase) GetAppVersion() (appVersion *domain.AppVersion, err error) {
	appVersion, err = a.repo.GetAppVersion()
	if err != nil {
		return nil, err
	}
	return appVersion, nil
}

func (a *AppVersionUseCase) Update(appVersion *domain.AppVersion) error {
	if err := a.repo.Update(appVersion); err != nil {
		return err
	}
	return nil
}

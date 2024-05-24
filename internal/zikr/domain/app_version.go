package domain

type AppVersion struct {
	AndroidVersion string
	IosVersion     string
	ForceUpdate    bool
}

type AppVersionUsecase interface {
	GetAppVersion() (appVersion *AppVersion, err error)
	Update(appVersion *AppVersion) error
}

type AppVersionRepository interface {
	GetAppVersion() (appVersion *AppVersion, err error)
	Update(appVersion *AppVersion) error
}

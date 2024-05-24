package adapter

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"zikr-app/internal/zikr/domain"
)

type appVersionRepo struct {
	db *pgxpool.Pool
}

func NewAppVersionRepo(db *pgxpool.Pool) *appVersionRepo {
	return &appVersionRepo{
		db: db,
	}
}

func (a appVersionRepo) GetAppVersion() (appVersion *domain.AppVersion, err error) {
	appVersion = &domain.AppVersion{}
	query := `
		SELECT
			android_version,
			ios_version,
			force_update
		FROM app_version`
	err = a.db.QueryRow(context.Background(), query).Scan(
		&appVersion.AndroidVersion,
		&appVersion.IosVersion,
		&appVersion.ForceUpdate,
	)
	if err != nil {
		log.Println("err: ", err)
		return &domain.AppVersion{}, err
	}
	return appVersion, nil
}

func (a appVersionRepo) Update(version *domain.AppVersion) error {
	query := `
		UPDATE app_version
		SET
		    android_version = $1,
			ios_version = $2,
			force_update = $3`
	_, err := a.db.Exec(context.Background(), query, version.AndroidVersion, version.IosVersion, version.ForceUpdate)
	if err != nil {
		log.Println("err: ", err)
		return err
	}
	return nil
}

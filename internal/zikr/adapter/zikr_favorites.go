package adapter

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
	"zikr-app/internal/zikr/domain"
)

type zikrFavoritesRepo struct {
	db      *pgxpool.Pool
	factory domain.Factory
}

type zikrModel struct {
	guid       string
	userGUID   string
	arabic     string
	uzbek      string
	pronounce  string
	count      int
	isFavorite bool
	createdAt  time.Time
	updatedAt  time.Time
}

func NewZikrFavoritesRepo(db *pgxpool.Pool) domain.ZikrFavoritesRepository {
	return &zikrFavoritesRepo{
		db: db,
	}
}

func (z zikrFavoritesRepo) FavoriteDua(userId, zikrId string) (bool, error) {
	query := `
				UPDATE zikr
				SET is_favorite=true
				WHERE user_guid=$1 and guid=$2
				`

	_, err := z.db.Exec(context.Background(), query, userId, zikrId)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (z zikrFavoritesRepo) UnFavoriteDua(userId, zikrId string) (bool, error) {
	query := `
				UPDATE zikr
				SET is_favorite=false
				WHERE user_guid=$1 and guid=$2
				`

	_, err := z.db.Exec(context.Background(), query, userId, zikrId)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (z zikrFavoritesRepo) GetAllFavorites(userId string) (zikrs []domain.Zikr, err error) {

	query := `
				SELECT
    				z.guid,
    				z.user_guid,
    				z.arabic,
    				z.uzbek,
    				z.pronounce,
    				z.count,
    				z.is_favorite,
    				z.created_at,
    				z.updated_at
				FROM
    				zikr z
				INNER JOIN
    				users u ON u.guid = z.user_guid
				WHERE
    				z.user_guid = $1 AND z.is_favorite = true;
					`

	rows, err := z.db.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var zikr zikrModel
		if err := rows.Scan(
			&zikr.guid,
			&zikr.userGUID,
			&zikr.arabic,
			&zikr.uzbek,
			&zikr.pronounce,
			&zikr.count,
			&zikr.isFavorite,
			&zikr.createdAt,
			&zikr.updatedAt,
		); err != nil {
			return nil, err
		}
		//zikrDomain := z.factory.ParseToDomain(zikr.guid, zikr.userGUID, zikr.arabic, zikr.uzbek, zikr.pronounce, zikr.count, zikr.isFavorite, zikr.createdAt, zikr.updatedAt)
		//zikrs = append(zikrs, *zikrDomain)
	}

	return zikrs, nil
}

func (z zikrFavoritesRepo) GetAllUnFavorites(userId string) (zikrs []domain.Zikr, err error) {

	query := `
				SELECT
    				z.guid,
    				z.user_guid,
    				z.arabic,
    				z.uzbek,
    				z.pronounce,
    				z.count,
    				z.is_favorite,
    				z.created_at,
    				z.updated_at
				FROM
    				zikr z
				INNER JOIN
    				users u ON u.guid = z.user_guid
				WHERE
    				z.user_guid = $1 AND z.is_favorite = false;
					`

	rows, err := z.db.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var zikr zikrModel
		if err := rows.Scan(
			&zikr.guid,
			&zikr.userGUID,
			&zikr.arabic,
			&zikr.uzbek,
			&zikr.pronounce,
			&zikr.count,
			&zikr.isFavorite,
			&zikr.createdAt,
			&zikr.updatedAt,
		); err != nil {
			return nil, err
		}
		//	zikrDomain := z.factory.ParseToDomain(zikr.guid, zikr.userGUID, zikr.arabic, zikr.uzbek, zikr.pronounce, zikr.count, zikr.isFavorite, zikr.createdAt, zikr.updatedAt)
		//zikrs = append(zikrs, *zikrDomain)
	}

	return zikrs, nil
}

package adapter

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
	"zikr-app/internal/zikr/domain"
)

type zikrFavoritesRepo struct {
	db      *pgxpool.Pool
	factory domain.ZikrFactory
}

func NewZikrFavoritesRepo(db *pgxpool.Pool, factory domain.ZikrFactory) domain.ZikrFavoritesRepository {
	return &zikrFavoritesRepo{
		db:      db,
		factory: factory,
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
	var guid string
	var userGUID string
	var arabic string
	var uzbek string
	var pronounce string
	var count int
	var isFavorite bool
	var createdAt *time.Time
	var updatedAt *time.Time

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
		var zikr domain.Zikr
		if err := rows.Scan(
			&guid,
			&userGUID,
			&arabic,
			&uzbek,
			&pronounce,
			&count,
			&isFavorite,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, err
		}
		log.Println("",
			guid,
			userGUID,
			arabic,
			uzbek,
			pronounce,
			count,
			isFavorite,
			createdAt,
			updatedAt)
		zikrs = append(zikrs, zikr)
	}
	return zikrs, nil
}

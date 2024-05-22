package adapter

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/domain/factory"
)

type zikrRepo struct {
	db      *pgxpool.Pool
	factory factory.Factory
}

func NewZikrRepo(db *pgxpool.Pool) domain.ZikrRepo {
	return &zikrRepo{
		db: db,
	}
}

func (z *zikrRepo) Create(zikr *domain.Zikr) error {
	query := `
		INSERT INTO zikr(
			guid,            
		    arabic, 
		    uzbek, 
		    pronounce
		)
		VALUES($1, $2, $3, $4)
	`

	_, err := z.db.Exec(context.Background(), query,
		zikr.Guid,
		zikr.Arabic,
		zikr.Uzbek,
		zikr.Pronounce,
	)
	if err != nil {
		return err
	}

	return nil
}

func (z *zikrRepo) Get(guid string) (zikr *domain.Zikr, err error) {
	var (
		id         string
		userGuid   string
		arabic     string
		uzbek      string
		pronounce  string
		count      int
		isFavorite bool
	)

	query := `
		SELECT
			z.guid,
			z.user_guid,
			z.arabic,
			z.uzbek,
			z.pronounce,
			z.count,
			z.is_favorite
		FROM zikr z
		WHERE z.guid = $1
    `

	err = z.db.QueryRow(context.Background(), query, guid).Scan(
		&id,
		&userGuid,
		&arabic,
		&uzbek,
		&pronounce,
		&count,
		&isFavorite,
	)

	if err != nil {
		return nil, err
	}

	//zikr = z.factory.ParseToDomainSpecial(id, userGuid, arabic, uzbek, pronounce, count, isFavorite)
	return zikr, nil
}

func (z *zikrRepo) GetAll(userGuid string) (zikrs []domain.Zikr, err error) {
	query := `
		SELECT       
            z.guid, 
            z.arabic,
            z.uzbek,
            z.pronounce,
    		COALESCE(uz.zikr_count, 0) AS count,
    		COALESCE(uz.isFavorite, false) AS isFavorite
        FROM zikr z 
        LEFT JOIN users_zikr uz on z.guid = uz.zikr_guid
		AND uz.user_guid = $1`

	rows, err := z.db.Query(context.Background(), query, userGuid)
	if err != nil {
		log.Println("err: ", err)
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var zikr domain.Zikr
		if err := rows.Scan(
			&zikr.Guid,
			&zikr.Arabic,
			&zikr.Uzbek,
			&zikr.Pronounce,
			&zikr.Count,
			&zikr.IsFavorite,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		zikrs = append(zikrs, zikr)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return zikrs, nil
}

func (z *zikrRepo) Update(zikr *domain.Zikr) error {
	query := `
		UPDATE zikr SET
		    arabic = $1,
			uzbek = $2,
			pronounce = $3
		WHERE guid = $4
   `

	_, err := z.db.Exec(context.Background(), query) //zikr.GetArabic(),
	//zikr.GetUzbek(),
	//zikr.GetPronounce(),
	//zikr.GetGuid(),

	if err != nil {
		return err
	}

	return nil
}

func (z *zikrRepo) Delete(guid string) error {
	query := `DELETE FROM zikr WHERE guid = $1`

	_, err := z.db.Exec(context.Background(), query, guid)
	if err != nil {
		return err
	}

	return nil
}

package adapter

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
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
		    coalesce(uz.count, 0) as count,
		    coalesce(uz.isFavorite, false) as isFavorite
		FROM zikr z 
		FULL JOIN users_zikr uz on z.guid = uz.zikr_guid AND uz.user_guid = $1;`

	rows, err := z.db.Query(context.Background(), query)
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

	log.Println("zikrs: ", zikrs)

	return zikrs, nil
}

func (z *zikrRepo) GetUserZikrByMail(email, username string) ([]domain.Zikr, error) {
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
		FROM zikr z
		INNER JOIN 
			users u ON u.guid = z.user_guid
		WHERE u.email = $1 AND u.unique_username = $2`

	rows, err := z.db.Query(context.Background(), query, email, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var zikrs []domain.Zikr
	for rows.Next() {
		var (
			guid       string
			userGUID   string
			arabic     string
			uzbek      string
			pronounce  string
			count      int
			isFavorite bool
			createdAt  time.Time
			updatedAt  time.Time
		)
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
		//zikrDomain := z.factory.ParseToDomain(guid, userGUID, arabic, uzbek, pronounce, count, isFavorite, createdAt, updatedAt)
		//zikrs = append(zikrs, *zikrDomain)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
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

func (z *zikrRepo) UpdateZikrCount(zikr *domain.Zikr) error {
	query := `
		UPDATE zikr 
		SET count = count + $1 
		WHERE guid = $2 AND user_guid = $3;
	`

	_, err := z.db.Exec(context.Background(), query) //zikr.GetCount(),
	//zikr.GetGuid(),
	//zikr.GetUserGUID(),

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

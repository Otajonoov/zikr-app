package adapter

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"zikr-app/internal/zikr/domain"
)

type zikrRepo struct {
	db      *pgxpool.Pool
	factory domain.ZikrFactory
}

func NewZikrRepo(db *pgxpool.Pool, factory domain.ZikrFactory) domain.ZikrRepo {
	return &zikrRepo{
		db:      db,
		factory: factory,
	}
}

type Zikr struct {
	guid       string
	userGuid   string
	arabic     string
	uzbek      string
	pronounce  string
	count      int
	isFavorite bool
}

func (z *zikrRepo) Create(zikr *domain.Zikr) error {
	query := `
		INSERT INTO zikr(
			guid,
		    user_guid,             
		    arabic, 
		    uzbek, 
		    pronounce,
		    count,             
		    is_favorite,
		    created_at,
		    updated_at
		)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);
	`

	_, err := z.db.Exec(context.Background(), query,
		zikr.GetGuid(),
		zikr.GetUserGUID(),
		zikr.GetArabic(),
		zikr.GetUzbek(),
		zikr.GetPronounce(),
		zikr.GetCount(),
		zikr.GetIsFavorite(),
		zikr.GetCreatedAt(),
		zikr.GetUpdatedAt(),
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

	zikr = z.factory.ParseToDomainSpecial(id, userGuid, arabic, uzbek, pronounce, count, isFavorite)
	return zikr, nil
}

func (z *zikrRepo) GetAll() (zikrs []domain.Zikr, err error) {
	zikr := Zikr{}

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
   `

	rows, err := z.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&zikr.guid,
			&zikr.userGuid,
			&zikr.arabic,
			&zikr.uzbek,
			&zikr.pronounce,
			&zikr.count,
			&zikr.isFavorite,
		); err != nil {
			return nil, err
		}

		newZikr := z.factory.ParseToDomainSpecial(zikr.guid, zikr.userGuid, zikr.arabic, zikr.uzbek, zikr.pronounce, zikr.count, zikr.isFavorite)
		zikrs = append(zikrs, *newZikr)
	}
	if err = rows.Err(); err != nil {
		return nil, err
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

	_, err := z.db.Exec(context.Background(), query,
		zikr.GetArabic(),
		zikr.GetUzbek(),
		zikr.GetPronounce(),
		zikr.GetGuid(),
	)

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

	_, err := z.db.Exec(context.Background(), query,
		zikr.GetCount(),
		zikr.GetGuid(),
		zikr.GetUserGUID(),
	)

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

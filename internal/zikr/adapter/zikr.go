package adapter

import (
	"context"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
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
	id         int
	userId     int
	arabic     string
	uzbek      string
	pronounce  string
	isFavorite bool
}

func (z *zikrRepo) Create(zikr *domain.Zikr) error {

	query := `
		INSERT INTO zikr(
		                 user_id,
		                 arabic, 
		                 uzbek, 
		                 pronounce,
		                 is_favorite,
		                 created_at,
		                 updated_at
		)
		VALUES($1, $2, $3, $4, $5, $6, $7)
	`
	// id removed
	// get fifth field

	_, err := z.db.Exec(context.Background(), query,
		zikr.GetUserId(),
		zikr.GetArabic(),
		zikr.GetUzbek(),
		zikr.GetPronounce(),
		zikr.GetIsFavourite(),
		zikr.GetCreatedAt(),
		zikr.GetUpdatedAt(),
	)
	if err != nil {
		log.Println("in repo", err.Error())
		return err
	}

	return nil
}

func (z *zikrRepo) Get(id int) (zikr *domain.Zikr, err error) {
	var zikrRes Zikr

	query := `
		SELECT 
		    z.id,
		    z.user_id,
			z.arabic, 
			z.uzbek, 
			z.pronounce,
			z.is_favorite
		FROM zikr z
		WHERE z.id=$1
   `

	err = z.db.QueryRow(context.Background(), query, id).Scan(
		&zikrRes.id,
		&zikrRes.userId,
		&zikrRes.arabic,
		&zikrRes.uzbek,
		&zikrRes.pronounce,
		&zikrRes.isFavorite,
	)

	if err != nil {
		log.Println("in repo", err.Error())
		return nil, err
	}

	newZikr := z.factory.ParseToController(zikrRes.id, zikrRes.userId, zikrRes.arabic, zikrRes.uzbek, zikrRes.pronounce, zikrRes.isFavorite)
	log.Println(newZikr)
	return newZikr, nil
}

func (z *zikrRepo) GetAll() (zikrs []domain.Zikr, err error) {
	var id int
	var userID int
	var arabic string
	var uzbek string
	var pronounce string
	var isFavorite bool
	var createdAt time.Time
	var updatedAt time.Time

	query := `
		SELECT	
		    id,
		    user_id,
		    arabic, 
		    uzbek, 
		    pronounce,
		    is_favorite,
		    created_at,
		    updated_at
		FROM zikr
    `

	rows, err := z.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&userID,
			&arabic,
			&uzbek,
			&pronounce,
			&isFavorite,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, err
		}

		newZikr := z.factory.ParseToDomain(id, userID, arabic, uzbek, pronounce, isFavorite, createdAt, updatedAt)

		zikrs = append(zikrs, *newZikr)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return zikrs, nil
}

func (z *zikrRepo) FavoriteDua(userId, zikrId int) (bool, error) {
	query := `
			UPDATE zikr 
			SET is_favorite=true 
			WHERE user_id=$2 and id=$3
			`

	_, err := z.db.Exec(context.Background(), query, userId, zikrId)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (z *zikrRepo) UnFavoriteDua(userId, zikrId int) (bool, error) {
	query := `
			UPDATE zikr 
			SET is_favorite=false 
			WHERE id=$2 and user_id=$3
			`

	_, err := z.db.Exec(context.Background(), query, zikrId, userId)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (z *zikrRepo) GetAllFavorites(userId int) (zikrs []domain.Zikr, err error) {
	var id int
	var userID int
	var arabic string
	var uzbek string
	var pronounce string
	var isFavorite bool
	var createdAt time.Time
	var updatedAt time.Time

	query := `
			SELECT 
				z.id,
				z.user_id,
				z.arabic,
				z.uzbek,
				z.pronounce,
				z.is_favorite,
				z.created_at,
				z.updated_at
			FROM zikr z
			INNER JOIN users u
				ON u.id = z.user_id
			WHERE z.user_id=$1 and z.is_favorite=true
			`

	rows, err := z.db.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&userID,
			&arabic,
			&uzbek,
			&pronounce,
			&isFavorite,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, err
		}
		newZikr := z.factory.ParseToDomain(id, userID, arabic, uzbek, pronounce, isFavorite, createdAt, updatedAt)
		zikrs = append(zikrs, *newZikr)
	}
	return zikrs, nil
}

func (z *zikrRepo) Update(zikr *domain.Zikr) error {
	query := `
		UPDATE zikr SET
		    arabic = $2, 
			uzbek = $3, 
			pronounce = $4
		WHERE id = $1
    `

	res, err := z.db.Exec(context.Background(), query,
		zikr.GetGUID(),
		zikr.GetArabic(),
		zikr.GetUzbek(),
		zikr.GetPronounce(),
	)

	if err != nil {
		return err
	}

	affected := res.RowsAffected()
	if affected == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (z *zikrRepo) Delete(id int) error {
	query := `DELETE FROM zikr WHERE id = $1`

	res, err := z.db.Exec(context.Background(), query, id)
	if err != nil {
		log.Println("err repo", err.Error())
		return err
	}
	affected := res.RowsAffected()
	if affected == 0 {
		return err
	}

	return nil
}

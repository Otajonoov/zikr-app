package adapter

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
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
	id        string
	arabic    string
	uzbek     string
	pronounce string
}

func (z *zikrRepo) Create(zikr *domain.Zikr) (id string, err error) {
	var resId string

	newUUID := uuid.New()
	query := `
		INSERT INTO zikr(id, arabic, uzbek, pronounce)
		VALUES($1, $2, $3, $4)
		RETURNING id
	`

	err = z.db.QueryRow(context.Background(), query, newUUID, zikr.GetArabic(), zikr.GetUzbek(), zikr.GetPronounce()).Scan(&resId)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return resId, nil
}

func (z *zikrRepo) Get(id string) (zikr *domain.Zikr, err error) {
	var zikrRes Zikr

	query := `
		SELECT 
			arabic, 
			uzbek, 
			pronounce
		FROM zikr
		WHERE id=$1
    `

	err = z.db.QueryRow(context.Background(), query, id).Scan(
		&zikrRes.arabic,
		&zikrRes.uzbek,
		&zikrRes.pronounce,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return &domain.Zikr{}, err
	} else if err != nil {
		log.Println(err.Error())
		return &domain.Zikr{}, err
	}

	newZikr := z.factory.ParseToController(zikrRes.arabic, zikrRes.uzbek, zikrRes.pronounce)

	return newZikr, nil
}

func (z *zikrRepo) GetAll() (zikrs []*domain.ZikrWithId, err error) {
	log.Println("Qwerty : ")

	query := `
		SELECT	
		    id,
		    arabic, 
		    uzbek, 
		    pronounce
		FROM zikr
    `

	rows, err := z.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var zikrRes Zikr
		if err := rows.Scan(
			&zikrRes.id,
			&zikrRes.arabic,
			&zikrRes.uzbek,
			&zikrRes.pronounce,
		); err != nil {
			return nil, err
		}
		zikrs = append(zikrs, &domain.ZikrWithId{
			Id:        zikrRes.id,
			Arabic:    zikrRes.arabic,
			Uzbek:     zikrRes.uzbek,
			Pronounce: zikrRes.pronounce,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Log zikrs
	log.Println(zikrs)
	return zikrs, nil
}

func (z *zikrRepo) Update(zikr *domain.ZikrWithId) error {
	query := `
		UPDATE zikr SET
		    arabic = $1, 
			uzbek = $2, 
			pronounce = $3
		WHERE id = $4
    `

	res, err := z.db.Exec(context.Background(), query,
		zikr.Arabic,
		zikr.Uzbek,
		zikr.Pronounce,
		zikr.Id,
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

func (z *zikrRepo) Delete(id string) error {
	query := `DELETE FROM zikr WHERE id = $1`

	res, err := z.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	affected := res.RowsAffected()
	if affected == 0 {
		return err
	}

	return nil
}

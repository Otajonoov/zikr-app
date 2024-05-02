package adapter

import (
	"context"
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
	guid      string
	arabic    string
	uzbek     string
	pronounce string
}

func (z *zikrRepo) Create(zikr *domain.Zikr) error {

	query := `
		INSERT INTO zikr(
			guid,
		    arabic, 
		    uzbek, 
		    pronounce,
		    created_at
		)
		VALUES($1, $2, $3, $4, $5)
	`

	_, err := z.db.Exec(context.Background(), query,
		zikr.GetGuid(),
		zikr.GetArabic(),
		zikr.GetUzbek(),
		zikr.GetPronounce(),
		zikr.GetCreatedAt(),
	)
	if err != nil {
		log.Println("in repo", err.Error())
		return err
	}

	return nil
}

func (z *zikrRepo) Get(guid string) (zikr *domain.Zikr, err error) {
	var (
		id        string
		arabic    string
		uzbek     string
		pronounce string
	)

	query := `
		SELECT
			guid,
			arabic,
			uzbek,
			pronounce
		FROM zikr
		WHERE guid = $1
    `

	err = z.db.QueryRow(context.Background(), query, guid).Scan(
		&id,
		&arabic,
		&uzbek,
		&pronounce,
	)

	if err != nil {
		log.Println("error while getting zikr: ", err.Error())
		return &domain.Zikr{}, err
	}

	zikr = z.factory.ParseToDomain(id, arabic, uzbek, pronounce)
	return zikr, nil
}

func (z *zikrRepo) GetAll() (zikrs []domain.Zikr, err error) {
	zikr := Zikr{}

	query := `
		SELECT
		    guid,
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
		if err := rows.Scan(
			&zikr.guid,
			&zikr.arabic,
			&zikr.uzbek,
			&zikr.pronounce,
		); err != nil {
			return nil, err
		}

		newZikr := z.factory.ParseToDomain(zikr.guid, zikr.arabic, zikr.uzbek, zikr.pronounce)
		zikrs = append(zikrs, *newZikr)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return zikrs, nil
}

//	func (z *zikrRepo) FavoriteDua(userId, zikrId int) (bool, error) {
//		query := `
//				UPDATE zikr
//				SET is_favorite=true
//				WHERE user_id=$2 and guid=$3
//				`
//
//		_, err := z.db.Exec(context.Background(), query, userId, zikrId)
//		if err != nil {
//			return false, err
//		}
//
//		return true, nil
//	}
//
//	func (z *zikrRepo) UnFavoriteDua(userId, zikrId int) (bool, error) {
//		query := `
//				UPDATE zikr
//				SET is_favorite=false
//				WHERE guid=$2 and user_id=$3
//				`
//
//		_, err := z.db.Exec(context.Background(), query, zikrId, userId)
//		if err != nil {
//			return false, err
//		}
//
//		return true, nil
//	}
//
//	func (z *zikrRepo) GetAllFavorites(userId int) (zikrs []domain.Zikr, err error) {
//		var guid int
//		var userID int
//		var arabic string
//		var uzbek string
//		var pronounce string
//		var isFavorite bool
//		var createdAt time.Time
//		var updatedAt time.Time
//
//		query := `
//				SELECT
//					z.guid,
//					z.user_id,
//					z.arabic,
//					z.uzbek,
//					z.pronounce,
//					z.is_favorite,
//					z.created_at,
//					z.updated_at
//				FROM zikr z
//				INNER JOIN users u
//					ON u.guid = z.user_id
//				WHERE z.user_id=$1 and z.is_favorite=true
//				`
//
//		rows, err := z.db.Query(context.Background(), query, userId)
//		if err != nil {
//			return nil, err
//		}
//		defer rows.Close()
//
//		for rows.Next() {
//			if err := rows.Scan(
//				&guid,
//				&userID,
//				&arabic,
//				&uzbek,
//				&pronounce,
//				&isFavorite,
//				&createdAt,
//				&updatedAt,
//			); err != nil {
//				return nil, err
//			}
//			newZikr := z.factory.ParseToDomain(guid, userID, arabic, uzbek, pronounce, isFavorite, createdAt, updatedAt)
//			zikrs = append(zikrs, *newZikr)
//		}
//		return zikrs, nil
//	}
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

func (z *zikrRepo) Delete(guid string) error {
	query := `DELETE FROM zikr WHERE guid = $1`

	_, err := z.db.Exec(context.Background(), query, guid)
	if err != nil {
		log.Println("err repo", err.Error())
		return err
	}

	return nil
}

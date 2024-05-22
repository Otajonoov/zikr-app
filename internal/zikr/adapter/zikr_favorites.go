package adapter

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/domain/factory"
)

type zikrFavoritesRepo struct {
	db      *pgxpool.Pool
	factory factory.Factory
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

func (z zikrFavoritesRepo) Update(userId, zikrId string, isFavorite bool) error {
	query := `
		UPDATE users_zikr
		SET isFavorite = $1
		WHERE user_guid = $2 and zikr_guid = $3`

	_, err := z.db.Exec(context.Background(), query, isFavorite, userId, zikrId)
	if err != nil {
		log.Println("err: ", err)
		return err
	}

	return nil
}

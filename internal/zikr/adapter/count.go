package adapter

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/domain/factory"
)

type countRepo struct {
	db      *pgxpool.Pool
	factory factory.Factory
}

func NewCountRepo(db *pgxpool.Pool) *countRepo {
	return &countRepo{
		db: db,
	}
}

func (u *countRepo) Create(count *domain.Count) error {
	query := `
		INSERT INTO users_zikr
			(guid, user_guid, zikr_guid, count)
		VALUES
			($1, $2, $3, $4)`

	_, err := u.db.Exec(context.Background(), query, count.Guid, count.UserGuid, count.ZikrGuid, count.Count)
	if err != nil {
		log.Printf("failed to execute insert: %v", err)
		return err
	}
	return nil
}

func (u *countRepo) CountUpdate(count *domain.Count) error {
	query := `
		UPDATE 
		    users_zikr 
		SET 
		    count = $1 	
		WHERE 
		    user_guid = $2 AND zikr_guid = $3`

	_, err := u.db.Exec(context.Background(), query, count.Count, count.UserGuid, count.ZikrGuid)
	if err != nil {
		log.Printf("failed to execute update: %v", err)
		return err
	}
	return nil
}

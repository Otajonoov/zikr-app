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

func (u *countRepo) CountUpdate(count *domain.UsersZikr) error {
	query := `
		UPDATE 
		    users_zikr 
		SET 
		    zikr_count = $1 	
		WHERE 
		    user_guid = $2 AND zikr_guid = $3`

	_, err := u.db.Exec(context.Background(), query, count.Count, count.UserGuid, count.ZikrGuid)
	if err != nil {
		log.Printf("failed to execute update: %v", err)
		return err
	}
	return nil
}

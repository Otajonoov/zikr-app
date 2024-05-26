package adapter

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/domain/factory"
)

type usersZikrRepo struct {
	db      *pgxpool.Pool
	factory factory.Factory
}

func NewCountRepo(db *pgxpool.Pool) *usersZikrRepo {
	return &usersZikrRepo{
		db: db,
	}
}

func (u *usersZikrRepo) CountUpdate(count *domain.UsersZikr) error {
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

func (u *usersZikrRepo) Update(userId, zikrId string, isFavorite bool) error {
	query := `
		UPDATE users_zikr
		SET isFavorite = $1
		WHERE user_guid = $2 and zikr_guid = $3`

	_, err := u.db.Exec(context.Background(), query, isFavorite, userId, zikrId)
	if err != nil {
		log.Println("err: ", err)
		return err
	}

	return nil
}

func (u *usersZikrRepo) Reyting(request *domain.Reyting) (*domain.ReytingResponse, error) {

	query := `
		SELECT
			u.guid,
			u.username,
			uz.zikr_count
		FROM users u
		JOIN users_zikr uz ON uz.user_guid = u.guid
		WHERE uz.zikr_guid = $1 AND uz.zikr_count > 0
		ORDER BY uz.zikr_count DESC
		LIMIT $2 OFFSET $3
	`

	offset := (request.Page - 1) * request.Limit
	rows, err := u.db.Query(context.Background(), query, request.ZikrGuid, request.Limit, offset)
	if err != nil {
		log.Println("Error querying request: ", err)
		return nil, err
	}
	defer rows.Close()

	reytings := domain.ReytingResponse{}
	for rows.Next() {
		var r domain.ReytingInfo
		if err := rows.Scan(
			&r.UserGuid,
			&r.Username,
			&r.ZikrCount,
		); err != nil {
			log.Println("Error scanning row: ", err)
			return nil, err
		}
		reytings.Reyting = append(reytings.Reyting, r)
	}

	if err = rows.Err(); err != nil {
		log.Println("Row iteration error: ", err)
		return nil, err
	}

	return &reytings, nil
}

package adapter

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/domain/factory"
)

type authRepo struct {
	db      *pgxpool.Pool
	factory factory.Factory
}

func NewAuthRepo(db *pgxpool.Pool) *authRepo {
	return &authRepo{
		db: db,
	}
}

func (u authRepo) CreateUser(ctx context.Context, user *domain.User) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var guid string
	query := `
		INSERT INTO users(
			guid,
			email,
		    username
		) VALUES ($1, $2, $3)
		RETURNING guid
	`

	err := u.db.QueryRow(ctx, query, user.Guid, user.Email, user.Username).Scan(&guid)
	if err != nil {
		log.Println("err: ", err)
		return "", err
	}

	return guid, nil
}

func (u authRepo) GetUserInfo(ctx context.Context, email string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var guid string

	query := `
		SELECT
			u.guid
		FROM users u
		WHERE u.email = $1
	`
	err := u.db.QueryRow(ctx, query, email).Scan(&guid)
	if err != nil {
		log.Println("err: ", err)
		return "", err
	}

	return guid, nil
}

func (u authRepo) GetAllZikrGuid(ctx context.Context) ([]domain.Guid, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	guids := []domain.Guid{}
	query := `
		SELECT
			z.guid
		FROM zikr z
	`

	rows, err := u.db.Query(ctx, query)
	if err != nil {
		log.Println("err: ", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var guid domain.Guid
		if err := rows.Scan(&guid.Guid); err != nil {
			log.Println("err: ", err)
			return nil, err
		}
		guids = append(guids, guid)
	}

	return guids, nil
}

func (u authRepo) CreateZikrsForUser(ctx context.Context, userGuid, zikrGuid string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO users_zikr(
			user_guid,
			zikr_guid     
		) VALUES ($1, $2)`
	_, err := u.db.Exec(ctx, query, userGuid, zikrGuid)
	if err != nil {
		log.Println("err: ", err)
		return err
	}

	return nil
}

package adapter

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
	"zikr-app/internal/zikr/domain"
)

type authRepo struct {
	db *pgxpool.Pool
}

func NewAuthRepo(db *pgxpool.Pool) *authRepo {
	return &authRepo{
		db: db,
	}
}

func (u authRepo) CreateUser(ctx context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO users(
			guid,
			email,
		    unique_username
		) VALUES ($1, $2, $3)
	`

	_, err := u.db.Exec(context.Background(), query, user.Guid, user.Email, user.UniqueUsername)
	if err != nil {
		return err
	}

	return nil
}

func (u authRepo) UserExistsByMail(ctx context.Context, mail string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var exists bool
	query := `SELECT EXISTS (
				SELECT 1
				FROM users u
				WHERE u.email = $1 
		)
	`

	err := u.db.QueryRow(ctx, query, mail).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (u authRepo) UserExistsByUsername(ctx context.Context, username string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var exists bool
	query := `SELECT EXISTS (
				SELECT 1
				FROM users u
				WHERE u.unique_username = $1 
		)
	`

	err := u.db.QueryRow(ctx, query, username).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

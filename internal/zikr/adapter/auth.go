package adapter

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
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
	query := `
		INSERT INTO users(
			fio,
			phone,
		    uniqe_username,
			password
		) VALUES ($1, $2, $3, $4)
	`
	_, err := u.db.Exec(context.Background(), query, user.FIO, user.PhoneNumber, user.UniqeUsername, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (u authRepo) GetUser(ctx context.Context, username string) (*domain.User, error) {
	query := `SELECT password 
                FROM users
                WHERE uniqe_username = $1`

	req := domain.User{}

	err := u.db.QueryRow(ctx, query, username).Scan(&req.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found for username: %s", username)
		}
		return nil, err
	}

	return &req, nil
}

func (u authRepo) FindOneByUsername(ctx context.Context, userName string) (*domain.User, error) {
	query := `SELECT u.fio,
    				 u.phone,
    				 u.uniqe_username,
    				 u.password
			  FROM users u
			  WHERE u.uniqe_username = $1;`

	req := domain.User{}

	if err := u.db.QueryRow(ctx, query, userName).Scan(
		&req.FIO,
		&req.PhoneNumber,
		&req.UniqeUsername,
		&req.Password,
	); err != nil {
		return nil, err
	}

	return &req, nil
}

func (u authRepo) UserExists(ctx context.Context, username string) (bool, error) {
	var exists bool
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM users
			WHERE uniqe_username = $1
		)
	`
	err := u.db.QueryRow(ctx, query, username).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

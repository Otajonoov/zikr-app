package adapter

import (
	"context"
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

func (u authRepo) GetUser(ctx context.Context, phone string) (*domain.User, error) {

	req := domain.User{}

	if err := u.db.QueryRow(ctx, "SELECT phone, password FROM users WHERE phone = $1", phone).Scan(
		&req.PhoneNumber,
		&req.Password,
	); err != nil {
		return nil, err
	}

	return &req, nil
}

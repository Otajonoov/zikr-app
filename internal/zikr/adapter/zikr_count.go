package adapter

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/context"
	"time"
	"zikr-app/internal/zikr/domain"
)

type zikrCountRepo struct {
	db *pgxpool.Pool
}

func NewZikrCountRepo(db *pgxpool.Pool) domain.ZikrCountRepository {
	return &zikrCountRepo{db: db}
}

func (z zikrCountRepo) CreateZikrCount(ctx context.Context, zikr *domain.ZikrCount) error {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	query := `
			INSERT INTO zikr_count(
			    user_id,
			    zikr_id,
			    count
				) VALUES ($1,$2,$3)
			`

	_, err := z.db.Exec(context.Background(), query, zikr.UserId, zikr.ZikrId, zikr.Count)
	if err != nil {
		return err
	}
	return nil
}

func (z zikrCountRepo) GetAllUserCount(ctx context.Context, userId int) (map[string]int, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	query := `
		SELECT z.zikr_id, z.count
		FROM zikr_count z
		WHERE z.user_id = $1
		`

	rows, err := z.db.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	zikrCount := map[string]int{}
	for rows.Next() {
		var zikrID string
		var count int
		if err := rows.Scan(&zikrID, &count); err != nil {
			return nil, err
		}
		zikrCount[zikrID] = count
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return zikrCount, nil
}

func (z zikrCountRepo) UpdateUserCount(ctx context.Context, userId int, count int) error {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	query := `
			UPDATE zikr_count	
			SET count = count + $1
			WHERE user_id = $2
			`

	_, err := z.db.Exec(context.Background(), query, count, userId)
	if err != nil {
		return err
	}
	return nil
}

func (z zikrCountRepo) DeleteCount(ctx context.Context, userId int) error {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	query := `
			UPDATE zikr_count
			SET count = 0
			WHERE user_id = $1
			`

	_, err := z.db.Exec(context.Background(), query, userId)
	if err != nil {
		return err
	}
	return nil
}

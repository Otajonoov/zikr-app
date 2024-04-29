package domain

import "context"

type ZikrCount struct {
	UserId int
	ZikrId int
	Count  int
}

type ZikrCountUsecase interface {
	CreateCount(ctx context.Context, count *ZikrCount) error
	GetUserCounts(ctx context.Context, userId int) (map[string]int, error)
	PatchCount(ctx context.Context, userId int, count int) error
	ResetCount(ctx context.Context, userId int) error
}

type ZikrCountRepository interface {
	CreateZikrCount(ctx context.Context, user *ZikrCount) error
	GetAllUserCount(ctx context.Context, userId int) (map[string]int, error)
	UpdateUserCount(ctx context.Context, userId int, count int) error
	DeleteCount(ctx context.Context, userId int) error
}

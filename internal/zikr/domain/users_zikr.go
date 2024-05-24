package domain

type UsersZikr struct {
	Guid       string
	UserGuid   string
	ZikrGuid   string
	Count      int64
	IsFavorite bool
}

type CountRepo interface {
	CountUpdate(count *UsersZikr) error
}

type CountUsecase interface {
	CountUpdate(count *UsersZikr) error
}

type ZikrFavoritesRepository interface {
	Update(userId, zikrId string, isFavorite bool) error
}

type ZikrFavoritesUsecase interface {
	Update(userId, zikrId string, isFavorite bool) error
}

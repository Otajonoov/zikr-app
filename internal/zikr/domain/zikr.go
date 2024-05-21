package domain

type Zikr struct {
	Guid       string
	UserEmail  string
	Arabic     string
	Uzbek      string
	Pronounce  string
	Count      int
	IsFavorite bool
}

type ZikrRepo interface {
	Create(zikr *Zikr) error
	Get(guid string) (zikr *Zikr, err error)
	GetAll() (zikrs []Zikr, err error)
	GetUserZikrByMail(email, username string) (zikrs []Zikr, err error)
	Update(zikr *Zikr) error
	UpdateZikrCount(zikr *Zikr) error
	Delete(guid string) error
}

type ZikrFavoritesRepository interface {
	FavoriteDua(userId, zikrId string) (bool, error)
	UnFavoriteDua(userId, zikrId string) (bool, error)
	GetAllFavorites(userId string) (zikrs []Zikr, err error)
	GetAllUnFavorites(userId string) (zikrs []Zikr, err error)
}

type ZikrUsecase interface {
	Create(zikr *Zikr) error
	Get(guid string) (zikr *Zikr, err error)
	GetAll() (zikrs []Zikr, err error)
	Update(zikr *Zikr) error
	UpdateZikrCount(zikr *Zikr) error
	Delete(guid string) error
}

type ZikrFavoritesUsecase interface {
	FavoriteDua(userId, zikrId string) (bool, error)
	UnFavoriteDua(userId, zikrId string) (bool, error)
	GetAllFavorites(userId string) (zikrs []Zikr, err error)
	GetAllUnFavorites(userId string) (zikrs []Zikr, err error)
}

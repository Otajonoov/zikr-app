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
	GetAll(guid string) (zikrs []Zikr, err error)
	Update(zikr *Zikr) error
	Delete(guid string) error
}

type ZikrUsecase interface {
	Create(zikr *Zikr) error
	GetAll(guid string) (zikrs []Zikr, err error)
	Update(zikr *Zikr) error
	Delete(guid string) error
}

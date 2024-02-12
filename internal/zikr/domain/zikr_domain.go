package domain

type Zikr struct {
	arabic    string
	uzbek     string
	pronounce string
}

type ZikrWithId struct {
	Id        string
	Arabic    string
	Uzbek     string
	Pronounce string
}

type Zikrs struct {
	Zikr []*ZikrWithId
}

func (z *Zikr) GetArabic() string {
	return z.arabic
}

func (z *Zikr) SetArabic(arabic string) {
	z.arabic = arabic
}

func (z *Zikr) GetUzbek() string {
	return z.uzbek
}

func (z *Zikr) SetUzbek(uzbek string) {
	z.uzbek = uzbek
}

func (z *Zikr) GetPronounce() string {
	return z.pronounce
}

func (z *Zikr) SetPronounce(pronounce string) {
	z.pronounce = pronounce
}

type ZikrRepo interface {
	Create(zikr *Zikr) (id string, err error)
	Get(id string) (zikr *Zikr, err error)
	GetAll() (zikrs []*ZikrWithId, err error)
	Update(zikr *Zikr) error
	Delete(id string) error
}

type ZikrUsecase interface {
	Create(zikr *Zikr) (id string, err error)
	Get(id string) (zikr *Zikr, err error)
	GetAll() (zikrs []*ZikrWithId, err error)
	Update(zikr *Zikr) error
	Delete(id string) error
}

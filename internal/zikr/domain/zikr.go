package domain

import (
	"time"
)

type Zikr struct {
	guid       string
	userGUID   string
	arabic     string
	uzbek      string
	pronounce  string
	count      int
	isFavorite bool
	createdAt  time.Time
	updatedAt  time.Time
}

type Zikrs struct {
	Zikr []*Zikr
}

func (z *Zikr) GetGuid() string {
	return z.guid
}

func (z *Zikr) SetGuid(id string) {
	z.guid = id
}

func (z *Zikr) GetUserGUID() string {
	return z.userGUID
}

func (z *Zikr) SetUserGUID(id string) {
	z.userGUID = id
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

func (z *Zikr) GetIsFavorite() bool {
	return z.isFavorite
}

func (z *Zikr) SetIsFavorite(isFavorite bool) {
	z.isFavorite = isFavorite
}

func (z *Zikr) GetCount() int {
	return z.count
}

func (z *Zikr) SetCount(count int) {
	z.count = count
}

func (z *Zikr) GetCreatedAt() time.Time {
	return z.createdAt
}

func (z *Zikr) SetCreatedAt(createdAt time.Time) {
	z.createdAt = createdAt
}

func (z *Zikr) GetUpdatedAt() time.Time {
	return z.updatedAt
}

func (z *Zikr) SetUpdatedAt(updatedAt time.Time) {
	z.updatedAt = updatedAt
}

type ZikrRepo interface {
	Create(zikr *Zikr) error
	Get(guid string) (zikr *Zikr, err error)
	GetAll() (zikrs []Zikr, err error)
	Update(zikr *Zikr) error
	UpdateZikrCount(zikr *Zikr) error
	Delete(guid string) error
}

type ZikrFavoritesRepository interface {
	FavoriteDua(userId, zikrId string) (bool, error)
	UnFavoriteDua(userId, zikrId string) (bool, error)
	GetAllFavorites(userId string) (zikrs []Zikr, err error)
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
}

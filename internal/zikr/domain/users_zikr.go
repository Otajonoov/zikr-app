package domain

type UsersZikr struct {
	Guid       string
	UserGuid   string
	ZikrGuid   string
	Count      int64
	IsFavorite bool
}

type ReytingInfo struct {
	UserGuid  string
	Username  string
	ZikrCount int64
}

type ReytingResponse struct {
	Reyting []ReytingInfo
}

type Reyting struct {
	Limit    int64
	Page     int64
	ZikrGuid string
}

type UsersZikrRepo interface {
	CountUpdate(count *UsersZikr) error
	Update(userId, zikrId string, isFavorite bool) error
	Reyting(reyting *Reyting) (reytings *ReytingResponse, err error)
}

type UsersZikrUseCase interface {
	CountUpdate(count *UsersZikr) error
	Update(userId, zikrId string, isFavorite bool) error
	Reyting(reyting *Reyting) (reytings *ReytingResponse, err error)
}

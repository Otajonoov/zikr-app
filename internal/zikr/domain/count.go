package domain

type Count struct {
	Guid     string
	UserGuid string
	ZikrGuid string
	Count    int64
}

type CountRepo interface {
	Create(count *Count) error
	CountUpdate(count *Count) error
}

type CountUsecase interface {
	Create(count *Count) error
	CountUpdate(count *Count) error
}

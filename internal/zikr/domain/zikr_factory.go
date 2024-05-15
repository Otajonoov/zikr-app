package domain

import (
	"time"
)

type ZikrFactory struct{}

func NewZikrFactory() ZikrFactory {
	return ZikrFactory{}
}

func (z *ZikrFactory) ParseToDomain(id, userGUID, arabic, uzbek, pronounce string, count int, isFavorite bool, createdAt, updatedAt time.Time) *Zikr {
	return &Zikr{
		guid:       id,
		userGUID:   userGUID,
		arabic:     arabic,
		uzbek:      uzbek,
		pronounce:  pronounce,
		count:      count,
		isFavorite: isFavorite,
		createdAt:  createdAt,
		updatedAt:  updatedAt,
	}
}

func (z *ZikrFactory) ParseToDomainSpecial(id, userGuid, arabic, uzbek, pronounce string, count int, isFavorite bool) *Zikr {
	return &Zikr{
		guid:       id,
		userGUID:   userGuid,
		arabic:     arabic,
		uzbek:      uzbek,
		pronounce:  pronounce,
		count:      count,
		isFavorite: isFavorite,
	}
}

func (z *ZikrFactory) ParseToDomainHandler(id, arabic, uzbek, pronounce string) *Zikr {
	return &Zikr{
		guid:      id,
		arabic:    arabic,
		uzbek:     uzbek,
		pronounce: pronounce,
	}
}

func (z *ZikrFactory) ParseToControllerForCreate(arabik, uzbek, pronounce string) *Zikr {
	return &Zikr{
		arabic:    arabik,
		uzbek:     uzbek,
		pronounce: pronounce,
	}
}

func (z *ZikrFactory) ParseToDomainToPatch(guid, userGuid string, count int) *Zikr {
	return &Zikr{
		guid:     guid,
		userGUID: userGuid,
		count:    count,
	}
}

package domain

import "time"

type ZikrFactory struct{}

func NewZikrFactory() ZikrFactory {
	return ZikrFactory{}
}

func (z *ZikrFactory) ParseToDomain(id, userId int, arabic, uzbek, pronounce string, isFav bool, created, updated time.Time) *Zikr {
	return &Zikr{
		id:         id,
		userId:     userId,
		arabic:     arabic,
		uzbek:      uzbek,
		pronounce:  pronounce,
		isFavorite: isFav,
		createdAt:  created,
		updatedAt:  updated,
	}
}

func (z *ZikrFactory) ParseToDomain2(id, userId int, arabic, uzbek, pronounce string, isFav bool) *Zikr {
	return &Zikr{
		id:         id,
		userId:     userId,
		arabic:     arabic,
		uzbek:      uzbek,
		pronounce:  pronounce,
		isFavorite: isFav,
	}
}

func (z *ZikrFactory) ParseToDomainHandler(id int, arabic, uzbek, pronounce string, isFav bool) *Zikr {
	return &Zikr{}
}

func (z *ZikrFactory) ParseToControllerForCreate(userId int, arabic, uzbek, pronounce string, isFavorites bool) *Zikr {
	return &Zikr{
		userId:     userId,
		arabic:     arabic,
		uzbek:      uzbek,
		pronounce:  pronounce,
		isFavorite: isFavorites,
	}
}

func (z *ZikrFactory) ParseToController(id, userId int, arabic, uzbek, pronounce string, isFavorites bool) *Zikr {
	return &Zikr{
		id:         id,
		userId:     userId,
		arabic:     arabic,
		uzbek:      uzbek,
		pronounce:  pronounce,
		isFavorite: isFavorites,
	}
}

//func (z *ZikrFactory) ParseToDomainArray(rows *pgx.Rows) *Zikrs {
//	var result Zikrs
//	result.Zikr = make([]*Zikr, 0)
//	log.Println(" : ", rows)
//
//	for rows.Next() {
//		var zikr Zikr
//		err := rows.Scan(
//			&zikr.arabic,
//			&zikr.uzbek,
//			&zikr.pronounce,
//		)
//		if err != nil {
//			return nil
//		}
//
//		result.Zikr = append(result.Zikr, &zikr)
//	}
//
//	return &result
//}

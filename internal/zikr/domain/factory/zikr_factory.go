package factory

import "zikr-app/internal/zikr/domain"

type Factory struct{}

//
//func (z *Factory) ParseToDomainSpecial(id, userGuid, arabic, uzbek, pronounce string, count int, isFavorite bool) *Zikr {
//	return &Zikr{
//		guid:       id,
//		userEmail:  userGuid,
//		arabic:     arabic,
//		uzbek:      uzbek,
//		pronounce:  pronounce,
//		count:      count,
//		isFavorite: isFavorite,
//	}
//}

//func (z *Factory) ParseToDomainHandler(id, arabic, uzbek, pronounce string) *Zikr {
//	return &Zikr{
//		guid:      id,
//		arabic:    arabic,
//		uzbek:     uzbek,
//		pronounce: pronounce,
//	}
//}

func (z *Factory) ParseToControllerForCreate(arabik, uzbek, pronounce string) *domain.Zikr {
	return &domain.Zikr{
		Arabic:    arabik,
		Uzbek:     uzbek,
		Pronounce: pronounce,
	}
}

func (z *Factory) ParseToDomainToUpdate(guid, arabic, uzbek, pronounce string) *domain.Zikr {
	return &domain.Zikr{
		Guid:      guid,
		Arabic:    arabic,
		Uzbek:     uzbek,
		Pronounce: pronounce,
	}
}
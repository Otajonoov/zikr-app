package factory

import "zikr-app/internal/zikr/domain"

type Factory struct{}

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

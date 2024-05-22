package factory

import "zikr-app/internal/zikr/domain"

func (z *Factory) ParseToDomainForCount(userGuid, zikrGuid string, count int64) *domain.UsersZikr {
	return &domain.UsersZikr{
		UserGuid: userGuid,
		ZikrGuid: zikrGuid,
		Count:    count,
	}
}

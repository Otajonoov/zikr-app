package factory

import "zikr-app/internal/zikr/domain"

func (z *Factory) ParseToDomainForCount(userGuid, zikrGuid string, count int64) *domain.UsersZikr {
	return &domain.UsersZikr{
		UserGuid: userGuid,
		ZikrGuid: zikrGuid,
		Count:    count,
	}
}

func (z *Factory) ParseToDomainForAppVersion(androidVersion, iosVersion string, forceUpdate bool) *domain.AppVersion {
	return &domain.AppVersion{
		AndroidVersion: androidVersion,
		IosVersion:     iosVersion,
		ForceUpdate:    forceUpdate,
	}
}

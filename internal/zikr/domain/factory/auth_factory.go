package factory

import "zikr-app/internal/zikr/domain"

func (a *Factory) ParseToDomainForAuth(email, username string) *domain.User {
	return &domain.User{
		Email:    email,
		Username: username,
	}
}

package domain

func (a *Factory) ParseToDomainForAuth(email, username string) *User {
	return &User{
		Email:    email,
		Username: username,
	}
}

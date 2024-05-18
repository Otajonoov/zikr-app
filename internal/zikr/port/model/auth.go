package model

type User struct {
	Email          string `json:"email"`
	UniqueUsername string `json:"unique_username"`
}

type UserGuid struct {
	Guid string
}

package model

type User struct {
	FIO           string `json:"fio"`
	UniqeUsername string `json:"uniqeUsername"`
	Password      string `json:"password"`
	PhoneNumber   string `json:"phone_number" example:"998999999999"`
}

type SignIn struct {
	UserName string
	Password string
}

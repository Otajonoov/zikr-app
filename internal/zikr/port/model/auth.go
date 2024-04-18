package model

type User struct {
	FIO           string
	UniqeUsername string
	Password      string
	PhoneNumber   string
} // removed json tags

type SignIn struct {
	UserName string
	Password string
}

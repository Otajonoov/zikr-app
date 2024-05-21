package model

type User struct {
	Guid     string `json:"guid"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserGuid struct {
	Guid string `json:"guid"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

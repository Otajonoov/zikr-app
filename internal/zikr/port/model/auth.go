package model

type UserGuid struct {
	Guid string `json:"guid"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

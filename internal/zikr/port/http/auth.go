package http

import (
	"context"
	"encoding/json"
	"net/http"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/port/model"
)

type AuthHandler struct {
	usecase domain.AuthUsecase
}

func NewAuthHandler(u domain.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: u}
}

func (u *AuthHandler) UserRegister(w http.ResponseWriter, r *http.Request) {
	var req model.User

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := u.usecase.CreateUser(context.Background(), &domain.User{
		Guid:           req.Guid,
		Email:          req.Email,
		UniqueUsername: req.UniqueUsername,
	}); err != nil {
		http.Error(w, "failed to create user: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("user created"))
	w.WriteHeader(http.StatusCreated)
}

func (u *AuthHandler) CheckUserRegister(w http.ResponseWriter, r *http.Request) {
	var req domain.UserLoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := u.usecase.CheckUser(context.Background(), req)
	if err != nil {
		http.Error(w, "bad credentials: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

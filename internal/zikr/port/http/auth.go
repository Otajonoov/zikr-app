package http

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"zikr-app/internal/pkg/jwt"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/port/model"
)

type AuthHandler struct {
	usecase domain.AuthUsecase
}

func NewAuthHandler(u domain.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: u}
}

// @Summary 	Sign-Up user
// @Description This api can Sign-Up new user
// @Tags 		User
// @Accept 		json
// @Produce 	json
// @Param body  body model.User true "body"
// @Failure 404 string Error response
// @Router /v1/sign-up [post]
func (u *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req model.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := u.usecase.CreateUser(context.Background(), &domain.User{
		FIO:           req.FIO,
		UniqeUsername: req.UniqeUsername,
		PhoneNumber:   req.PhoneNumber,
		Password:      req.Password,
	}); err != nil {
		http.Error(w, "failed to create user: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Signed up"))
	w.WriteHeader(http.StatusCreated)
}

// @Summary 	Sign-In user
// @Description This api can Sign-In user
// @Tags 		User
// @Accept 		json
// @Produce 	json
// @Param body  body model.SignIn true "body"
// @Failure 404 string Error response
// @Router /v1/sign-in [post]
func (u *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req model.SignIn
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	ok, err := u.usecase.CheckUser(context.Background(), req.UserName, req.Password)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusInternalServerError)
		return
	}

	if ok {
		token, err := jwt.CreateToken(req.UserName)
		if err != nil {
			http.Error(w, "username not found: "+err.Error(), http.StatusInternalServerError)
			return
		}
		response := map[string]string{
			"access_token": token,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
	}
}

func (u *AuthHandler) GetUserByUserName(w http.ResponseWriter, r *http.Request) {
	userName := chi.URLParam(r, "Username")

	user, err := u.usecase.GetByUserName(context.Background(), userName)
	if err != nil {
		http.Error(w, "user not found: "+err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

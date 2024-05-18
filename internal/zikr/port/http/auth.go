package http

import (
	"context"
	"encoding/json"
	"net/http"
	"zikr-app/internal/zikr/domain"
)

type AuthHandler struct {
	usecase domain.AuthUsecase
}

func NewAuthHandler(u domain.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: u}
}

// @Summary 	Get or Create user
// @Description register-user
// @Tags 		User
// @Accept 		json
// @Produce 	json
// @Param body  body domain.UserLoginRequest true "account info"
// @Success   200 {object} model.UserGuid "Successful response"
// @Failure 404 string Error response
// @Router /user/check-or-register [post]
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

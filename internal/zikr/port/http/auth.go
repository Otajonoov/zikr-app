package http

import (
	"context"
	"encoding/json"
	"net/http"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/domain/factory"
	"zikr-app/internal/zikr/port/model"
)

type AuthHandler struct {
	usecase domain.AuthUsecase
	factory factory.Factory
}

func NewAuthHandler(u domain.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: u}
}

// @Summary 	Get or Create user
// @Description register-user
// @Tags 		user
// @Accept 		json
// @Produce 	json
// @Param body  body model.UserLoginRequest true "account info"
// @Success   200 {object} model.UserGuid "Successful response"
// @Failure 404 string Error response
// @Router /user [post]
func (u *AuthHandler) CheckUserRegister(w http.ResponseWriter, r *http.Request) {
	var req model.UserLoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	user := u.factory.ParseToDomainForAuth(req.Email, req.Username)
	guid, err := u.usecase.GetOrCreateUser(context.Background(), user)
	if err != nil {
		http.Error(w, "bad credentials: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.UserGuid{Guid: guid})
}

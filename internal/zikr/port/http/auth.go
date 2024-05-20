package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/port/model"
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
// @Param body  body model.UserLoginRequest true "account info"
// @Success   200 {object} model.UserGuid "Successful response"
// @Failure 404 string Error response
// @Router /user/check-or-register [post]
func (u *AuthHandler) CheckUserRegister(w http.ResponseWriter, r *http.Request) {
	var req model.UserLoginRequest

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

	lines := strings.Split(user, ", ")

	formattedUser := strings.Join(lines, "\n")

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(formattedUser))
}

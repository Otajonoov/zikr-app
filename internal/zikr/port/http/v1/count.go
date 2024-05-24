package v1

import (
	"encoding/json"
	"net/http"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/domain/factory"
	"zikr-app/internal/zikr/port/http/model"
)

type CountHandler struct {
	usecase domain.CountUsecase
	factory factory.Factory
}

func NewCountHandler(u domain.CountUsecase) *CountHandler {
	return &CountHandler{usecase: u}
}

// @Summary 	Update zikr count
// @Description This API updates zikr count
// @Tags 		zikr-count
// @Accept 		json
// @Produce 	json
// @Param body  body model.Count true "account info"
// @Success   200 {object} model.Response "Successful response"
// @Failure 404 string Error response
// @Router /count [patch]
func (c *CountHandler) Count(w http.ResponseWriter, r *http.Request) {
	var req model.Count

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	userInfo := c.factory.ParseToDomainForCount(req.UserGuid, req.ZikrGuid, req.Count)
	err = c.usecase.CountUpdate(userInfo)
	if err != nil {
		http.Error(w, "bad credentials: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Response{Result: "ok"})
}

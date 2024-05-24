package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/domain/factory"
	"zikr-app/internal/zikr/port/http/model"
)

type zikrFavoriteHandler struct {
	service domain.ZikrFavoritesRepository
	factory factory.Factory
}

func NewZikrFavoriteHandler(service domain.ZikrFavoritesUsecase) *zikrFavoriteHandler {
	return &zikrFavoriteHandler{
		service: service,
	}
}

// @Summary Mark zikr as favorite
// @Description This API marks zikr as favorite
// @Tags 	zikr-favorite
// @Accept 	json
// @Produce json
// @Param 	body body model.IsFavorite true "body"
// @Success 200 {string} string "updated to favorite"
// @Failure 400 {string} string "invalid request body"
// @Failure 404 {string} string "could not update"
// @Router /favorite [patch]
func (z *zikrFavoriteHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req model.IsFavorite
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("err favorite: ", err)
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err = z.service.Update(req.UserGuId, req.ZikrGuid, req.IsFav); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

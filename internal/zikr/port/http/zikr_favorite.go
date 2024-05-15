package http

import (
	"encoding/json"
	"net/http"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/port/model"
)

type zikrFavoriteHandler struct {
	service domain.ZikrFavoritesRepository
	factory domain.ZikrFactory
}

func NewZikrFavoriteHandler(service domain.ZikrFavoritesUsecase) *zikrFavoriteHandler {
	return &zikrFavoriteHandler{
		service: service,
	}
}

func (z *zikrFavoriteHandler) UpdateToFavorite(w http.ResponseWriter, r *http.Request) {
	var req model.Patch
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	ok, err := z.service.FavoriteDua(req.UserId, req.ZikrId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if !ok {
		http.Error(w, "could not update", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("updated to favorite"))
	w.WriteHeader(http.StatusOK)
}

func (z *zikrFavoriteHandler) UpdateToUnFavorite(w http.ResponseWriter, r *http.Request) {
	var req model.Patch
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	ok, err := z.service.UnFavoriteDua(req.UserId, req.ZikrId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if !ok {
		http.Error(w, "could not update to Un favorite", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("updated to unfavorite"))
	w.WriteHeader(http.StatusOK)
}

func (z *zikrFavoriteHandler) GetAllFavorites(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBodyForUser
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		return
	}

	favs, err := z.service.GetAllFavorites(requestBody.UserGuid)

	if err != nil {
		http.Error(w, "failed to retrieve favorites : "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(favs)
}

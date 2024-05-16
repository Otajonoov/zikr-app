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

func (z *zikrFavoriteHandler) ToggleFavorite(w http.ResponseWriter, r *http.Request) {
	var req model.Patch
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	ok, err := z.service.FavoriteDua(req.UserGuId, req.Guid)
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

func (z *zikrFavoriteHandler) ToggleUnFavorite(w http.ResponseWriter, r *http.Request) {
	var req model.Patch
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	ok, err := z.service.UnFavoriteDua(req.UserGuId, req.Guid)
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
		http.Error(w, "failed to decode request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	favs, err := z.service.GetAllFavorites(requestBody.UserGuid)
	if err != nil {
		http.Error(w, "failed to retrieve favorites: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var zikr model.Zikrs
	for _, fav := range favs {
		zikr.Zikrs = append(zikr.Zikrs, model.GetZikr{
			Guid:       fav.GetGuid(),
			UserGuid:   fav.GetUserGUID(),
			Arabic:     fav.GetArabic(),
			Uzbek:      fav.GetUzbek(),
			Pronounce:  fav.GetPronounce(),
			Count:      fav.GetCount(),
			IsFavorite: true,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(zikr); err != nil {
		http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (z *zikrFavoriteHandler) GetAllUNFavorites(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBodyForUser
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "failed to decode request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	favs, err := z.service.GetAllFavorites(requestBody.UserGuid)
	if err != nil {
		http.Error(w, "failed to retrieve favorites: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var zikr model.Zikrs
	for _, fav := range favs {
		zikr.Zikrs = append(zikr.Zikrs, model.GetZikr{
			Guid:       fav.GetGuid(),
			UserGuid:   fav.GetUserGUID(),
			Arabic:     fav.GetArabic(),
			Uzbek:      fav.GetUzbek(),
			Pronounce:  fav.GetPronounce(),
			Count:      fav.GetCount(),
			IsFavorite: false,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(zikr); err != nil {
		http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

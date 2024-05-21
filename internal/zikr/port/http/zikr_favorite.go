package http

import (
	"encoding/json"
	"net/http"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/port/model"
)

type zikrFavoriteHandler struct {
	service domain.ZikrFavoritesRepository
	factory domain.Factory
}

func NewZikrFavoriteHandler(service domain.ZikrFavoritesUsecase) *zikrFavoriteHandler {
	return &zikrFavoriteHandler{
		service: service,
	}
}

// @Summary Mark zikr as favorite
// @Description This API marks zikr as favorite
// @Tags zikr-favs
// @Accept json
// @Produce json
// @Param body body model.Patch true "body"
// @Success 200 {string} string "updated to favorite"
// @Failure 400 {string} string "invalid request body"
// @Failure 404 {string} string "could not update"
// @Router /zikr-favs/favorite [patch]
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

// @Summary Mark zikr as unfavorite
// @Description This API marks zikr as unfavorite
// @Tags zikr-favs
// @Accept json
// @Produce json
// @Param body body model.Patch true "body"
// @Success 200 {string} string "updated to unfavorite"
// @Failure 400 {string} string "invalid request body"
// @Failure 404 {string} string "could not update to unfavorite"
// @Router /zikr-favs/unfavorite [patch]
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

// @Summary Get all favorites or unfavorites
// @Description This API retrieves a list of favorite or unfavorite zikrs based on the endpoint
// @Tags zikr-favs
// @Accept json
// @Produce json
// @Param    user_guid query string true "UserGuid of the zikr"
// @Success 200 {object} model.Zikrs "body"
// @Failure 400 {string} string "invalid request body"
// @Failure 404 {string} string "failed to retrieve favorites/unfavorites"
// @Router /zikr-favs/all-favorites [get]
func (z *zikrFavoriteHandler) GetAllFavorites(w http.ResponseWriter, r *http.Request) {
	userGuid := r.URL.Query().Get("user_guid")

	favs, err := z.service.GetAllFavorites(userGuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var zikr model.Zikrs
	for _, fav := range favs {
		zikr.Zikrs = append(zikr.Zikrs, model.GetZikr{
			Guid: fav.Guid,
			//	UserGuid:   fav.GetUserGUID(),

			//Arabic:     fav.GetArabic(),
			//Uzbek:      fav.GetUzbek(),
			//Pronounce:  fav.GetPronounce(),
			//Count:      fav.GetCount(),
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

// @Summary Get all favorites or unfavorites
// @Description This API retrieves a list of favorite or unfavorite zikrs based on the endpoint
// @Tags zikr-favs
// @Accept json
// @Produce json
// @Param  user_guid query string true "GUID of the user"
// @Success 200 {object} model.Zikrs "body"
// @Failure 400 {string} string "invalid request body"
// @Failure 404 {string} string "failed to retrieve favorites/unfavorites"
// @Router /zikr-favs/all-unfavorites [get]
func (z *zikrFavoriteHandler) GetAllUnFavorites(w http.ResponseWriter, r *http.Request) {
	//userGuid := r.URL.Query().Get("user_guid")

	//favs, err := z.service.GetAllUnFavorites(userGuid)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	var zikr model.Zikrs
	//for _, fav := range favs {
	//	zikr.Zikrs = append(zikr.Zikrs, model.GetZikr{
	//		//Guid:       fav.GetGuid(),
	//		//UserGuid:   fav.GetUserGUID(),
	//		//Arabic:     fav.GetArabic(),
	//		//Uzbek:      fav.GetUzbek(),
	//		//Pronounce:  fav.GetPronounce(),
	//		//Count:      fav.GetCount(),
	//		IsFavorite: false,
	//	})
	//}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(zikr); err != nil {
		http.Error(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

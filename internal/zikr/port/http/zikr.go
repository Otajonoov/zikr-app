package http

import (
	"encoding/json"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/port/model"

	"log"
	"net/http"
)

type zikrHandler struct {
	service domain.ZikrUsecase
	factory domain.ZikrFactory
}

func NewZikrHandler(service domain.ZikrUsecase) *zikrHandler {
	return &zikrHandler{
		service: service,
	}
}

// @Summary 	Create zikr
// @Description This api can create new zikr
// @Tags 		Zikr
// @Accept 		json
// @Produce 	json
// @Param body body model.Zikr true "Create"
// @Failure 400 string Error response
// @Router /zikr/create [post]
func (z *zikrHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var zikr model.Zikr
		if err := json.NewDecoder(r.Body).Decode(&zikr); err != nil {
			log.Println("error", err)
			return
		}

		res := z.factory.ParseToControllerForCreate(zikr.Arabic, zikr.Uzbek, zikr.Pronounce)
		err := z.service.Create(res)
		if err != nil {
			log.Println(err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("ok"))
	}
}

// @Summary 	Get by ID zikr
// @Description This api can get by ID zikr
// @Tags 		Zikr
// @Accept 		json
// @Produce 	json
// @Param 		guid query string true "GUID"
// @Success 	200 {object} model.GetZikr
// @Failure 400 string Error response
// @Router /zikr/get [get]
func (z *zikrHandler) Get(w http.ResponseWriter, r *http.Request) {
	guid := r.URL.Query().Get("guid")

	log.Println("Guid: ", guid)
	zikr, err := z.service.Get(guid)
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	res := model.GetZikr{
		Guid:      zikr.GetGuid(),
		Arabic:    zikr.GetArabic(),
		Uzbek:     zikr.GetUzbek(),
		Pronounce: zikr.GetPronounce(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&res)
}

// @Summary 	Get all zikr
// @Description This api can get all zikr
// @Tags 		Zikr
// @Accept 		json
// @Produce 	json
// @Success 200 {object} model.Zikrs "Created successfully"
// @Failure 400 string Error response
// @Router /zikr/get-all [get]
func (z *zikrHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	zikrs, err := z.service.GetAll()
	if err != nil {
		http.Error(w, "failed to retrieve duas : "+err.Error(), http.StatusNotFound)
		return
	}

	var zikr model.Zikrs
	for _, v := range zikrs {
		zikr.Zikrs = append(zikr.Zikrs, model.GetZikr{
			Guid:      v.GetGuid(),
			Arabic:    v.GetArabic(),
			Uzbek:     v.GetUzbek(),
			Pronounce: v.GetPronounce(),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(zikr)
}

// @Summary 	Update zikr
// @Description This api can update zikr
// @Tags 		Zikr
// @Accept 		json
// @Produce 	json
// @Param 		guid query string true "GUID"
// @Param body body model.Zikr true "Update"
// @Failure 400 string Error response
// @Router /zikr/update [put]
func (z *zikrHandler) Update(w http.ResponseWriter, r *http.Request) {
	var zikr model.Zikr

	guid := r.URL.Query().Get("guid")

	if err := json.NewDecoder(r.Body).Decode(&zikr); err != nil {
		log.Println("error", err)
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	res := z.factory.ParseToDomainHandler(guid, zikr.Arabic, zikr.Uzbek, zikr.Pronounce)
	err := z.service.Update(res)
	if err != nil {
		http.Error(w, "failed to update zikr : "+err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("ok"))
}

// @Summary 	Delete zikr
// @Description This api can delete zikr
// @Tags 		Zikr
// @Accept 		json
// @Produce 	json
// @Param 		guid query string true "GUID"
// @Failure 400 string Error response
// @Router /zikr/delete [delete]
func (z *zikrHandler) Delete(w http.ResponseWriter, r *http.Request) {
	guid := r.URL.Query().Get("guid")

	err := z.service.Delete(guid)
	if err != nil {
		http.Error(w, "failed to delete zikr: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("ok")
}

//func (z *zikrHandler) Favorites(w http.ResponseWriter, r *http.Request) {
//	var req model.Favorites
//
//	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
//		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	_, err := z.service.FavoritedDua(req.UserId, req.ZikrId)
//	if err != nil {
//		http.Error(w, "failed to add favorites : "+err.Error(), http.StatusNotFound)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	w.Write([]byte("Updated to favorite"))
//}
//
//func (z *zikrHandler) UnFavorites(w http.ResponseWriter, r *http.Request) {
//	var req model.Favorites
//
//	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
//		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	_, err := z.service.UnFavoritedDua(req.UserId, req.ZikrId)
//	if err != nil {
//		http.Error(w, "failed to add unfavorites : "+err.Error(), http.StatusNotFound)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//}
//
//func (z *zikrHandler) GetAllFavorites(w http.ResponseWriter, r *http.Request) {
//	//UserIdStr := chi.URLParam(r, "user_id")
//	//userId, err := strconv.Atoi(UserIdStr)
//
//	userId := 2
//
//	//if err := nil {
//	//	http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
//	//	return
//	//}
//
//	favs, err := z.service.GetAllFavoriteDuas(userId)
//	if err != nil {
//		http.Error(w, "failed to add favorites : "+err.Error(), http.StatusNotFound)
//		return
//	}
//	var zikrs []domain.Zikr
//	zikrs = make([]domain.Zikr, 0, len(favs))
//	for _, zikr := range favs {
//		zikrs = append(zikrs, zikr)
//	}
//
//	log.Println("favs", zikrs)
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(zikrs)
//}
//

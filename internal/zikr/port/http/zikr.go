package http

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"strconv"
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
// @Success 	200 {object} model.Id
// @Failure 400 string Error response
// @Router /zikr/create [post]
func (z *zikrHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var zikr model.Zikr
		if err := json.NewDecoder(r.Body).Decode(&zikr); err != nil {
			log.Println("error", err)
			return
		}

		res := z.factory.ParseToControllerForCreate(zikr.UserId, zikr.Arabic, zikr.Uzbek, zikr.Pronounce, zikr.IsFavorite)
		err := z.service.Create(res)
		if err != nil {
			log.Println(err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Dua created"))
	}
}

// @Summary 	Get by ID zikr
// @Description This api can get by ID zikr
// @Tags 		Zikr
// @Accept 		json
// @Produce 	json
// @Param 		id query string true "ID"
// @Success 	200 {object} model.Zikr
// @Failure 400 string Error response
// @Router /v1/get [get]
func (z *zikrHandler) Get(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	var newZikr model.Zikr

	id = 6

	zikr, err := z.service.Get(id)
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	newZikr.Id = zikr.GetGUID()
	newZikr.UserId = zikr.GetUserId()
	newZikr.Arabic = zikr.GetArabic()
	newZikr.Uzbek = zikr.GetUzbek()
	newZikr.Pronounce = zikr.GetPronounce()
	newZikr.IsFavorite = zikr.GetIsFavourite()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&newZikr)
}

// @Summary 	Get all zikr
// @Description This api can get all zikr
// @Tags 		Zikr
// @Accept 		json
// @Produce 	json
// @Success 200 {object} model.Zikrs "Created successfully"
// @Failure 400 string Error response
// @Router /v1/get-all [get]
func (z *zikrHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	zikrs, err := z.service.GetAll()
	if err != nil {
		http.Error(w, "failed to retrieve duas : "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(zikrs)
}

func (z *zikrHandler) Favorites(w http.ResponseWriter, r *http.Request) {
	var req model.Favorites

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err := z.service.FavoritedDua(req.UserId, req.ZikrId)
	if err != nil {
		http.Error(w, "failed to add favorites : "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Updated to favorite"))
}

func (z *zikrHandler) UnFavorites(w http.ResponseWriter, r *http.Request) {
	var req model.Favorites

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err := z.service.UnFavoritedDua(req.UserId, req.ZikrId)
	if err != nil {
		http.Error(w, "failed to add unfavorites : "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (z *zikrHandler) GetAllFavorites(w http.ResponseWriter, r *http.Request) {
	//UserIdStr := chi.URLParam(r, "user_id")
	//userId, err := strconv.Atoi(UserIdStr)

	userId := 2

	//if err := nil {
	//	http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
	//	return
	//}

	favs, err := z.service.GetAllFavoriteDuas(userId)
	if err != nil {
		http.Error(w, "failed to add favorites : "+err.Error(), http.StatusNotFound)
		return
	}
	var zikrs []domain.Zikr
	zikrs = make([]domain.Zikr, 0, len(favs))
	for _, zikr := range favs {
		zikrs = append(zikrs, zikr)
	}

	log.Println("favs", zikrs)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(zikrs)
}

// @Summary 	Update zikr
// @Description This api can update zikr
// @Tags 		Zikr
// @Accept 		json
// @Produce 	json
// @Param 		id query string true "ID"
// @Param body body model.Zikr true "Create"
// @Success 	200 {object} model.Id
// @Failure 400 string Error response
// @Router /v1/update [put]
func (z *zikrHandler) Update(w http.ResponseWriter, r *http.Request) {
	var zikr model.Zikr

	res := z.factory.ParseToDomainHandler(zikr.Id, zikr.Arabic, zikr.Uzbek, zikr.Pronounce, zikr.IsFavorite)
	err := z.service.Update(res)
	if err != nil {
		http.Error(w, "failed to update zikr : "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("updated"))
}

func (z *zikrHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var id int

	err := z.service.Delete(id)
	if err != nil {
		http.Error(w, "failed to delete zikr: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("deleted"))
}

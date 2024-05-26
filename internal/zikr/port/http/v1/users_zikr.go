package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/domain/factory"
	"zikr-app/internal/zikr/port/http/model"
)

type usersZikrHandler struct {
	usecase domain.UsersZikrUseCase
	factory factory.Factory
}

func NewCountHandler(u domain.UsersZikrUseCase) *usersZikrHandler {
	return &usersZikrHandler{usecase: u}
}

// @Summary 	Update zikr count
// @Description This API updates zikr count
// @Tags 		users-zikr
// @Accept 		json
// @Produce 	json
// @Param body  body model.Count true "account info"
// @Success   	200 {object} model.Response "Successful response"
// @Failure 	404 string Error response
// @Router /users-zikr/count [patch]
func (c *usersZikrHandler) Count(w http.ResponseWriter, r *http.Request) {
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

// @Summary 	Update zikr favorite
// @Description This API update zikr favorite
// @Tags 		users-zikr
// @Accept 		json
// @Produce 	json
// @Param 		body body model.IsFavorite true "body"
// @Success 	200 {string} string "updated to favorite"
// @Failure 	400 {string} string "invalid request body"
// @Failure 	404 {string} string "could not update"
// @Router /users-zikr/favorite [patch]
func (z *usersZikrHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req model.IsFavorite
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("err favorite: ", err)
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err = z.usecase.Update(req.UserGuId, req.ZikrGuid, req.IsFav); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// @Summary 	Count zikr
// @Description This API count zikr
// @Tags 		users-zikr
// @Accept 		json
// @Produce 	json
// @Param 		filter query model.Reyting false "Filter"
// @Success   	200 {object} model.ReytingResponse "Successful response"
// @Failure 	400 {string} string "invalid request body"
// @Failure 	404 {string} string "could not update"
// @Router /users-zikr/reyting [get]
func (z *usersZikrHandler) Reyting(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	page := r.URL.Query().Get("page")
	zikrGuid := r.URL.Query().Get("zikr_guid")

	l, _ := strconv.Atoi(limit)
	p, _ := strconv.Atoi(page)

	res, err := z.usecase.Reyting(&domain.Reyting{
		Limit:    int64(l),
		Page:     int64(p),
		ZikrGuid: zikrGuid,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

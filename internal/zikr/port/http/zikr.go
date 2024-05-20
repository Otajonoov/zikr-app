package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/port/model"
)

type zikrHandler struct {
	service domain.ZikrUsecase
	factory domain.ZikrFactory
}

type RequestBody struct {
	GUID string
}

func NewZikrHandler(service domain.ZikrUsecase) *zikrHandler {
	return &zikrHandler{
		service: service,
	}
}

// @Summary  Create zikr
// @Description create-zikr
// @Tags      zikr
// @Accept    json
// @Produce   json
// @Param     body body model.ZikrSave true "body"
// @Success   201 {string} string "created"
// @Failure   404 {string} string "Error response"
// @Router    /zikr/create [post]
func (z *zikrHandler) Create(w http.ResponseWriter, r *http.Request) {
	var zikr model.ZikrSave
	if err := json.NewDecoder(r.Body).Decode(&zikr); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	res := z.factory.ParseToDomain(zikr.Guid, zikr.UserGuid, zikr.Arabic, zikr.Uzbek, zikr.Pronounce, zikr.Count, zikr.IsFavorite, zikr.CreatedAt, zikr.UpdatedAt)
	err := z.service.Create(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("created"))
}

// @Summary  Get zikr
// @Description This API gets a zikr
// @Tags      zikr
// @Accept    json
// @Produce   json
// @Param     guid  query string true "GUID of the zikr"
// @Success   200   {object} model.GetZikr "Successful response"
// @Failure   404   {string} Error "Error response"
// @Router    /zikr/get [get]
func (z *zikrHandler) Get(w http.ResponseWriter, r *http.Request) {
	guid := r.URL.Query().Get("guid")
	if guid == "" {
		http.Error(w, "missing GUID parameter", http.StatusBadRequest)
		return
	}

	zikr, err := z.service.Get(guid)
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	res := model.GetZikr{
		Guid:       zikr.GetGuid(),
		UserGuid:   zikr.GetUserGUID(),
		Arabic:     zikr.GetArabic(),
		Uzbek:      zikr.GetUzbek(),
		Pronounce:  zikr.GetPronounce(),
		Count:      zikr.GetCount(),
		IsFavorite: zikr.GetIsFavorite(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&res)
}

// @Summary  Get zikr list
// @Description This API gets a list of zikr
// @Tags      zikr
// @Accept    json
// @Produce   json
// @Success   200 {object} model.Zikrs "Successful response"
// @Failure   404 {string} string "Error response"
// @Router    /zikr/list [get]
func (z *zikrHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	zikrs, err := z.service.GetAll()
	if err != nil {
		http.Error(w, "failed to retrieve duas : "+err.Error(), http.StatusNotFound)
		return
	}

	var zikr model.Zikrs
	for _, v := range zikrs {
		zikr.Zikrs = append(zikr.Zikrs, model.GetZikr{
			Guid:       v.GetGuid(),
			UserGuid:   v.GetUserGUID(),
			Arabic:     v.GetArabic(),
			Uzbek:      v.GetUzbek(),
			Pronounce:  v.GetPronounce(),
			Count:      v.GetCount(),
			IsFavorite: v.GetIsFavorite(),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(zikr)
}

// @Summary  Update zikr
// @Description This API updates a zikr
// @Tags      zikr
// @Accept    json
// @Produce   json
// @Param     guid  query string true "GUID of the zikr"
// @Param     body body model.Zikr true "body"
// @Success   200 {string} string "updated"
// @Failure   404 {string} string "Error response"
// @Router    /zikr/update [put]
func (z *zikrHandler) Update(w http.ResponseWriter, r *http.Request) {
	guid := r.URL.Query().Get("guid")

	if guid == "" {
		http.Error(w, "missing GUID parameter", http.StatusBadRequest)
		return
	}
	var zikr model.Zikr

	if err := json.NewDecoder(r.Body).Decode(&zikr); err != nil {
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
	w.Write([]byte("updated"))
}

// @Summary  Patch zikr count
// @Description This API patches the zikr count
// @Tags      zikr
// @Accept    json
// @Produce   json
// @Param     guid      query string true "GUID of the zikr"
// @Param     user_guid query string true "GUID of the user"
// @Param     count  query string true "count of zikr"
// @Success   200 {string} string "count updated"
// @Failure   400 {string} string "Invalid request"
// @Failure   404 {string} string "Zikr not found"
// @Router    /zikr/count [patch]
func (z *zikrHandler) PatchCount(w http.ResponseWriter, r *http.Request) {
	guid := r.URL.Query().Get("guid")
	if guid == "" {
		http.Error(w, "missing GUID parameter", http.StatusBadRequest)
		return
	}

	userGuid := r.URL.Query().Get("user_guid")
	if userGuid == "" {
		http.Error(w, "missing user_guid parameter", http.StatusBadRequest)
		return
	}

	countStr := r.URL.Query().Get("count")
	count, err := strconv.Atoi(countStr)
	if count == 0 || count < 0 || err != nil {
		http.Error(w, "invalid count parameter: "+err.Error(), http.StatusBadRequest)
		return
	}

	res := z.factory.ParseToDomainToPatch(guid, userGuid, count)

	err = z.service.UpdateZikrCount(res)
	if err != nil {
		http.Error(w, "failed to update zikr count: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`"count updated"`))
}

// @Summary Delete zikr
// @Description This api Delete zikr
// @Tags zikr
// @Accept json
// @Produce json
// @Param body body RequestBody true "body"
// @Failure 404 string Error response
// @Router /zikr/delete [delete]
func (z *zikrHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		return
	}

	err := z.service.Delete(requestBody.GUID)
	if err != nil {
		http.Error(w, "failed to delete zikr: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("deleted")
}

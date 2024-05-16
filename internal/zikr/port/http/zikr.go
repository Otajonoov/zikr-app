package http

import (
	"encoding/json"
	"net/http"
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

type RequestBodyForUser struct {
	UserGuid string
}

func NewZikrHandler(service domain.ZikrUsecase) *zikrHandler {
	return &zikrHandler{
		service: service,
	}
}

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

func (z *zikrHandler) Get(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		return
	}

	zikr, err := z.service.Get(requestBody.GUID)
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

func (z *zikrHandler) Update(w http.ResponseWriter, r *http.Request) {
	var zikr model.Zikr

	guid := r.URL.Query().Get("guid")

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
	w.Write([]byte("ok"))
}

func (z *zikrHandler) PatchCount(w http.ResponseWriter, r *http.Request) {
	var patch model.PatchCount

	if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	res := z.factory.ParseToDomainToPatch(patch.Guid, patch.UserGuid, patch.Count)
	err := z.service.UpdateZikrCount(res)
	if err != nil {
		http.Error(w, "failed to update zikr count : "+err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("count updated"))
}

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
	json.NewEncoder(w).Encode("ok")
}

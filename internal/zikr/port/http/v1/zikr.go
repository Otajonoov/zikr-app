package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/domain/factory"
	"zikr-app/internal/zikr/port/http/model"
)

type zikrHandler struct {
	service domain.ZikrUsecase
	factory factory.Factory
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
// @Param     body body model.Zikr true "body"
// @Success   200   {object} model.Response "success"
// @Failure   400   {string} string "Invalid request body"
// @Failure   500   {string} string "Internal server error"
// @Router    /zikr [post]
func (z *zikrHandler) Create(w http.ResponseWriter, r *http.Request) {
	var zikr model.Zikr
	if err := json.NewDecoder(r.Body).Decode(&zikr); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	res := z.factory.ParseToControllerForCreate(zikr.Arabic, zikr.Uzbek, zikr.Pronounce)
	err := z.service.Create(res)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.Response{Result: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.Response{Result: "ok"})
}

// @Summary  Get zikr list
// @Description This API gets a list of zikr
// @Tags      zikr
// @Accept    json
// @Produce   json
// @Param  guid query string true "GUID of the user"
// @Success   200 {object} model.Zikrs "Successful response"
// @Failure   404 {string} string "Error response"
// @Router    /zikr [get]
func (z *zikrHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	guid := r.URL.Query().Get("guid")
	zikrs, err := z.service.GetAll(guid)
	if err != nil {
		log.Println("failed to retrieve duas : ", err)
		http.Error(w, "failed to retrieve duas : "+err.Error(), http.StatusNotFound)
		return
	}

	var zikr model.Zikrs
	for _, z := range zikrs {
		zikr.Zikrs = append(zikr.Zikrs, model.GetZikr{
			Guid:       z.Guid,
			Arabic:     z.Arabic,
			Uzbek:      z.Uzbek,
			Pronounce:  z.Pronounce,
			Count:      z.Count,
			IsFavorite: z.IsFavorite,
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
// @Param     zikrId  query string true "GUID of the zikr"
// @Param     body body model.Zikr true "body"
// @Success   200 {string} string "updated"
// @Failure   404 {string} string "Error response"
// @Router    /zikr [put]
func (z *zikrHandler) Update(w http.ResponseWriter, r *http.Request) {
	zikrId := r.URL.Query().Get("zikrId")
	var zikr model.Zikr

	if err := json.NewDecoder(r.Body).Decode(&zikr); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	res := z.factory.ParseToDomainToUpdate(zikrId, zikr.Arabic, zikr.Uzbek, zikr.Pronounce)
	err := z.service.Update(res)
	if err != nil {
		http.Error(w, "failed to update zikr : "+err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("updated"))
}

// @Summary Delete zikr
// @Description This api Delete zikr
// @Tags 	zikr
// @Accept 	json
// @Produce json
// @Param   zikrId  query string true "GUID of the zikr"
// @Failure 404 string Error response
// @Router /zikr [delete]
func (z *zikrHandler) Delete(w http.ResponseWriter, r *http.Request) {
	zikrId := r.URL.Query().Get("zikrId")

	err := z.service.Delete(zikrId)
	if err != nil {
		http.Error(w, "failed to delete zikr: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("deleted")
}

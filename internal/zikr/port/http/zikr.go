package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/domain/factory"
	"zikr-app/internal/zikr/port/model"
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
// @Router    /zikr/create [post]
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

//func (z *zikrHandler) Get(w http.ResponseWriter, r *http.Request) {
//	guid := r.URL.Query().Get("guid")
//	if guid == "" {
//		http.Error(w, "missing GUID parameter", http.StatusBadRequest)
//		return
//	}
//
//	zikr, err := z.service.Get(guid)
//	if err != nil {
//		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	res := model.GetZikr{
//		Guid:       zikr.Guid,
//		UserEmail:  zikr.UserEmail,
//		Arabic:     zikr.Arabic,
//		Uzbek:      zikr.Uzbek,
//		Pronounce:  zikr.Pronounce,
//		Count:      zikr.Count,
//		IsFavorite: zikr.IsFavorite,
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(&res)
//}

// @Summary  Get zikr list
// @Description This API gets a list of zikr
// @Tags      zikr
// @Accept    json
// @Produce   json
// @Param  guid query string true "GUID of the user"
// @Success   200 {object} model.Zikrs "Successful response"
// @Failure   404 {string} string "Error response"
// @Router    /zikr/get-all [get]
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

// sdffd
// // @Summary  Update zikr
// // @Description This API updates a zikr
// // @Tags      zikr
// // @Accept    json
// // @Produce   json
// // @Param     userGuid  query string true "GUID of the user"
// // @Param     zikrGuid  query string true "GUID of the zikr"
// // @Param     body body model.Zikr true "body"
// // @Success   200 {string} string "updated"
// // @Failure   404 {string} string "Error response"
// // @Router    /zikr/update [put]
func (z *zikrHandler) Update(w http.ResponseWriter, r *http.Request) {
	userGuid := r.URL.Query().Get("userGuid")

	var zikr model.Zikr

	if err := json.NewDecoder(r.Body).Decode(&zikr); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	res := z.factory.ParseToDomainToUpdate(userGuid, zikr.Arabic, zikr.Uzbek, zikr.Pronounce)
	err := z.service.Update(res)
	if err != nil {
		http.Error(w, "failed to update zikr : "+err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("updated"))
}

// zxcds
// // @Summary  Patch zikr count
// // @Description This API patches the zikr count
// // @Tags      zikr
// // @Accept    json
// // @Produce   json
// // @Param     zikr_guid      query string true "GUID of the zikr"
// // @Param     user_guid query string true "GUID of the user"
// // @Param     count  query string true "count of zikr"
// // @Success   200 {string} string "count updated"
// // @Failure   400 {string} string "Invalid request"
// // @Failure   404 {string} string "Zikr not found"
// // @Router    /zikr/count [patch]
func (z *zikrHandler) PatchCount(w http.ResponseWriter, r *http.Request) {
	//zikrGuid := r.URL.Query().Get("zikr_guid")
	//userGuid := r.URL.Query().Get("user_guid")

	countStr := r.URL.Query().Get("count")
	count, err := strconv.Atoi(countStr)
	if count == 0 || count < 0 || err != nil {
		log.Println("invalid count parameter: ", err)
		http.Error(w, "invalid count parameter: "+err.Error(), http.StatusBadRequest)
		return
	}

	//res := z.factory.

	//err = z.service.UpdateZikrCount(res)
	//if err != nil {
	//	http.Error(w, "failed to update zikr count: "+err.Error(), http.StatusNotFound)
	//	return
	//}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`"count updated"`))
}

// zcz
// // @Summary Delete zikr
// // @Description This api Delete zikr
// // @Tags zikr
// // @Accept json
// // @Produce json
// // @Param body body RequestBody true "body"
// // @Failure 404 string Error response
// // @Router /zikr/delete [delete]
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

package http

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/port/model"
)

type zikrCountHandler struct {
	service domain.ZikrCountUsecase
}

func NewZikrCountHandler(service domain.ZikrCountUsecase) zikrCountHandler {
	return zikrCountHandler{service: service}
}

func (z zikrCountHandler) CreateCount(w http.ResponseWriter, r *http.Request) {
	var req model.ZikrCount

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := z.service.CreateCount(context.Background(), &domain.ZikrCount{
		UserId: req.UserId,
		ZikrId: req.ZikrId,
		Count:  req.Count,
	}); err != nil {
		http.Error(w, "failed to create zikr count: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Signed up"))
	w.WriteHeader(http.StatusCreated)
}

func (z zikrCountHandler) ListCount(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "user_id")

	userId, err := strconv.Atoi(userIdStr)

	res, err := z.service.GetUserCounts(context.Background(), userId)
	if err != nil {
		http.Error(w, "failed to list zikr count: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (z zikrCountHandler) PatchUserCount(w http.ResponseWriter, r *http.Request) {
	var req model.Patch

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = z.service.PatchCount(context.Background(), req.UserId, req.Count)
	if err != nil {
		http.Error(w, "failed to update zikr count: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("updated"))
}

func (z zikrCountHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "user_id")
	userId, err := strconv.Atoi(userIdStr)

	if err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = z.service.ResetCount(context.Background(), userId)
	if err != nil {
		http.Error(w, "failed to reset zikr count: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("deleted"))
}

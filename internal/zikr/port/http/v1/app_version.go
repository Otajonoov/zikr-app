package v1

import (
	"encoding/json"
	"net/http"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/domain/factory"
	"zikr-app/internal/zikr/port/http/model"
)

type AppVersionHandler struct {
	usecase domain.AppVersionUsecase
	factory factory.Factory
}

func NewAppVersionHandler(u domain.AppVersionUsecase) *AppVersionHandler {
	return &AppVersionHandler{usecase: u}
}

// @Summary 	Get app version
// @Description This API gets app version
// @Tags 		app-version
// @Accept 		json
// @Produce 	json
// @Success   200 {object} model.AppVersion "Successful response"
// @Failure 404 string Error response
// @Router /app-version [get]
func (a *AppVersionHandler) GetAppVersion(w http.ResponseWriter, r *http.Request) {
	appVersion, err := a.usecase.GetAppVersion()
	if err != nil {
		http.Error(w, "error getting app version", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(appVersion)
}

// @Summary 	Update app version
// @Description This API updates app version
// @Tags 		app-version
// @Accept 		json
// @Produce 	json
// @Param body  body model.AppVersion true "app version"
// @Success   200 {object} model.AppVersion "Successful response"
// @Failure 404 string Error response
// @Router /app-version [put]
func (a *AppVersionHandler) Update(w http.ResponseWriter, r *http.Request) {
	var appVersion model.AppVersion
	err := json.NewDecoder(r.Body).Decode(&appVersion)
	if err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	version := a.factory.ParseToDomainForAppVersion(appVersion.Anndroid, appVersion.Ios, appVersion.ForceUpdate)
	err = a.usecase.Update(version)
	if err != nil {
		http.Error(w, "error updating app version: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

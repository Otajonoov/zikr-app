package http

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
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
			log.Println(err)
			return
		}

		res := z.factory.ParseToController(zikr.Arabic, zikr.Uzbek, zikr.Pronounce)
		err := z.service.Create(res)
		if err != nil {
			log.Println(err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
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
func (z *zikrHandler) Get(ctx *gin.Context) {
	id := ctx.Query("id")

	zikr, err := z.service.Get(id)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error while getting value",
		})
		return
	}
	ctx.JSON(http.StatusOK, model.Zikr{
		Arabic:    zikr.GetArabic(),
		Uzbek:     zikr.GetUzbek(),
		Pronounce: zikr.GetPronounce(),
	})
}

// @Summary 	Get all zikr
// @Description This api can get all zikr
// @Tags 		Zikr
// @Accept 		json
// @Produce 	json
// @Success 200 {object} model.Zikrs "Created successfully"
// @Failure 400 string Error response
// @Router /v1/get-all [get]
func (z *zikrHandler) GetAll(ctx *gin.Context) {

	zikrs, err := z.service.GetAll()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error while getting value",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"zikrs": zikrs,
	})
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
func (z *zikrHandler) Update(ctx *gin.Context) {
	id := ctx.Query("id")

	var zikr model.Zikr
	if err := ctx.ShouldBindJSON(&zikr); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error while getting value",
		})
		return
	}

	res := z.factory.ParseToDomain(id, zikr.Arabic, zikr.Uzbek, zikr.Pronounce)
	err := z.service.Update(res)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error while update value",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"ok": "successfully update",
	})
}

// @Summary 	Delete zikr
// @Description This api can delete zikr
// @Tags 		Zikr
// @Accept 		json
// @Produce 	json
// @Param 		id query string true "ID"
// @Failure 400 string Error response
// @Router /v1/delete [delete]
func (z *zikrHandler) Delete(ctx *gin.Context) {
	id := ctx.Query("id")

	err := z.service.Delete(id)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error while delete",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id": "successfully delete",
	})
}

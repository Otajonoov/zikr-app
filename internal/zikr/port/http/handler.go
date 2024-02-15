package http

import (
	"github.com/gin-gonic/gin"
	"zikr-app/internal/zikr/domain"
)

type ZikrController struct {
	Service domain.ZikrUsecase
	Factory domain.ZikrFactory
}

func NewZikrController(service domain.ZikrUsecase) *ZikrController {
	return &ZikrController{
		Service: service,
	}
}

type ZikrHandler interface {
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	//Delete(ctx *gin.Context)
}

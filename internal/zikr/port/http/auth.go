package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/port/model"
)

type AuthHandler struct {
	usecase domain.AuthUsecase
}

func NewAuthHandler(u domain.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: u}
}

// @Summary 	Sign-Up user
// @Description This api can Sign-Up new user
// @Tags 		User
// @Accept 		json
// @Produce 	json
// @Param body  body model.User true "body"
// @Failure 404 string Error response
// @Router /v1/sign-up [post]
func (u *AuthHandler) SignUp(ctx *gin.Context) {
	var req model.User
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	if err := u.usecase.CreateUser(ctx, &domain.User{
		FIO:           req.FIO,
		UniqeUsername: req.UniqeUsername,
		PhoneNumber:   req.PhoneNumber,
		Password:      req.Password,
	}); err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, "ok")
}

// @Summary 	Sign-In user
// @Description This api can Sign-In user
// @Tags 		User
// @Accept 		json
// @Produce 	json
// @Param body  body model.SignIn true "body"
// @Failure 404 string Error response
// @Router /v1/sign-in [post]
func (u *AuthHandler) SignIn(ctx *gin.Context) {
	var req model.SignIn
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	if ok := u.usecase.CheckUser(ctx, &domain.User{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	}); !ok {
		ctx.JSON(http.StatusNotFound, "Bunday foydalanuvchi mavjud emas")
		return
	}

	ctx.JSON(http.StatusOK, "ok")
}

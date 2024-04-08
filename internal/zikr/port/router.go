package port

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/port/http"
	_ "zikr-app/internal/zikr/port/http/docs"
)

type RouterOption struct {
	ZikrUsecase domain.ZikrUsecase
	AuthUsecase domain.AuthUsecase

	Factory domain.ZikrFactory
}

// @Description Created by Otajonov Quvonchbek
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(option RouterOption) *gin.Engine {
	router := gin.New()

	zikrController := http.NewZikrController(option.ZikrUsecase)
	authController := http.NewAuthHandler(option.AuthUsecase)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	corConfig := cors.DefaultConfig()
	corConfig.AllowAllOrigins = true
	corConfig.AllowCredentials = true
	corConfig.AllowBrowserExtensions = true
	corConfig.AllowHeaders = append(corConfig.AllowHeaders, "*")
	router.Use(cors.New(corConfig))

	api := router.Group("/v1")

	// Zikr
	api.POST("/create", zikrController.Create)
	api.GET("/get", zikrController.Get)
	api.GET("/get-all", zikrController.GetAll)
	api.PUT("/update", zikrController.Update)
	api.DELETE("/delete", zikrController.Delete)

	// User
	api.POST("sign-up", authController.SignUp)
	api.POST("sign-in", authController.SignIn)

	url := ginSwagger.URL("/v1/swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

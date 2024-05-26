package port

import (
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	httpSwagger "github.com/swaggo/http-swagger"
	"zikr-app/internal/zikr/adapter"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/domain/factory"
	_ "zikr-app/internal/zikr/port/http/docs"
	"zikr-app/internal/zikr/port/http/v1"
	"zikr-app/internal/zikr/usecase"
)

type RouterOption struct {
	ZikrUsecase domain.ZikrUsecase
	AuthUsecase domain.AuthUsecase
	DB          *pgxpool.Pool

	Factory factory.Factory
}

// @Description Created by Otajonov Quvonchbek and Usmonov Azizbek
// securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(option RouterOption) *chi.Mux {

	router := chi.NewRouter()

	router.Use(chimiddleware.RequestID)
	router.Use(chimiddleware.Logger)
	router.Use(chimiddleware.Recoverer)
	router.Use(chimiddleware.URLFormat)

	// Repos
	authRepo := adapter.NewAuthRepo(option.DB)
	zikrRepo := adapter.NewZikrRepo(option.DB)
	usersZikrRepo := adapter.NewCountRepo(option.DB)
	appVersionRepo := adapter.NewAppVersionRepo(option.DB)

	// Usecase
	authUsecase := usecase.NewAuthUsecase(authRepo, zikrRepo)
	zikrUsecase := usecase.NewZikrUsecase(zikrRepo)
	usersZikrtUseCase := usecase.NewCountUsecase(usersZikrRepo)
	appVersionUseCase := usecase.NewAppVersionUsecase(appVersionRepo)

	// Handlers
	authHandler := v1.NewAuthHandler(authUsecase)
	zikrHandler := v1.NewZikrHandler(zikrUsecase)
	usersZikrHandler := v1.NewCountHandler(usersZikrtUseCase)
	appVersionHandler := v1.NewAppVersionHandler(appVersionUseCase)

	// User registration
	router.Route("/auth", func(r chi.Router) {
		r.Post("/", authHandler.CheckUserRegister)
	})

	// App version
	router.Route("/app-version", func(r chi.Router) {
		r.Get("/", appVersionHandler.GetAppVersion)
		r.Put("/", appVersionHandler.Update)
	})

	// Zikr
	router.Route("/zikr", func(r chi.Router) {
		r.Post("/", zikrHandler.Create)
		r.Get("/", zikrHandler.GetAll)
	})

	// Users Zikr
	router.Route("/users-zikr", func(r chi.Router) {
		r.Patch("/count", usersZikrHandler.Count)
		r.Patch("/favorite", usersZikrHandler.Update)
		r.Get("/reyting", usersZikrHandler.Reyting)

	})

	router.Get("/swagger/*", httpSwagger.Handler())
	return router
}

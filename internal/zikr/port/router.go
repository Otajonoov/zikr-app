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
	countRpo := adapter.NewCountRepo(option.DB)
	zikrFavoriteRepo := adapter.NewZikrFavoritesRepo(option.DB)
	appVersionRepo := adapter.NewAppVersionRepo(option.DB)

	// Usecase
	authUsecase := usecase.NewAuthUsecase(authRepo, zikrRepo)
	zikrUsecase := usecase.NewZikrUsecase(zikrRepo)
	countUseCase := usecase.NewCountUsecase(countRpo)
	zikrFavoriteUseCase := usecase.NewZikrFavoritesUsecase(zikrFavoriteRepo)
	appVersionUseCase := usecase.NewAppVersionUsecase(appVersionRepo)

	// Handlers
	authHandler := v1.NewAuthHandler(authUsecase)
	zikrHandler := v1.NewZikrHandler(zikrUsecase)
	countHandler := v1.NewCountHandler(countUseCase)
	zikrFavoriteHandler := v1.NewZikrFavoriteHandler(zikrFavoriteUseCase)
	appVersionHandler := v1.NewAppVersionHandler(appVersionUseCase)

	// User registration
	router.Route("/user", func(r chi.Router) {
		r.Post("/", authHandler.CheckUserRegister)
	})

	// Zikr
	router.Route("/zikr", func(r chi.Router) {
		r.Post("/", zikrHandler.Create)
		r.Get("/", zikrHandler.GetAll)

		//r.Put("/update", zikrHandler.Update)
		//r.Patch("/count", zikrHandler.PatchCount)
		//r.Delete("/delete", zikrHandler.Delete)
	})

	// Zikr count
	router.Route("/count", func(r chi.Router) {
		r.Patch("/", countHandler.Count)
	})

	// Zikr favorites
	router.Route("/favorite", func(r chi.Router) {
		r.Patch("/", zikrFavoriteHandler.Update)
	})

	// App version
	router.Route("/app-version", func(r chi.Router) {
		r.Get("/", appVersionHandler.GetAppVersion)
		r.Put("/", appVersionHandler.Update)
	})

	router.Get("/swagger/*", httpSwagger.Handler())
	return router
}

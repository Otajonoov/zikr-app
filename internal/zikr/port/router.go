package port

import (
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	httpSwagger "github.com/swaggo/http-swagger"
	"zikr-app/internal/zikr/adapter"
	"zikr-app/internal/zikr/domain"
	handler "zikr-app/internal/zikr/port/http"
	_ "zikr-app/internal/zikr/port/http/docs"
	"zikr-app/internal/zikr/usecase"
)

type RouterOption struct {
	ZikrUsecase domain.ZikrUsecase
	AuthUsecase domain.AuthUsecase
	DB          *pgxpool.Pool

	Factory domain.Factory
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
	zikrFavoriteRepo := adapter.NewZikrFavoritesRepo(option.DB)

	// Usecase
	authUsecase := usecase.NewAuthUsecase(authRepo, zikrRepo)
	zikrUsecase := usecase.NewZikrUsecase(zikrRepo)
	zikrFavoriteUseCase := usecase.NewZikrFavoritesUsecase(zikrFavoriteRepo)

	// Handlers
	authHandler := handler.NewAuthHandler(authUsecase)
	zikrHandler := handler.NewZikrHandler(zikrUsecase)
	zikrFavoriteHandler := handler.NewZikrFavoriteHandler(zikrFavoriteUseCase)

	// User registration
	router.Route("/user", func(r chi.Router) {
		r.Post("/create-user", authHandler.CheckUserRegister)
	})

	// Zikr
	router.Route("/zikr", func(r chi.Router) {
		r.Post("/create", zikrHandler.Create)
		r.Get("/get", zikrHandler.Get)
		r.Get("/list", zikrHandler.GetAll)
		r.Put("/update", zikrHandler.Update)
		r.Patch("/count", zikrHandler.PatchCount)
		r.Delete("/delete", zikrHandler.Delete)
	})

	// Zikr favorites
	router.Route("/zikr-favs", func(r chi.Router) {
		r.Patch("/favorite", zikrFavoriteHandler.ToggleFavorite)
		r.Patch("/unfavorite", zikrFavoriteHandler.ToggleUnFavorite)
		r.Get("/all-favorites", zikrFavoriteHandler.GetAllFavorites)
		r.Get("/all-unfavorites", zikrFavoriteHandler.GetAllUnFavorites)
	})

	router.Get("/swagger/*", httpSwagger.Handler())
	return router
}

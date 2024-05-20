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

	Factory domain.ZikrFactory
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

	factory := domain.NewZikrFactory()

	authRepo := adapter.NewAuthRepo(option.DB)
	zikrRepo := adapter.NewZikrRepo(option.DB, factory)
	authUsecase := usecase.NewAuthUsecase(authRepo, zikrRepo)
	authHandler := handler.NewAuthHandler(authUsecase)

	// Routers
	router.Route("/user", func(r chi.Router) {
		r.Post("/check-or-register", authHandler.CheckUserRegister)
	})

	// Zikr
	zikrUsecase := usecase.NewZikrUsecase(zikrRepo, factory)
	zikrHandler := handler.NewZikrHandler(zikrUsecase)

	// Routers
	router.Route("/zikr", func(r chi.Router) {
		r.Post("/create", zikrHandler.Create)
		r.Get("/get", zikrHandler.Get)
		r.Get("/list", zikrHandler.GetAll)
		r.Put("/update", zikrHandler.Update)
		r.Patch("/count", zikrHandler.PatchCount)
		r.Delete("/delete", zikrHandler.Delete)
	})

	// Zikr Favorites
	zikrFavoriteRepo := adapter.NewZikrFavoritesRepo(option.DB, factory)
	zikrFavoriteUseCase := usecase.NewZikrFavoritesUsecase(zikrFavoriteRepo, factory)
	zikrFavoriteHandler := handler.NewZikrFavoriteHandler(zikrFavoriteUseCase)
	// Routers
	router.Route("/zikr-favs", func(r chi.Router) {
		r.Patch("/favorite", zikrFavoriteHandler.ToggleFavorite)
		r.Patch("/unfavorite", zikrFavoriteHandler.ToggleUnFavorite)
		r.Get("/all-favorites", zikrFavoriteHandler.GetAllFavorites)
		r.Get("/all-unfavorites", zikrFavoriteHandler.GetAllUnFavorites)
	})

	router.Get("/swagger/*", httpSwagger.Handler())
	return router
}

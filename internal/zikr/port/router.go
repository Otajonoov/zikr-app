package port

import (
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"zikr-app/internal/zikr/adapter"
	"zikr-app/internal/zikr/domain"
	handler "zikr-app/internal/zikr/port/http"
	"zikr-app/internal/zikr/usecase"
)

type RouterOption struct {
	ZikrUsecase domain.ZikrUsecase
	AuthUsecase domain.AuthUsecase
	DB          *pgxpool.Pool

	Factory domain.ZikrFactory
}

func New(option RouterOption) *chi.Mux {

	router := chi.NewRouter()

	router.Use(chimiddleware.RequestID)
	router.Use(chimiddleware.Logger)
	router.Use(chimiddleware.Recoverer)
	router.Use(chimiddleware.URLFormat)

	factory := domain.NewZikrFactory()

	authRepo := adapter.NewAuthRepo(option.DB)
	authUsecase := usecase.NewAuthUsecase(authRepo)
	authHandler := handler.NewAuthHandler(authUsecase)

	// Routers
	router.Route("/user", func(r chi.Router) {
		r.Post("/register", authHandler.UserRegister)
		r.Post("/check-or-register", authHandler.CheckUserRegister)
	})

	// Zikr
	zikrRepo := adapter.NewZikrRepo(option.DB, factory)
	zikrUsecase := usecase.NewZikrUsecase(zikrRepo, factory)
	zikrHandler := handler.NewZikrHandler(zikrUsecase)

	// Routers
	router.Route("/zikr", func(r chi.Router) {
		r.Post("/create", zikrHandler.Create())
		r.Get("/get", zikrHandler.Get)
		r.Get("/get-all", zikrHandler.GetAll)
		r.Put("/update", zikrHandler.Update)
		r.Delete("/delete", zikrHandler.Delete)
	})

	// Zikr Favorites
	zikrFavoriteRepo := adapter.NewZikrFavoritesRepo(option.DB, factory)
	zikrFavoriteUseCase := usecase.NewZikrFavoritesUsecase(zikrFavoriteRepo, factory)
	zikrFavoriteHandler := handler.NewZikrFavoriteHandler(zikrFavoriteUseCase)
	// Routers
	router.Route("/favorite", func(r chi.Router) {
		r.Patch("/favor", zikrFavoriteHandler.UpdateToFavorite)
		r.Patch("/unfavor", zikrFavoriteHandler.UpdateToUnFavorite)
		r.Get("/get-favorites", zikrFavoriteHandler.GetAllFavorites)
	})

	return router
}

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

// @Description Created by Otajonov Quvonchbek
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(option RouterOption) *chi.Mux {

	router := chi.NewRouter()

	router.Use(chimiddleware.RequestID)
	router.Use(chimiddleware.Logger)
	//router.Use(mwLogger.New(log))
	router.Use(chimiddleware.Recoverer)
	router.Use(chimiddleware.URLFormat)

	factory := domain.NewZikrFactory()

	// Zikr
	zikrRepo := adapter.NewZikrRepo(option.DB, factory)
	zikrUsecase := usecase.NewZikrUsecase(zikrRepo, factory)
	zikrHandler := handler.NewZikrHandler(zikrUsecase)

	// Routers
	router.Route("/zikr", func(r chi.Router) {
		r.Post("/create", zikrHandler.Create())
		r.Get("/get", zikrHandler.Create())
		r.Get("/get-all", zikrHandler.Create())
		r.Put("/update", zikrHandler.Create())
	})

	// Swagger integration
	router.Get("/swagger/*", httpSwagger.Handler())
	return router
}

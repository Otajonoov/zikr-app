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
	"zikr-app/pkg/jwt"
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

	// Auth
	authRepo := adapter.NewAuthRepo(option.DB)
	authUsecase := usecase.NewAuthUsecase(authRepo)
	authHandler := handler.NewAuthHandler(authUsecase)

	// Routers
	router.Route("/user", func(r chi.Router) {
		r.Post("/sign-up-user", authHandler.SignUp)
		r.Post("/sign-in-user", authHandler.SignIn)
		r.With(jwt.AuthMiddleWare).Get("/get-user/{username}", authHandler.GetUserByUserName)
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
		//r.Get("/get-favorites", zikrHandler.GetAllFavorites)
		//r.Patch("/favorite", zikrHandler.Favorites)
		//r.Patch("/unfavorite", zikrHandler.UnFavorites)
		r.Put("/update", zikrHandler.Update)
		r.Delete("/delete", zikrHandler.Delete)
	})

	// ZikrCount
	zikrCountRepo := adapter.NewZikrCountRepo(option.DB)
	zikrCountUseCase := usecase.NewZikrCountUsecase(zikrCountRepo)
	zikrCountHandler := handler.NewZikrCountHandler(zikrCountUseCase)

	// Routers
	router.Route("/zikr-count", func(r chi.Router) {
		r.Post("/add-count", zikrCountHandler.CreateCount)
		r.Get("/list-user-counts", zikrCountHandler.ListCount)
		r.Patch("/increment", zikrCountHandler.PatchUserCount)
		r.Patch("/reset", zikrCountHandler.Delete)
	})

	// Swagger integration
	router.Get("/swagger/*", httpSwagger.Handler())
	return router
}

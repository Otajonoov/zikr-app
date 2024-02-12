package main

import (
	"log"
	"net/http"
	"zikr-app/internal/pkg/config"
	db2 "zikr-app/internal/pkg/db"
	"zikr-app/internal/zikr/adapter"
	"zikr-app/internal/zikr/domain"
	"zikr-app/internal/zikr/port"
	"zikr-app/internal/zikr/usecase"
)

func main() {
	cfg := config.Load()
	pgxConn, err := db2.ConnDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pgxConn.Close()

	factory := domain.NewZikrFactory()
	zikrRepo := adapter.NewZikrRepo(pgxConn, factory)
	zikrUsecase := usecase.NewZikrUsecase(zikrRepo, factory)

	apiServer := port.New(port.RouterOption{UseCase: zikrUsecase, Factory: factory})

	server := &http.Server{
		Addr:    ":" + cfg.HttpPort, // 5005
		Handler: apiServer,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	select {}
}

package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	abrands "dresser/internal/application/brands"
	dbrands "dresser/internal/domain/brands"
	"dresser/internal/infrastructure/config"
	"dresser/internal/infrastructure/db"
	pbrands "dresser/internal/infrastructure/persistence/brands"
	httpserver "dresser/internal/interface/http"

	"net/http"
)

func main() {
	// Load configuration
	cfg, err := config.LoadDefault()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database
	database, err := db.New(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	// Initialize repositories
	repo := pbrands.NewBrandRepository(database)
	fact := dbrands.NewFactory(repo)

	// Initialize application services
	brandRegisterService := abrands.NewBrandRegisterService(repo, &fact)
	brandQueryService := abrands.NewBrandQueryService(repo)
	brandDeleteService := abrands.NewBrandDeleteService(repo)
	// Initialize HTTP server
	server := httpserver.NewHTTPServer(httpserver.ServerConfig{
		Port:                 "8080",
		BrandRegisterService: brandRegisterService,
		BrandQueryService:    brandQueryService,
		BrandDeleteService:   brandDeleteService,
	})

	err = listenAndServeGracefully(server)
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("server shutdown: %v", err)
		return
	}
	log.Println("server shutdown")

}

func listenAndServeGracefully(hs *httpserver.Server) error {
	errChan := make(chan error, 2)
	defer close(errChan)

	timeout := time.Duration(60) * time.Second

	// http server
	go func() {
		err := hs.Start()
		errChan <- err
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer close(quit)

	select {
	case srvErr := <-errChan: // 에러 발생으로 인한 서버 종료
		return srvErr
	case <-quit: // 서버 종료 시그널 수신
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// http server shutdown
		err := hs.Stop(ctx)

		return err
	}
}

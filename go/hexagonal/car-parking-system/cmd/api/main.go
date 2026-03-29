package main

import (
	"car-parking-system/internal/adapters/in/httpfiber"
	"car-parking-system/internal/adapters/out/postgresql"
	"car-parking-system/internal/config"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	// --- start DB adapter (outbound) ---
	dbCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pg, err := postgresql.New(dbCtx, cfg.PostgreSQL)
	if err != nil {
		log.Fatalf("postgres init: %v", err)
	}
	defer func() {
		_ = pg.Close()
	}()

	// --- build services/usecases (application layer) ---
	// Example (pseudo):
	// userSvc := services.NewUserService(pg.UserRepo())

	// --- build HTTP adapter (inbound) ---
	app := httpfiber.NewApp(httpfiber.Options{
		// UserService: userSvc,
		// Logger: ...
	})

	// --- run server ---
	go func() {
		if err := app.Listen(cfg.HTTP.ListenAddr); err != nil {
			// On shutdown Fiber may return an error; treat it as info unless you want strict handling.
			log.Printf("fiber listen: %v", err)
		}
	}()

	// --- graceful shutdown ---
	sigCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	<-sigCtx.Done()

	shutdownCtx, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel2()

	if err := app.ShutdownWithContext(shutdownCtx); err != nil {
		log.Printf("fiber shutdown: %v", err)
	}

	// Close DB pool after server stops accepting new requests
	if err := pg.Close(); err != nil {
		log.Printf("postgres close: %v", err)
	}

	log.Println("shutdown complete")
}

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

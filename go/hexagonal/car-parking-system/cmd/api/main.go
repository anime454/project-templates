package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/anime454/project-templates/go/hexagonal/car-parking-system/internal/adapters/in/httpfiber"
	"github.com/anime454/project-templates/go/hexagonal/car-parking-system/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	// --- start DB adapter (outbound) ---
	dbCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("gorm init: %v", err)
	}
	pg := db.WithContext(dbCtx)

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
	sqlDB, err := pg.DB()
	if err != nil {
		log.Printf("gorm get db: %v", err)
	} else {
		if err := sqlDB.Close(); err != nil {
			log.Printf("gorm close: %v", err)
		}
	}

	log.Println("shutdown complete")
}

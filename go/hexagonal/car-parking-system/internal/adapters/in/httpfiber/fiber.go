package httpfiber

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

type Options struct {
	// Put your inbound dependencies here, typically application ports (interfaces).
	// Example:
	// UserService ports.in.UserService
}

func NewApp(opts Options) *fiber.App {

	app := fiber.New(fiber.Config{
		// Keep defaults unless you need custom error handling, timeouts, etc.
	})

	// Middleware (framework-specific => inbound adapter)
	app.Use(recover.New())
	// app.Use(requestid.New())

	registerRoutes(app, opts)

	return app
}

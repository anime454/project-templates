package httpfiber

import "github.com/gofiber/fiber/v3"

func registerRoutes(app *fiber.App, opts Options) {

	// Health check that doesn't touch core
	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("ok")
	})

	apiV1 := app.Group("/api/v1")

	vehicle := apiV1.Group("/vehicles")

	vehicle.Post("/in", func(c fiber.Ctx) error { return nil })
	vehicle.Post("/out", func(c fiber.Ctx) error { return nil })

}

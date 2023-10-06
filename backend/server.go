package main

import (
	endpoints "backend/endpoints"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Crea una nueva instancia de Fiber
	app := fiber.New()

	// Configura el CORS una vez
	app.Use(cors.New())

	// Configura tus rutas
	app.Post("/parser", endpoints.Analyzer())

	// Escucha en el puerto 5000
	app.Listen(":5000")
}

package routes

import (
	"github.com/akhiltn/go-link/internal/api"
	"github.com/gofiber/fiber/v2"
)

func AppRouteInit(app *fiber.App) {
	app.Get("/allkv", api.GetAllKV)
	app.Get("/:key", api.ResolveShortURL)
	app.Delete("/:key", api.DeleteShortURL)
	app.Post("/", api.CreateShortURL)
}

package handler

import (
	"cmd/internal/client"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Client client.Client
}

func (h Handler) Register(app *fiber.App) Handler {
	app.Post("/api/send", h.Publish)
	app.Get("/api/catch", h.Subscribe)

	return h
}
package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ko-ding-in/go-boilerplate/internal/appctx"
	"net/http"
)

type router struct {
	cfg   *appctx.Config
	fiber *fiber.App
}

func NewRouter(cfg *appctx.Config, fiber *fiber.App) Router {
	return &router{cfg: cfg, fiber: fiber}
}

func (rtr *router) Route() {
	rtr.fiber.Get("/ruok", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(map[string]string{"message": "Perfectly fine"})
	})
}

package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ko-ding-in/go-boilerplate/internal/appctx"
	"github.com/ko-ding-in/go-boilerplate/internal/bootstrap"
	"github.com/ko-ding-in/go-boilerplate/internal/controller"
	"github.com/ko-ding-in/go-boilerplate/internal/controller/contract"
)

type router struct {
	cfg   *appctx.Config
	fiber *fiber.App
}

func NewRouter(cfg *appctx.Config, fiber *fiber.App) Router {
	bootstrap.RegistryLogger(cfg)
	return &router{cfg: cfg, fiber: fiber}
}

func (rtr *router) handle(hfn httpHandleFunc, ctrl contract.Controller) fiber.Handler {
	return func(c *fiber.Ctx) error {
		resp := hfn(c, ctrl, rtr.cfg)
		return rtr.response(c, resp)
	}
}

func (rtr *router) response(c *fiber.Ctx, resp appctx.Response) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(resp.StatusCode).Send(resp.Byte())
}

func (rtr *router) Route() {
	// controllers
	controllers := controller.NewController(&contract.Dependency{
		Core: contract.Core{
			Cfg: rtr.cfg,
		},
	})

	rtr.fiber.Get("/ruok", rtr.handle(
		HttpRequest,
		controllers.HealthCheck.Liveness,
	))
}

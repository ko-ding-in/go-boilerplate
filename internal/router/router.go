package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ko-ding-in/go-boilerplate/internal/appctx"
	"github.com/ko-ding-in/go-boilerplate/internal/bootstrap"
	"github.com/ko-ding-in/go-boilerplate/internal/controller"
	"github.com/ko-ding-in/go-boilerplate/internal/controller/contract"
	"github.com/ko-ding-in/go-boilerplate/internal/dependencies"
	"github.com/ko-ding-in/go-boilerplate/internal/middleware"
	"github.com/ko-ding-in/go-boilerplate/internal/providers"
	"github.com/ko-ding-in/go-boilerplate/internal/repositories"
	"github.com/ko-ding-in/go-boilerplate/internal/services"
	"github.com/ko-ding-in/go-boilerplate/pkg/logger"
	"net/http"
	"runtime/debug"
)

type router struct {
	cfg   *appctx.Config
	fiber *fiber.App
}

func NewRouter(cfg *appctx.Config, fiber *fiber.App) Router {
	bootstrap.RegistryLogger(cfg)
	return &router{cfg: cfg, fiber: fiber}
}

func (rtr *router) handleRoute(hfn httpHandleFunc, ctrl contract.Controller, mdws ...middleware.Func) fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if err := recover(); err != nil {
				logger.Error(logger.MessageFormat("got panic error: %v", err))
				logger.Error(logger.MessageFormat("got panic: stack trace: %v", string(debug.Stack())))
				res := *appctx.NewResponse().
					WithStatusCode(http.StatusInternalServerError).
					WithMessage(http.StatusText(http.StatusInternalServerError))

				_ = rtr.response(c, res)
			}
		}()

		if rm := middleware.FilterFunc(rtr.cfg, c, mdws); rm.StatusCode != fiber.StatusOK {
			// return response base on middleware
			res := *appctx.NewResponse().
				WithCode(rm.Code).
				WithStatusCode(rm.StatusCode).
				WithMessage(rm.Message).
				WithErrors(rm.Errors)
			return rtr.response(c, res)
		}

		resp := hfn(c, ctrl, rtr.cfg)
		return rtr.response(c, resp)
	}
}

func (rtr *router) handleGroup(mdws ...middleware.Func) fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if err := recover(); err != nil {
				logger.Error(logger.MessageFormat("got panic error: %v", err))
				logger.Error(logger.MessageFormat("got panic: stack trace: %v", string(debug.Stack())))
				res := *appctx.NewResponse().
					WithStatusCode(http.StatusInternalServerError).
					WithMessage(http.StatusText(http.StatusInternalServerError))

				_ = rtr.response(c, res)
			}
		}()

		if rm := middleware.FilterFunc(rtr.cfg, c, mdws); rm.StatusCode != fiber.StatusOK {
			// return response base on middleware
			res := *appctx.NewResponse().
				WithCode(rm.Code).
				WithStatusCode(rm.StatusCode).
				WithMessage(rm.Message).
				WithErrors(rm.Errors)
			return rtr.response(c, res)
		}

		return c.Next()
	}
}

func (rtr *router) response(c *fiber.Ctx, resp appctx.Response) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(resp.StatusCode).Send(resp.Byte())
}

func (rtr *router) Route() {
	rtr.fiber.Use(rtr.handleGroup(
		middleware.Injector,
	))

	dep := dependencies.NewDependency(rtr.cfg)
	repo := repositories.NewRepository(dep)
	provider := providers.NewProvider(rtr.cfg)
	svcs := services.NewService(&services.Dependency{
		Repository: repo,
		Provider:   provider,
	})

	// controllers
	controllers := controller.NewController(&contract.Dependency{
		Core: contract.Core{
			Cfg:      rtr.cfg,
			Services: svcs,
		},
	})

	rtr.fiber.Get("/ruok", rtr.handleRoute(
		HttpRequest,
		controllers.HealthCheck.Liveness,
	))
}

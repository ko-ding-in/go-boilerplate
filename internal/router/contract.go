package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ko-ding-in/go-boilerplate/internal/appctx"
	"github.com/ko-ding-in/go-boilerplate/internal/controller/contract"
)

type httpHandleFunc func(*fiber.Ctx, contract.Controller, *appctx.Config) appctx.Response

func HttpRequest(xCtx *fiber.Ctx, svc contract.Controller, conf *appctx.Config) appctx.Response {
	data := appctx.Data{
		Ctx:    xCtx,
		Config: conf,
	}
	return svc.Serve(data)
}

type Router interface {
	Route()
}

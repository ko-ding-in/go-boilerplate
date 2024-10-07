package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ko-ding-in/go-boilerplate/internal/appctx"
)

type Func func(xCtx *fiber.Ctx, conf *appctx.Config) appctx.Response

// FilterFunc is a iterator resolver in each middleware registered
func FilterFunc(conf *appctx.Config, xCtx *fiber.Ctx, mfs []Func) appctx.Response {
	// Initiate postive case
	var response = appctx.Response{StatusCode: fiber.StatusOK}
	for _, mf := range mfs {
		if response = mf(xCtx, conf); response.StatusCode != fiber.StatusOK {
			return response
		}
	}

	return response
}

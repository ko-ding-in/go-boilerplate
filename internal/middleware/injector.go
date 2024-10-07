package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/ko-ding-in/go-boilerplate/internal/appctx"
	"github.com/ko-ding-in/go-boilerplate/internal/consts"
	"time"
)

func Injector(c *fiber.Ctx, _ *appctx.Config) appctx.Response {
	reqId := c.Get(consts.HeaderXRequestID)
	if reqId == "" {
		uid, _ := uuid.NewV7()
		reqId = uid.String()
	}
	// Set new id to response header
	c.Set(consts.HeaderXRequestID, reqId)

	c.UserContext()

	c.Context().SetUserValue(consts.ContextKeyRequestID, reqId)
	c.Context().SetUserValue(consts.ContextKeyStartTime, time.Now())
	c.Context().SetUserValue(consts.ContextKeyIP, c.IP())
	c.Context().SetUserValue(consts.ContextKeyPath, c.Path())
	c.Context().SetUserValue(consts.ContextKeyMethod, c.Method())

	return *appctx.NewResponse().WithStatusCode(fiber.StatusOK)
}

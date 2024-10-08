package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ko-ding-in/go-boilerplate/internal/consts"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"testing"
	"time"
)

func TestInjector(t *testing.T) {
	tests := []struct {
		name             string
		initialReqId     string
		expectNewReqId   bool
		expectReqIdInCtx bool
	}{
		{
			name:             "Request with existing X-Request-ID",
			initialReqId:     "existing-request-id",
			expectNewReqId:   false,
			expectReqIdInCtx: true,
		},
		{
			name:             "Request without X-Request-ID",
			initialReqId:     "",
			expectNewReqId:   true,
			expectReqIdInCtx: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new fiber context
			app := fiber.New()
			ctx := app.AcquireCtx(&fasthttp.RequestCtx{})

			// Set the initial request ID in the header, if any
			if tt.initialReqId != "" {
				ctx.Request().Header.Set(consts.HeaderXRequestID, tt.initialReqId)
			}

			// Call the Injector function
			resp := Injector(ctx, nil)

			// Verify response status is OK
			assert.Equal(t, fiber.StatusOK, resp.StatusCode)

			// Check if X-Request-ID is set correctly in the response headers
			reqId := ctx.Response().Header.Peek(consts.HeaderXRequestID)
			if tt.expectNewReqId {
				assert.NotEqual(t, tt.initialReqId, string(reqId), "A new request ID should be generated")
			} else {
				assert.Equal(t, tt.initialReqId, string(reqId), "Existing request ID should be kept")
			}

			// Verify request ID is present in context
			ctxReqId := ctx.UserContext().Value(consts.HeaderXRequestID)
			assert.NotNil(t, ctxReqId)
			assert.Equal(t, string(reqId), ctxReqId, "Request ID should be in the context")

			// Verify other context values (start time, IP, path, method)
			startTime := ctx.UserContext().Value(consts.ContextKeyStartTime)
			assert.IsType(t, time.Now(), startTime)

			ip := ctx.UserContext().Value(consts.ContextKeyIP)
			assert.NotNil(t, ip)

			path := ctx.UserContext().Value(consts.ContextKeyPath)
			assert.NotNil(t, path)

			method := ctx.UserContext().Value(consts.ContextKeyMethod)
			assert.NotNil(t, method)
		})
	}
}

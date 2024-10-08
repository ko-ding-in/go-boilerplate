package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ko-ding-in/go-boilerplate/internal/appctx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/valyala/fasthttp"
	"testing"
)

// Mock middleware function
type MockFunc struct {
	mock.Mock
}

func (m *MockFunc) Handle(xCtx *fiber.Ctx, conf *appctx.Config) appctx.Response {
	args := m.Called(xCtx, conf)
	return args.Get(0).(appctx.Response)
}

func TestFilterFunc(t *testing.T) {
	mockConf := &appctx.Config{}

	tests := []struct {
		name         string
		middlewares  []Func
		expectedResp appctx.Response
	}{
		{
			name: "All middleware return StatusOK",
			middlewares: []Func{
				func(xCtx *fiber.Ctx, conf *appctx.Config) appctx.Response {
					return appctx.Response{StatusCode: fiber.StatusOK}
				},
				func(xCtx *fiber.Ctx, conf *appctx.Config) appctx.Response {
					return appctx.Response{StatusCode: fiber.StatusOK}
				},
			},
			expectedResp: appctx.Response{StatusCode: fiber.StatusOK},
		},
		{
			name: "One middleware returns non-StatusOK",
			middlewares: []Func{
				func(xCtx *fiber.Ctx, conf *appctx.Config) appctx.Response {
					return appctx.Response{StatusCode: fiber.StatusOK}
				},
				func(xCtx *fiber.Ctx, conf *appctx.Config) appctx.Response {
					return appctx.Response{StatusCode: fiber.StatusBadRequest}
				},
			},
			expectedResp: appctx.Response{StatusCode: fiber.StatusBadRequest},
		},
		{
			name:         "No middlewares provided",
			middlewares:  nil,
			expectedResp: appctx.Response{StatusCode: fiber.StatusOK},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new fiber app and context
			app := fiber.New()
			req := app.AcquireCtx(&fasthttp.RequestCtx{}) // Use an actual request context
			defer app.ReleaseCtx(req)                     // Ensure proper resource cleanup

			// Run the FilterFunc with the test case data
			resp := FilterFunc(mockConf, req, tt.middlewares)

			// Assert the result
			assert.Equal(t, tt.expectedResp.StatusCode, resp.StatusCode)
		})
	}
}

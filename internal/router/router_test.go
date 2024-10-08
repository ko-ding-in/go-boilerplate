package router

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ko-ding-in/go-boilerplate/internal/appctx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockController is a mock implementation of the contract.Controller interface
type MockContractController struct {
	mock.Mock
}

func (m *MockContractController) EventName() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockContractController) Serve(data appctx.Data) appctx.Response {
	args := m.Called(data)
	return args.Get(0).(appctx.Response)
}

// TestHttpRequest tests the HttpRequest function in the router package
func TestHttpRequest(t *testing.T) {
	tests := []struct {
		name       string
		controller *MockContractController
		data       appctx.Data
		expected   appctx.Response
	}{
		{
			name:       "Successful response",
			controller: &MockContractController{},
			data: appctx.Data{
				Ctx:    nil, // Normally you would create a real *fiber.Ctx here
				Config: &appctx.Config{},
			},
			expected: appctx.Response{
				StatusCode: http.StatusOK,
				Message:    "Success",
			},
		},
		{
			name:       "Controller returns error response",
			controller: new(MockContractController),
			data: appctx.Data{
				Ctx:    nil,
				Config: &appctx.Config{},
			},
			expected: appctx.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    "Internal Server Error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.controller.On("Serve", tt.data).Return(tt.expected)

			resp := HttpRequest(nil, tt.controller, tt.data.Config)

			assert.Equal(t, tt.expected.StatusCode, resp.StatusCode)
			assert.Equal(t, tt.expected.Message, resp.Message)

			tt.controller.AssertExpectations(t)
		})
	}
}

func TestRoute(t *testing.T) {
	tests := []struct {
		name            string
		expectedCode    int
		expectedMessage string // description of the test case
		route           string
		mockController  *MockContractController
	}{
		{
			name:            "Liveness check returns success",
			expectedCode:    http.StatusOK,
			expectedMessage: "Perfectly Fine",
			route:           "/ruok",
			mockController:  &MockContractController{},
		},
		{
			name:            "Liveness check returns not found",
			expectedCode:    http.StatusNotFound,
			expectedMessage: "Internal Server Error",
			route:           "/ruk",
			mockController:  &MockContractController{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			cfg := &appctx.Config{
				App: appctx.App{Port: 9000},
			} // Mock your config as needed

			// Create a new router instance
			r := NewRouter(cfg, app)

			// Mock the controller behavior
			if tt.name == "Liveness check returns success" {
				tt.mockController.On("Serve", mock.Anything).Return(appctx.Response{
					StatusCode: http.StatusOK,
					Message:    "Perfectly Fine",
				})
			} else {
				tt.mockController.On("Serve", mock.Anything).Return(appctx.Response{
					StatusCode: http.StatusInternalServerError,
					Message:    "Internal Server Error",
				})
			}

			// Set up the route
			r.Route()

			req := httptest.NewRequest("GET", tt.route, nil)

			// Perform the request plain with the app,
			// the second argument is a request latency
			// (set to -1 for no latency)
			resp, _ := app.Test(req)

			// Verify, if the status code is as expected
			assert.Equal(t, tt.expectedCode, resp.StatusCode)
		})
	}
}

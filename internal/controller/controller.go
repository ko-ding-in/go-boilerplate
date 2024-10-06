package controller

import (
	"github.com/ko-ding-in/go-boilerplate/internal/controller/contract"
	healthcheck "github.com/ko-ding-in/go-boilerplate/internal/controller/health_check"
)

type Controller struct {
	HealthCheck *healthcheck.HealthCheck
}

func NewController(dep *contract.Dependency) *Controller {
	return &Controller{
		HealthCheck: healthcheck.NewHealthCheck(),
	}
}

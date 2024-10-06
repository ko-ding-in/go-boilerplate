package health_check

import "github.com/ko-ding-in/go-boilerplate/internal/controller/contract"

type HealthCheck struct {
	Liveness contract.Controller
}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{
		Liveness: newLivenessController(),
	}
}

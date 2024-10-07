package services

import (
	"github.com/ko-ding-in/go-boilerplate/internal/providers"
	"github.com/ko-ding-in/go-boilerplate/internal/repositories"
)

type (
	Services struct {
	}

	Dependency struct {
		Repository *repositories.Repository
		Provider   *providers.Provider
	}
)

func NewService(dep *Dependency) *Services {
	return &Services{}
}

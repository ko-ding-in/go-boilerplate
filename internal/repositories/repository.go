package repositories

import "github.com/ko-ding-in/go-boilerplate/internal/dependencies"

type Repository struct {
}

func NewRepository(dep *dependencies.Dependency) *Repository {
	return &Repository{}
}

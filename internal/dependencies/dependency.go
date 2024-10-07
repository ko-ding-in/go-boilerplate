package dependencies

import "github.com/ko-ding-in/go-boilerplate/internal/appctx"

type Dependency struct {
}

func NewDependency(cfg *appctx.Config) *Dependency {
	return &Dependency{}
}

package server

import (
	"context"
	"github.com/ko-ding-in/go-boilerplate/internal/appctx"
)

type Server interface {
	Run(context.Context) error
	Done()
	Config() *appctx.Config
}

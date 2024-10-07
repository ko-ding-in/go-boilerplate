package contract

import (
	"github.com/ko-ding-in/go-boilerplate/internal/appctx"
	"github.com/ko-ding-in/go-boilerplate/internal/services"
)

type Core struct {
	Cfg      *appctx.Config
	Services *services.Services
}

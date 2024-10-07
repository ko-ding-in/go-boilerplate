package providers

import (
	"github.com/ko-ding-in/go-boilerplate/internal/appctx"
)

type Provider struct {
}

func NewProvider(cfg *appctx.Config) *Provider {
	return &Provider{}
}

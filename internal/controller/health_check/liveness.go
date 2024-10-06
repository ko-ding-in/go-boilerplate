package health_check

import (
	"github.com/ko-ding-in/go-boilerplate/internal/appctx"
	"github.com/ko-ding-in/go-boilerplate/internal/controller/contract"
)

type livenessController struct{}

func newLivenessController() contract.Controller {
	return &livenessController{}
}

func (ctrl *livenessController) Serve(data appctx.Data) appctx.Response {
	const (
		LivenessMessage = `Perfectly Fine`
	)
	return *appctx.NewResponse().
		WithMessage(LivenessMessage)
}

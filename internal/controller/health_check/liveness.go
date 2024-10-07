package health_check

import (
	"github.com/ko-ding-in/go-boilerplate/internal/appctx"
	"github.com/ko-ding-in/go-boilerplate/internal/controller/contract"
	"github.com/ko-ding-in/go-boilerplate/pkg/logger"
)

type livenessController struct{}

func newLivenessController() contract.Controller {
	return &livenessController{}
}

func (ctrl *livenessController) eventName() string {
	return "controller.liveness"
}

func (ctrl *livenessController) Serve(data appctx.Data) appctx.Response {
	const (
		LivenessMessage = `Perfectly Fine`
	)

	var (
		lf = logger.NewFields(
			logger.EventName(ctrl.eventName()),
		)
		ctx = data.Ctx.UserContext()
	)

	logger.InfoWithContext(ctx, "Liveness Check", lf...)

	return *appctx.NewResponse().
		WithMessage(LivenessMessage)
}

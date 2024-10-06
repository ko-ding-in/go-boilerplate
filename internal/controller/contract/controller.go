package contract

import "github.com/ko-ding-in/go-boilerplate/internal/appctx"

type Controller interface {
	Serve(appctx.Data) appctx.Response
}

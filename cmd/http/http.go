package http

import (
	"context"
	"github.com/ko-ding-in/go-boilerplate/internal/server"
	"github.com/ko-ding-in/go-boilerplate/pkg/logger"
)

func Start(ctx context.Context) {
	httpServer := server.NewHttpServer()
	defer httpServer.Done()

	logger.Info(logger.MessageFormat("starting %s services... %d", httpServer.Config().App.Name, httpServer.Config().App.Port))

	if err := httpServer.Run(ctx); err != nil {
		logger.Fatal(logger.MessageFormat("http httpServer start got error: %v", err))
	}
}

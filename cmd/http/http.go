package http

import (
	"context"
	"fmt"
	"github.com/ko-ding-in/go-boilerplate/internal/server"
	"log"
)

func Start(ctx context.Context) {
	httpServer := server.NewHttpServer()
	defer httpServer.Done()

	log.Println(fmt.Sprintf("starting %s services... %d", "kodingin-boilerplate", 9990))

	if err := httpServer.Run(ctx); err != nil {
		log.Fatal(fmt.Sprintf("http httpServer start got error: %v", err))
	}
}

package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/ko-ding-in/go-boilerplate/internal/appctx"
	"log"
	"net/http"
	"time"
)

type httpServer struct {
	config *appctx.Config
}

func NewHttpServer() Server {
	return &httpServer{
		config: appctx.NewConfig(),
	}
}

func (s *httpServer) Run(ctx context.Context) error {
	var err error

	server := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", s.config.App.Port),
		ReadTimeout:  s.config.App.ReadTimeout,
		WriteTimeout: s.config.App.WriteTimeout,
	}

	go func() {
		err = server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("http server error: ", err)
		}
	}()

	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctxShutDown)
	if err != nil {
		log.Fatal("http server shutdown got error: ", err)
	}

	log.Println("server exited properly")

	if errors.Is(err, http.ErrServerClosed) {
		err = nil
	}

	return err
}

func (s *httpServer) Done() {
	log.Println("service has stopped")
}

func (s *httpServer) Config() *appctx.Config {
	return s.config
}

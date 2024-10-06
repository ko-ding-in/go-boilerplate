package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"
)

type httpServer struct {
}

func NewHttpServer() Server {
	return &httpServer{}
}

func (s *httpServer) Run(ctx context.Context) error {
	var err error

	server := &http.Server{
		Addr:         "0.0.0.0:9990",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
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

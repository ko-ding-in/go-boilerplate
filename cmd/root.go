package cmd

import (
	"context"
	"github.com/ko-ding-in/go-boilerplate/cmd/http"
	"github.com/ko-ding-in/go-boilerplate/pkg/logger"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	rootCmd := &cobra.Command{}
	ctx, cancel := context.WithCancel(context.Background())
	logger.SetJSONFormatter()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		cancel()
	}()

	httpCmd := &cobra.Command{
		Use:   "http",
		Short: "Start HTTP server",
		Run: func(cmd *cobra.Command, args []string) {
			http.Start(ctx)
		},
	}

	rootCmd.AddCommand(httpCmd)
	if err := rootCmd.Execute(); err != nil {
		logger.Fatal(err)
	}
}

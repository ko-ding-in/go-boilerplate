package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Start() {
	rootCmd := &cobra.Command{}
	_, cancel := context.WithCancel(context.Background())

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
			log.Println("Starting HTTP server")
		},
	}

	rootCmd.AddCommand(httpCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

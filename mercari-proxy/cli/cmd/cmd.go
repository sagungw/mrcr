package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sagungw/mercari/api/handler"
	"sagungw/mercari/api/route"
	"sagungw/mercari/core/cache"
	"sagungw/mercari/core/config"
	"sagungw/mercari/core/indoarea"
	"sync"
	"syscall"
	"time"

	"github.com/sagungw/gotrunks/log"
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
	}

	ServerCmd = &cobra.Command{
		Use:  "server",
		RunE: serverRunE,
	}
)

func serverRunE(cmd *cobra.Command, args []string) error {
	cache, err := cache.NewRedisCache(config.RedisAddress())
	if err != nil {
		return err
	}

	indoareaService := indoarea.NewCacheAwareClient(cache, 24*time.Hour, indoarea.NewClientStub())

	h := handler.Handler{IndoareaService: indoareaService}
	route := route.RouteHandler(h)

	srv := http.Server{
		Addr:    ":8181",
		Handler: route,
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

		s := <-sigCh
		log.Infof("Got signal %v, attempting graceful shutdown", s)

		if err := srv.Shutdown(context.Background()); err != nil {
			log.Errorf("HTTP server Shutdown: %v", err)
		}
	}()

	log.Info("HTTP server is running on port 8181")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	wg.Wait()
	return nil
}

func init() {
	RootCmd.AddCommand(ServerCmd)
}

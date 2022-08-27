package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sagungw/mercari/api/handler"
	"sagungw/mercari/api/route"
	"sagungw/mercari/core/config"
	"sagungw/mercari/core/datastore"
	"sagungw/mercari/core/service"
	"sync"
	"syscall"

	"github.com/sagungw/gotrunks/log"
	"github.com/spf13/cobra"
)

var (
	ServerCmd = &cobra.Command{
		Use:  "server",
		RunE: serverRunE,
	}
)

func serverRunE(cmd *cobra.Command, args []string) error {
	err := datastore.Open(config.DatabaseAddress())
	if err != nil {
		log.Fatal(err)
	}

	userDatastore := datastore.NewUserDatastore()
	loginHistoryDatastore := datastore.NewLoginHistoryDatastore()

	userService := service.NewUserService(userDatastore, loginHistoryDatastore)
	loginHistoryService := service.NewLoginHistoryService(loginHistoryDatastore)

	h := handler.Handler{UserService: userService, LoginHistoryService: loginHistoryService}
	m := handler.Middleware{UserService: userService}
	route := route.RouteHandler(h, m)

	srv := http.Server{
		Addr:    ":8282",
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

		if err := datastore.Close(); err != nil {
			log.Errorf("Shutting down datastore: %v", err)
		}

		if err := srv.Shutdown(context.Background()); err != nil {
			log.Errorf("HTTP server Shutdown: %v", err)
		}
	}()

	log.Info("HTTP server is running on port 8282")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	wg.Wait()
	return nil
}

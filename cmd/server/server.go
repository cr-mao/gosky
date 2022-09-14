package server

import (
	"context"
	"gosky/bootstrap"
	"gosky/cmd"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"

	httpServe "gosky/app/http"
	"gosky/infra/console"
	"gosky/infra/logger"
)

var ServeCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		if err := Run(cmd.Context()); err != nil {
			logger.ErrorString("cmd", "serve", err.Error())
			console.Exit("Unable to start server, error:" + err.Error())
		}
	},
	// rootCmd 的所有子命令都会执行以下代码
	PersistentPreRun: func(command *cobra.Command, args []string) {
		bootstrap.Bootstrap(cmd.Env)
	},
}

func Run(ctx context.Context) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//http start
	srv := httpServe.NewServe()
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			logger.ErrorString("http", "serve", err.Error())
		}
	}()
	<-quit
	// shutdown http server
	newCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(newCtx); err != nil {
		logger.WarnString("http", "shut_down", err.Error())
	}
	return nil
}

package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common/dependency"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/config"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/router"
)

func RunHttpServer() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()
	go func() {
		server := &http.Server{
			Addr:    fmt.Sprintf("192.168.1.4:%s", config.ServerConfig.Port),
			Handler: router.NewRouter(dependency.NewAppDependencies()),
		}
		slog.Info(
			fmt.Sprintf("Http Server start on port :%s",
				config.ServerConfig.Port,
			),
		)
		server.ListenAndServe()
	}()
	<-ctx.Done()
}

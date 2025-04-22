package app

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/router"
)

func init() {
	slog.SetDefault(
		slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	)
}

func RunHttpServer() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()
	go func() {
		server := &http.Server{
			Addr:    ":8000",
			Handler: router.NewRouter(),
		}
		slog.Info("Http Server start on port :8000")
		server.ListenAndServe()
	}()
	<-ctx.Done()
}

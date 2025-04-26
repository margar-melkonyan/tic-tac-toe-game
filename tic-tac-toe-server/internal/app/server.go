package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/router"
)

func init() {
	slog.SetDefault(
		slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	)
	err := godotenv.Load()
	if err != nil {
		slog.Error("can't load .env")
		panic(err)
	}
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
			Addr:    fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")),
			Handler: router.NewRouter(),
		}
		slog.Info(
			fmt.Sprintf("Http Server start on port :%s", os.Getenv("SERVER_PORT")),
		)
		server.ListenAndServe()
	}()
	<-ctx.Done()
}

package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/config"
	postgres_storage "github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/repository/postgres"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/router"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/storage/postgres"
)

func RunHttpServer() {
	db, _ := postgres.NewConnection()
	postgres_storage.NewRoomRepository(db)

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()
	go func() {
		server := &http.Server{
			Addr:    fmt.Sprintf(":%s", config.ServerConfig.Port),
			Handler: router.NewRouter(),
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

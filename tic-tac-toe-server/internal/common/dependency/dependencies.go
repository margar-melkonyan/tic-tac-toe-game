package dependency

import (
	"log/slog"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/handler/controller"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/repository"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/service"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/storage/postgres"
)

type AppDependencies struct {
	RoomHandler  controller.RoomHandler
	ScoreHandler controller.ScoreHandler
	UserHandler  controller.UserHandler
	AuthHandler  controller.AuthHandler
}

func NewAppDependencies() *AppDependencies {
	const op = "config.NewAppDependencides"
	store := postgres.Storage{
		ConnectionDriver: "postgres",
	}
	db, err := store.NewConnection()
	if err != nil {
		slog.With(op, err.Error())
		panic(err)
	}

	// repos
	roomRepo := repository.NewRoomRepository(db)
	scoreRepo := repository.NewScoreRepository(db)
	userRepo := repository.NewUserRepository(db)
	//services
	roomService := service.NewRoomService(roomRepo)
	scoreService := service.NewScoreService(scoreRepo)
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo)
	//handlers
	roomHandler := controller.NewRoomHandler(*roomService)
	scoreHandler := controller.NewScoreHandler(*scoreService)
	userHandler := controller.NewUserHandler(*userService)
	authHandler := controller.NewAuthHandler(*authService)

	return &AppDependencies{
		RoomHandler:  *roomHandler,
		ScoreHandler: *scoreHandler,
		UserHandler:  *userHandler,
		AuthHandler:  *authHandler,
	}
}

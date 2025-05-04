package dependency

import (
	"log/slog"

	http_handler "github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/handler/http"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/repository"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/service"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/storage/postgres"
)

type GlobalRepositories struct {
	UserRepository  repository.UserRepository
	ScoreRepository repository.ScoreRepository
}

type AppDependencies struct {
	RoomHandler  http_handler.RoomHandler
	ScoreHandler http_handler.ScoreHandler
	UserHandler  http_handler.UserHandler
	AuthHandler  http_handler.AuthHandler
	GlobalRepositories
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
	scoreService := service.NewScoreService(scoreRepo, userRepo)
	userService := service.NewUserService(userRepo, scoreRepo)
	authService := service.NewAuthService(userRepo)
	//handlers
	roomHandler := http_handler.NewRoomHandler(*roomService)
	scoreHandler := http_handler.NewScoreHandler(*scoreService)
	userHandler := http_handler.NewUserHandler(*userService)
	authHandler := http_handler.NewAuthHandler(*authService)

	return &AppDependencies{
		RoomHandler:  *roomHandler,
		ScoreHandler: *scoreHandler,
		UserHandler:  *userHandler,
		AuthHandler:  *authHandler,
		GlobalRepositories: GlobalRepositories{
			UserRepository:  userRepo,
			ScoreRepository: scoreRepo,
		},
	}
}

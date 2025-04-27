package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/config"
)

type Storage struct {
	ConnectionDriver string
}

func (stroage *Storage) NewConnection() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.ServerConfig.DbConfig.Host,
		config.ServerConfig.DbConfig.Port,
		config.ServerConfig.DbConfig.Username,
		config.ServerConfig.DbConfig.Password,
		config.ServerConfig.DbConfig.Name,
		config.ServerConfig.DbConfig.SSLMode,
	)

	db, err := sql.Open(stroage.ConnectionDriver, dsn)

	if err != nil {
		return nil, fmt.Errorf(
			"the applicaiton cannot open connection with %s",
			stroage.ConnectionDriver,
		)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf(
			"the application cannot connect to the database",
		)
	}

	return db, nil
}

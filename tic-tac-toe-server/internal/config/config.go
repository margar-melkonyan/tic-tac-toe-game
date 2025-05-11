package config

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/margar-melkonyan/tic-tac-toe-game/tic-tac-toe.git/internal/common"
)

var ServerConfig *common.ServerConfig

// init инициализирует конфигурацию сервера при старте приложения.
// Выполняет:
//   - Загрузку переменных окружения
//   - Инициализацию конфигурации, если она не была загружена ранее
//   - Настройку логгера с указанным уровнем логирования
func init() {
	mustLoadEnv()
	if ServerConfig == nil {
		NewConfig()
		setNewDefaultLogger(slog.Level(ServerConfig.LogLevel))
	}
}

// NewConfig создает и инициализирует новую конфигурацию сервера из переменных окружения.
// При ошибках парсинга числовых значений завершает работу приложения с panic.
//
// Загружаемые параметры:
//   - LOG_LEVEL: уровень логирования (число)
//   - BCRYPT_POWER: сложность хеширования bcrypt (число)
//   - SERVER_PORT: порт сервера
//   - DB_*: параметры подключения к БД
//   - JWT_*: параметры JWT токенов
//
// Возвращает:
//   - Инициализирует глобальную переменную ServerConfig
//
// Может вызвать panic при:
//   - Ошибках парсинга числовых параметров
//   - Отсутствии обязательных переменных окружения
func NewConfig() {
	logLevel, err := strconv.ParseInt(os.Getenv("LOG_LEVEL"), 10, 8)
	if err != nil {
		slog.Error(err.Error())
		panic(err.Error())
	}
	bcryptPower, err := strconv.ParseInt(os.Getenv("BCRYPT_POWER"), 10, 8)
	if err != nil {
		slog.Error(err.Error())
		panic(err.Error())
	}
	ServerConfig = &common.ServerConfig{
		Port:        os.Getenv("SERVER_PORT"),
		LogLevel:    int8(logLevel),
		BcryptPower: int(bcryptPower),
		DbConfig: common.DBConfig{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Name:     os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		},
		JWTConfig: common.JWTConfig{
			AccessTokenSecret: os.Getenv("JWT_ACCESS_TOKEN_SECRET"),
			AccessTokenTTL:    os.Getenv("JWT_ACCESS_TOKEN_TTL"),
		},
	}
}

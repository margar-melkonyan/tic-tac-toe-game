package common

type DBConfig struct {
	Username string
	Password string
	Name     string
	Host     string
	Port     string
	SSLMode  string
}

type JWTConfig struct {
	AccessTokenSecret string
	AccessTokenTTL    string
}

type ServerConfig struct {
	Port      string
	LogLevel  int8
	DbConfig  DBConfig
	JWTConfig JWTConfig
}

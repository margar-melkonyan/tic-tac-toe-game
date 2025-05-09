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
	LocalIP     string
	Port        string
	LogLevel    int8
	BcryptPower int
	DbConfig    DBConfig
	JWTConfig   JWTConfig
}

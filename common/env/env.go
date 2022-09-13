package env

type ServerEnv struct {
	PORT string `env:"SERVER_PORT"`
}

type DataBaseEnv struct {
	HOST     string `env:"POSTGRES_HOST"`
	USER     string `env:"POSTGRES_USER"`
	PORT     string `env:"POSTGRES_PORT"`
	DB_NAME  string `env:"POSTGRES_DB"`
	PASSWORD string `env:"POSTGRES_PASSWORD"`
}

type JwtEnv struct {
	SECRET_KEY string `env:"SECRET_KEY"`
	ISSURE     string `env:"ISSURE"`
}

var Server ServerEnv
var DataBase DataBaseEnv
var JWT JwtEnv

func Load() {
	LoadEnv(&Server, &DataBase, &JWT)
}

package env

type ServerEnv struct {
	PORT string `env:"SERVER_PORT"`
}

type DataBaseEnv struct {
	HOST     string `env:"MYSQL_HOST"`
	USER     string `env:"MYSQL_USER"`
	PORT     string `env:"MYSQL_PORT"`
	DB_NAME  string `env:"MYSQL_DATABASE"`
	PASSWORD string `env:"MYSQL_PASSWORD"`
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

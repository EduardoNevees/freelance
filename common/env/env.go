package env

type ServerEnv struct {
	PORT string `env:"SERVER_PORT"`
}

type DataBaseEnv struct {
	HOST     string `env:"MYSQL_HOST"`
	USER     string `env:"MYSQL_USER"`
	SQL_PORT string `env:"MYSQL_PORT"`
	DB_NAME  string `env:"MYSQL_DB_NAME"`
	PASSWORD string `env:"MYSQL_PASSWORD"`
}

var Server ServerEnv
var DataBase DataBaseEnv

func Load() {
	LoadEnv(&Server, &DataBase)
}

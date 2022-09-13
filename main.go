package main

import (
	"github.com/EduardoNevesRamos/frelancer.git/adapters/api"
	"github.com/EduardoNevesRamos/frelancer.git/adapters/infrastructure/db"
	"github.com/EduardoNevesRamos/frelancer.git/common/env"
)

func main() {
	env.Load()
	db.Start()

	server := api.New()
	server.Run()

}

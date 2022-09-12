package api

import "github.com/EduardoNevesRamos/frelancer.git/common/env"

func New() Server {
	server := NewServer()

	server.SetPort(env.Server.PORT)
	server.SetHealthResponse("Health Response Okay")
	return *server
}

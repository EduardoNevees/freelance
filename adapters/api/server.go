package api

import (
	"log"
	"net/http"

	"github.com/EduardoNevesRamos/frelancer.git/common/env"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port          string
	healthMessage string
	serverEngine  *gin.Engine
}

func NewServer() *Server {
	return &Server{
		port:          env.Server.PORT,
		healthMessage: "",
		serverEngine:  gin.Default(),
	}
}

func (s *Server) SetPort(port string) {
	s.port = port
}

func (s *Server) SetHealthResponse(message string) {
	s.healthMessage = message
}

func (s *Server) healthRoute(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, s.healthMessage)
}

func (s *Server) Run() {
	router := ConfigRoutes(s.serverEngine)

	s.serverEngine.GET("/health", s.healthRoute)

	var address string = ":" + s.port

	s.serverEngine.Run(address)

	log.Printf("Server running at port: %v", s.port)
	log.Fatal(router.Run(":" + s.port))
}

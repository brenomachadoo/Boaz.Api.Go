package server

import (
	"fmt"
	"log"
	"os"

	"bmachado/Boaz.Api.Go/cmd/boazApi/api/server/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServerCompany() Server {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file: " + err.Error())
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	newServer := Server{
		port:   port,
		server: gin.New(),
	}

	return newServer
}

func (s *Server) RunServerCompany() {
	router := routers.ContigureRoutesCompany(s.server)

	log.Print("server is running at port: ", s.port)
	log.Fatal(router.Run(":" + s.port))
}

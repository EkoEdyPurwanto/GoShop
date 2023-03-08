package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Server struct {
	DB     *gorm.DB
	Router *echo.Echo
}

func (server *Server) Initialize() {
	fmt.Println("WELCOME TO GoShop")

	server.Router = echo.New()
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to Port %s", addr)

	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func Run() {
	var server = Server{}

	server.Initialize()
	server.Run(":9000")
}

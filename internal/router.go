package internal

import (
	"GoShop/internal/controllers"
)

func (server *Server) InitializeRouter() {
	server.Router.GET("/", controllers.Home)
}

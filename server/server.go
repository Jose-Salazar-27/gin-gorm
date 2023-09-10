package server

import (
	"github.com/bricsport/handler"
	"github.com/bricsport/repository"
	"github.com/bricsport/router"
	"github.com/bricsport/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	Router  *gin.Engine
	Handler handler.Handler
}

func NewServer(r *gin.Engine, tx *gorm.DB) *Server {
	repositories := repository.InitRepositories(tx)
	services := service.InitServices(*repositories)
	handlers := handler.InitHandlers(*services)

	return &Server{
		Router:  r,
		Handler: *handlers,
	}

}

func (srvr *Server) Init() {
	// Load routes
	router.LoadUserRoutes(srvr.Router, srvr.Handler.Userhandler)
}

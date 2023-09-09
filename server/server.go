package server

import (
	"github.com/bricsport/handler"
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	Router  *gin.Engine
	Handler handler.Handler
}

func (as *APIServer) Init() {
	
}
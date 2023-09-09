package handler

import (
	"net/http"

	"github.com/bricsport/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Userhandler *userHandler
}

func InitHandlers(services service.Service) *Handler {
	return &Handler{
		Userhandler: &userHandler{Service: *services.UserService},
	}
}

// func Handle

func (h *Handler) HandleUser(ctx *gin.Context) {

	switch ctx.Request.Method {
	case http.MethodGet:
		ctx.String(200, "ok")
	default:
		ctx.String(400, "mehtod not allowed")
	}

}

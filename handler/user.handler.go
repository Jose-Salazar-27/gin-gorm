package handler

import (
	"net/http"

	"github.com/bricsport/common/exceptions"
	"github.com/bricsport/entity"
	"github.com/bricsport/service"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Create(ctx *gin.Context) (*Payload, *exceptions.Exception)
}

type userHandler struct {
	Service service.UserService
}

func (hdlr *userHandler) Create(ctx *gin.Context) (*Payload, *exceptions.Exception) {
	usr := entity.User{}
	if parseErr := ctx.BindJSON(&usr); parseErr != nil {
		return nil, exceptions.Throw(http.StatusBadRequest, "ERROR BIDING JSON: "+parseErr.Error())
	}

	result, err := hdlr.Service.CreateUser(&usr)

	if err != nil {
		return nil, err
	}

	return &Payload{Code: 201, Data: result}, nil
}

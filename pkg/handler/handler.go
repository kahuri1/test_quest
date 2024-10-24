package handler

import (
	"github.com/gin-gonic/gin"
)

type gRPCProxy interface {
}

type Handler struct {
	service gRPCProxy
}

func Newhandler(service gRPCProxy) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	return r
}

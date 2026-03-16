package service

import "github.com/gin-gonic/gin"

type WsService interface {
	AddClient(*gin.Context) error
}

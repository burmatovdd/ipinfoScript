package apiMethods

import "github.com/gin-gonic/gin"

type Service struct {
	method *Api
}

type Api interface {
	GetInfo(c *gin.Context)
}

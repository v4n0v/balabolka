package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterInfoService(router *gin.Engine) {
	messageRoute := router.Group("/api/info")
	{
		messageRoute.GET("/echo", echo)
	}
}

func echo(c *gin.Context) {
	c.String(http.StatusOK, c.Query("msg"))
}

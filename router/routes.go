package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ivovilar1/goledger-challenge-besu/handler"
)

func initRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/set", handler.SetHandler)
		v1.GET("/get", handler.GetHandler)
		v1.POST("/sync", handler.SyncHandler)
		v1.GET("/check", handler.CheckHandler)
	}
}

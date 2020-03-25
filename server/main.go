package server

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/ncov", WhiteList)

	r.Static("/uploads", "./uploads")

	return r
}

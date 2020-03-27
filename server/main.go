package server

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/ncov", WhiteList)

	//获取静态资源
	r.Static("/uploads", "./uploads")
	r.Static("/static", "./static")

	return r
}

package server

import (
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/gin-gonic/gin"
	"ncov-report-manage-system-GO/pages"
)

func WhiteList(ctx *gin.Context)  {
	engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
		return pages.MainPage()
	})
}

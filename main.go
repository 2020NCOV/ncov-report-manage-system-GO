package main

import (
	_ "github.com/GoAdminGroup/go-admin/adapter/gin" // 引入适配器，必须引入，如若不引入，则需要自己定义
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql" // 引入对应数据库引擎
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/GoAdminGroup/go-admin/template"
	_ "github.com/GoAdminGroup/themes/adminlte" // 引入主题，必须引入，不然报错
	"ncov-report-manage-system-GO/controller/template/login"
	"ncov-report-manage-system-GO/model"
	"ncov-report-manage-system-GO/server"
)

func main() {
	//实例化一个路由
	r := server.Router()

	// 实例化一个GoAdmin引擎对象
	eng := engine.Default()

	//从json文件中读取配置
	cfg := config.ReadFromJson("./config.json")

	//增加网站图标
	cfg.CustomHeadHtml = template.HTML(`<link rel="icon" type="image/png" sizes="32x32" href="/static/logo.png">`)

	//页面动画
	cfg.Animation = config.PageAnimation{
		Type:     "fadeInUp",
		Duration: 0.9,
	}

	// 这里引入你需要管理的业务表配置，在model中生成一系列表，详见model/table.go
	adminPlugin := admin.NewAdmin(model.Generators)

	//也可单个添加业务表
	//adminPlugin.AddGenerator("user", model.GetUserTable)

	// 增加配置与插件，使用Use方法挂载到Web框架中
	if err := eng.AddConfig(cfg).AddPlugins(adminPlugin).Use(r); err != nil {
		panic(err)
	}

	//增加登录模块,不写时有默认的登录页面，可在本项目中自定义登录页面，在此处加入
	template.AddLoginComp(login.GetLoginComponent())

	//监听端口9033
	_ = r.Run(":9033")
}

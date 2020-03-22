package model

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	editType "github.com/GoAdminGroup/go-admin/template/types/table"
)

func GetWxMpBindInfoTable(ctx *context.Context) table.Table {

	wxMpBindInfoTable := table.NewDefaultTable(table.Config{
		Driver:     db.DriverMysql,
		CanAdd:     true,
		Editable:   true,
		Deletable:  true,
		Exportable: true,	//可导出数据
		Connection: table.DefaultConnectionName,
		PrimaryKey: table.PrimaryKey{
			Type: db.Int,
			Name: table.DefaultPrimaryKeyName,
		},
	})

	//SetFilterFormLayout()设置筛选布局一行有几个筛选选项
	info := wxMpBindInfoTable.GetInfo().SetFilterFormLayout(form.LayoutThreeCol)

	/*
	* 1. FieldFilterable()，可筛选，在该表顶端的筛选布局中会多出此项。默认只有主键可筛选
	* 2. FieldSortable(), 可排序，
	* 3. FieldEditAble(), 可编辑，editType.xxx多种类型
	* 4. AddColumn(),增加一列
	* 5.
	**/

	info.AddField("Id", "id", db.Int).FieldSortable()
	info.AddField("Org_id", "org_id", db.Int).FieldFilterable()
	info.AddField("Wx_uid", "wx_uid", db.Int).FieldFilterable()
	info.AddField("用户名", "username", db.Varchar).FieldFilterable().FieldEditAble()
	info.AddField("是否绑定", "isbind", db.Int).FieldEditAble(editType.Switch).FieldEditOptions(types.FieldOptions{
		{Value:"1", Text:"Y"},
		{Value:"0", Text:"N"},
	})

	//注：如果此方法报错，则是因为版本问题，使用go 包管理获取到的
	info.AddColumn("personality", func(value types.FieldModel) interface{} {
		return "handsome"
	})

	info.AddColumnButtons("see more", types.GetDefaultButton("see more", icon.Info,
		action.PopUp("/see/more/example", "see more", func(ctx *context.Context) (success bool, msg string, data interface{}) {
			return true, "ok", "<h1>Detail</h1><p>balabala</p>"
		})))

	info.AddField("绑定时间", "bind_date", db.Datetime)
	info.AddField("解绑时间", "unbind_date", db.Datetime)
	info.AddField("Remark", "remark", db.Varchar)


	info.SetTable("wx_mp_bind_info").SetTitle("小程序绑定用户").SetDescription("小程序绑定用户信息")

	formList := wxMpBindInfoTable.GetForm()

	formList.AddField("Id", "id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("Org_id", "org_id", db.Int, form.Number)
	formList.AddField("Wx_uid", "wx_uid", db.Int, form.Number)
	formList.AddField("Username", "username", db.Varchar, form.Text)
	formList.AddField("Isbind", "isbind", db.Int, form.Number)
	formList.AddField("Bind_date", "bind_date", db.Datetime, form.Datetime)
	formList.AddField("Unbind_date", "unbind_date", db.Datetime, form.Datetime)
	formList.AddField("Remark", "remark", db.Varchar, form.Text)

	formList.SetTable("wx_mp_bind_info").SetTitle("Wx_mp_bind_info").SetDescription("Wx_mp_bind_info")

	return wxMpBindInfoTable
}

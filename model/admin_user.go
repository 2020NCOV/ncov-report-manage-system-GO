package model

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetAdminUserTable(ctx *context.Context) table.Table {

	adminUserTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := adminUserTable.GetInfo()

	info.AddField("Id", "id", db.Int).FieldSortable()
	info.AddField("Org_id", "org_id", db.Int).FieldFilterable()
	info.AddField("Dep_id", "dep_id", db.Int).FieldEditAble()
	info.AddField("Username", "username", db.Varchar)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Password", "password", db.Varchar)
	info.AddField("Role", "role", db.Int)
	info.AddField("Is_admin", "is_admin", db.Int)
	info.AddField("Is_del", "is_del", db.Int)
	info.AddField("Need_m_pass", "need_m_pass", db.Int)
	info.AddField("Remarks", "remarks", db.Varchar)
	info.AddField("Datetime", "datetime", db.Timestamp)

	info.SetTable("admin_user").SetTitle("Admin_user").SetDescription("Admin_user")

	formList := adminUserTable.GetForm()

	formList.AddField("Id", "id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("Org_id", "org_id", db.Int, form.Number)
	formList.AddField("Dep_id", "dep_id", db.Int, form.Number)
	formList.AddField("Username", "username", db.Varchar, form.Text)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Password", "password", db.Varchar, form.Password)
	formList.AddField("Role", "role", db.Int, form.Number)
	formList.AddField("Is_admin", "is_admin", db.Int, form.Number)
	formList.AddField("Is_del", "is_del", db.Int, form.Number)
	formList.AddField("Need_m_pass", "need_m_pass", db.Int, form.Number)
	formList.AddField("Remarks", "remarks", db.Varchar, form.Text)
	formList.AddField("Datetime", "datetime", db.Timestamp, form.Datetime)

	formList.SetTable("admin_user").SetTitle("Admin_user").SetDescription("Admin_user")

	return adminUserTable
}

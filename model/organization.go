package model

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetOrganizationTable(ctx *context.Context) table.Table {

	organizationTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := organizationTable.GetInfo()

	info.AddField("Id", "id", db.Int).FieldFilterable()
	info.AddField("Corpname", "corpname", db.Varchar)
	info.AddField("Corpname_full", "corpname_full", db.Varchar)
	info.AddField("Access_type", "access_type", db.Varchar)
	info.AddField("Template_code", "template_code", db.Varchar)
	info.AddField("Corp_code", "corp_code", db.Varchar)
	info.AddField("Type_corpname", "type_corpname", db.Varchar)
	info.AddField("Type_username", "type_username", db.Varchar)
	info.AddField("Admin_name", "admin_name", db.Varchar)
	info.AddField("Tel", "tel", db.Varchar)
	info.AddField("Is_del", "is_del", db.Int)
	info.AddField("Status", "status", db.Int)
	info.AddField("Remark", "remark", db.Varchar)
	info.AddField("Add_date", "add_date", db.Datetime)

	info.SetTable("organization").SetTitle("Organization").SetDescription("Organization")

	formList := organizationTable.GetForm()

	formList.AddField("Id", "id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("Corpname", "corpname", db.Varchar, form.Text)
	formList.AddField("Corpname_full", "corpname_full", db.Varchar, form.Text)
	formList.AddField("Access_type", "access_type", db.Varchar, form.Text)
	formList.AddField("Template_code", "template_code", db.Varchar, form.Text)
	formList.AddField("Corp_code", "corp_code", db.Varchar, form.Text)
	formList.AddField("Type_corpname", "type_corpname", db.Varchar, form.Text)
	formList.AddField("Type_username", "type_username", db.Varchar, form.Text)
	formList.AddField("Admin_name", "admin_name", db.Varchar, form.Text)
	formList.AddField("Tel", "tel", db.Varchar, form.Text)
	formList.AddField("Is_del", "is_del", db.Int, form.Number)
	formList.AddField("Status", "status", db.Int, form.Number)
	formList.AddField("Remark", "remark", db.Varchar, form.Text)
	formList.AddField("Add_date", "add_date", db.Datetime, form.Datetime)

	formList.SetTable("organization").SetTitle("Organization").SetDescription("Organization")

	return organizationTable
}

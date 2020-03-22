package model

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetOrgWhitelistTable(ctx *context.Context) table.Table{

	orgWhitelistTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := orgWhitelistTable.GetInfo()

	info.AddField("Id", "id", db.Int).FieldFilterable()
	info.AddField("Org_id", "org_id", db.Int)
	info.AddField("UserID", "userID", db.Varchar)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Gender", "gender", db.Varchar)
	info.AddField("Sub1_department_id", "sub1_department_id", db.Int)
	info.AddField("Sub2_department_id", "sub2_department_id", db.Int)
	info.AddField("Tag1", "tag1", db.Varchar)
	info.AddField("Tag2", "tag2", db.Varchar)
	info.AddField("Tag3", "tag3", db.Varchar)
	info.AddField("Tag4", "tag4", db.Varchar)
	info.AddField("Add_datetime", "add_datetime", db.Timestamp)
	info.AddField("Last_update_time", "last_update_time", db.Datetime)
	info.AddField("Add_remark", "add_remark", db.Varchar)
	info.AddField("Dep_name", "dep_name", db.Varchar)
	info.AddField("Report_id", "report_id", db.Int)
	info.AddField("Report_date", "report_date", db.Datetime)
	info.AddField("Status", "status", db.Int)

	info.SetTable("org_whitelist").SetTitle("Org_whitelist").SetDescription("Org_whitelist")

	formList := orgWhitelistTable.GetForm()

	formList.AddField("Id", "id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("Org_id", "org_id", db.Int, form.Number)
	formList.AddField("UserID", "userID", db.Varchar, form.Text)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Gender", "gender", db.Varchar, form.Text)
	formList.AddField("Sub1_department_id", "sub1_department_id", db.Int, form.Number)
	formList.AddField("Sub2_department_id", "sub2_department_id", db.Int, form.Number)
	formList.AddField("Tag1", "tag1", db.Varchar, form.Text)
	formList.AddField("Tag2", "tag2", db.Varchar, form.Text)
	formList.AddField("Tag3", "tag3", db.Varchar, form.Text)
	formList.AddField("Tag4", "tag4", db.Varchar, form.Text)
	formList.AddField("Add_datetime", "add_datetime", db.Timestamp, form.Datetime)
	formList.AddField("Last_update_time", "last_update_time", db.Datetime, form.Datetime)
	formList.AddField("Add_remark", "add_remark", db.Varchar, form.Text)
	formList.AddField("Dep_name", "dep_name", db.Varchar, form.Text)
	formList.AddField("Report_id", "report_id", db.Int, form.Number)
	formList.AddField("Report_date", "report_date", db.Datetime, form.Datetime)
	formList.AddField("Status", "status", db.Int, form.Number)

	formList.SetTable("org_whitelist").SetTitle("Org_whitelist").SetDescription("Org_whitelist")

	return orgWhitelistTable
}

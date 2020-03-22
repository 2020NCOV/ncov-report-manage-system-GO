package model

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetOrgDepTable(ctx *context.Context) table.Table {

	orgDepTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := orgDepTable.GetInfo()

	info.AddField("Id", "id", db.Int).FieldFilterable()
	info.AddField("Org_id", "org_id", db.Int)
	info.AddField("Dep_name", "dep_name", db.Varchar)
	info.AddField("Level", "level", db.Int)
	info.AddField("Is_del", "is_del", db.Int)
	info.AddField("Status", "status", db.Int)
	info.AddField("Remark", "remark", db.Varchar)
	info.AddField("Datetime", "datetime", db.Timestamp)

	info.SetTable("org_dep").SetTitle("Org_dep").SetDescription("Org_dep")

	formList := orgDepTable.GetForm()

	formList.AddField("Id", "id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("Org_id", "org_id", db.Int, form.Number)
	formList.AddField("Dep_name", "dep_name", db.Varchar, form.Text)
	formList.AddField("Level", "level", db.Int, form.Number)
	formList.AddField("Is_del", "is_del", db.Int, form.Number)
	formList.AddField("Status", "status", db.Int, form.Number)
	formList.AddField("Remark", "remark", db.Varchar, form.Text)
	formList.AddField("Datetime", "datetime", db.Timestamp, form.Datetime)

	formList.SetTable("org_dep").SetTitle("Org_dep").SetDescription("Org_dep")

	return orgDepTable
}

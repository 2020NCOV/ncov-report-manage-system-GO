package model

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetReportRecordDefaultTable(ctx *context.Context) table.Table {

	reportRecordDefaultTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := reportRecordDefaultTable.GetInfo()

	info.AddField("Id", "id", db.Bigint).FieldFilterable()
	info.AddField("Wxuid", "wxuid", db.Int)
	info.AddField("Org_id", "org_id", db.Int)
	info.AddField("Org_name", "org_name", db.Varchar)
	info.AddField("UserID", "userID", db.Varchar)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Is_return_school", "is_return_school", db.Int)
	info.AddField("Return_dorm_num", "return_dorm_num", db.Varchar)
	info.AddField("Return_time", "return_time", db.Date)
	info.AddField("Return_district_value", "return_district_value", db.Int)
	info.AddField("Return_district_path", "return_district_path", db.Varchar)
	info.AddField("Return_traffic_info", "return_traffic_info", db.Varchar)
	info.AddField("Current_district_value", "current_district_value", db.Int)
	info.AddField("Current_district_path", "current_district_path", db.Varchar)
	info.AddField("Current_health_value", "current_health_value", db.Int)
	info.AddField("Current_contagion_risk_value", "current_contagion_risk_value", db.Int)
	info.AddField("Current_temperature", "current_temperature", db.Float)
	info.AddField("Psy_status", "psy_status", db.Int)
	info.AddField("Psy_demand", "psy_demand", db.Int)
	info.AddField("Psy_knowledge", "psy_knowledge", db.Int)
	info.AddField("Return_company_date", "return_company_date", db.Varchar)
	info.AddField("Plan_company_date", "plan_company_date", db.Varchar)
	info.AddField("Template_code", "template_code", db.Varchar)
	info.AddField("Remarks", "remarks", db.Varchar)
	info.AddField("Time", "time", db.Timestamp)

	info.SetTable("report_record_default").SetTitle("Report_record_default").SetDescription("Report_record_default")

	formList := reportRecordDefaultTable.GetForm()

	formList.AddField("Id", "id", db.Bigint, form.Default).FieldNotAllowAdd()
	formList.AddField("Wxuid", "wxuid", db.Int, form.Number)
	formList.AddField("Org_id", "org_id", db.Int, form.Number)
	formList.AddField("Org_name", "org_name", db.Varchar, form.Text)
	formList.AddField("UserID", "userID", db.Varchar, form.Text)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Is_return_school", "is_return_school", db.Int, form.Number)
	formList.AddField("Return_dorm_num", "return_dorm_num", db.Varchar, form.Text)
	formList.AddField("Return_time", "return_time", db.Date, form.Datetime)
	formList.AddField("Return_district_value", "return_district_value", db.Int, form.Number)
	formList.AddField("Return_district_path", "return_district_path", db.Varchar, form.Text)
	formList.AddField("Return_traffic_info", "return_traffic_info", db.Varchar, form.Text)
	formList.AddField("Current_district_value", "current_district_value", db.Int, form.Number)
	formList.AddField("Current_district_path", "current_district_path", db.Varchar, form.Text)
	formList.AddField("Current_health_value", "current_health_value", db.Int, form.Number)
	formList.AddField("Current_contagion_risk_value", "current_contagion_risk_value", db.Int, form.Number)
	formList.AddField("Current_temperature", "current_temperature", db.Float, form.Text)
	formList.AddField("Psy_status", "psy_status", db.Int, form.Number)
	formList.AddField("Psy_demand", "psy_demand", db.Int, form.Number)
	formList.AddField("Psy_knowledge", "psy_knowledge", db.Int, form.Number)
	formList.AddField("Return_company_date", "return_company_date", db.Varchar, form.Text)
	formList.AddField("Plan_company_date", "plan_company_date", db.Varchar, form.Text)
	formList.AddField("Template_code", "template_code", db.Varchar, form.Text)
	formList.AddField("Remarks", "remarks", db.Varchar, form.Text)
	formList.AddField("Time", "time", db.Timestamp, form.Datetime)

	formList.SetTable("report_record_default").SetTitle("Report_record_default").SetDescription("Report_record_default")

	return reportRecordDefaultTable
}

package model

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "org_whitelist" => http://localhost:9033/admin/info/org_whitelist
//
// "admin_user" => http://localhost:9033/admin/info/admin_user
// "org_dep" => http://localhost:9033/admin/info/org_dep
// "report_record_default" => http://localhost:9033/admin/info/report_record_default
// "wx_mp_bind_info" => http://localhost:9033/admin/info/wx_mp_bind_info
//
// "organization" => http://localhost:9033/admin/info/organization
//
// example end
//
var Generators = map[string]table.Generator{
	"org_whitelist": GetOrgWhitelistTable,

	"admin_user":            GetAdminUserTable,
	"org_dep":               GetOrgDepTable,
	"report_record_default": GetReportRecordDefaultTable,
	"wx_mp_bind_info":       GetWxMpBindInfoTable,

	"organization": GetOrganizationTable,

	// generators end
}

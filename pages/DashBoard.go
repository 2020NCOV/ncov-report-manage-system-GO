package pages

import (
	tmpl "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/themes/adminlte/components/infobox"
)


func MainPage() (types.Panel, error) {
	components := tmpl.Default()
	colComp := components.Col()

	/**************************
	 * Info Box
	/**************************/

	infobox1 := infobox.New().
		SetText("CPU TRAFFIC").
		SetColor("aqua").
		SetNumber("100").
		SetIcon("ion-ios-gear-outline").
		GetContent()

	infobox2 := infobox.New().
		SetText("Likes").
		SetColor("red").
		SetNumber("1030.00<small>$</small>").
		SetIcon(icon.GooglePlus).
		GetContent()

	infobox3 := infobox.New().
		SetText("Sales").
		SetColor("green").
		SetNumber("760").
		SetIcon("ion-ios-cart-outline").
		GetContent()

	infobox4 := infobox.New().
		SetText("New Members").
		SetColor("yellow").
		SetNumber("2,349").
		SetIcon("ion-ios-people-outline"). // svg is ok
		GetContent()

	var size = types.SizeMD(3).SM(6).XS(12)
	infoboxCol1 := colComp.SetSize(size).SetContent(infobox1).GetContent()
	infoboxCol2 := colComp.SetSize(size).SetContent(infobox2).GetContent()
	infoboxCol3 := colComp.SetSize(size).SetContent(infobox3).GetContent()
	infoboxCol4 := colComp.SetSize(size).SetContent(infobox4).GetContent()


	row1 := components.Row().SetContent(infoboxCol1 + infoboxCol2 + infoboxCol3 + infoboxCol4).GetContent()

	//tableCol := colComp.SetSize(types.SizeMD(8)).SetContent(row1).GetContent()
	//
	//row5 := components.Row().SetContent(tableCol).GetContent()

	table := components.Table().SetType("table").SetInfoList([]map[string]types.InfoItem{
		{
			"Order ID":   {Content: "OR9842"},
			"Item":       {Content: "Call of Duty IV"},
			"Status":     {Content: "shipped"},
			"Popularity": {Content: "90%"},
		}, {
			"Order ID":   {Content: "OR9842"},
			"Item":       {Content: "Call of Duty IV"},
			"Status":     {Content: "shipped"},
			"Popularity": {Content: "90%"},
		}, {
			"Order ID":   {Content: "OR9842"},
			"Item":       {Content: "Call of Duty IV"},
			"Status":     {Content: "shipped"},
			"Popularity": {Content: "90%"},
		}, {
			"Order ID":   {Content: "OR9842"},
			"Item":       {Content: "Call of Duty IV"},
			"Status":     {Content: "shipped"},
			"Popularity": {Content: "90%"},
		},
	}).SetThead(types.Thead{
		{Head: "Order ID"},
		{Head: "Item"},
		{Head: "Status"},
		{Head: "Popularity"},
	}).GetContent()

	boxInfo := components.Box().
		WithHeadBorder().
		SetHeader("Latest Orders").
		SetHeadColor("#f7f7f7").
		SetBody(table).
		SetFooter(`<div class="clearfix"><a href="javascript:void(0)" class="btn btn-sm btn-info btn-flat pull-left">处理订单</a><a href="javascript:void(0)" class="btn btn-sm btn-default btn-flat pull-right">查看所有新订单</a> </div>`).
		GetContent()

	return types.Panel{
		Content:     row1+boxInfo,
		Title:       "Test",
		Description: "Test example",
	}, nil
}
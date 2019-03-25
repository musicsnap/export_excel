package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
	"strings"
	"github.com/tealeg/xlsx"
	"strconv"
)

type ExcelController struct {
	beego.Controller
}
//XSRF
func (c *ExcelController) Prepare() {
	c.EnableXSRF = false
}

func (c *ExcelController) Excel() {
	beego.BConfig.WebConfig.AutoRender = false
	db:=orm.NewOrm()
	sql:= c.GetString("sql")
	fmt.Println(sql)
	start := time.Now()
	var maps []orm.Params
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")

	if err != nil {
		fmt.Printf(err.Error())
	}

	num, err := db.Raw(sql).Values(&maps)
	if err == nil && num > 0{
		fmt.Println("nums: ", num)
		//fmt.Println("maps:",maps)
		input := c.Input()
		fmt.Println("input:",input)//input: map[4:[银行行号:number] 5:[邮编:zip] sql:[select id,name,phone,addr,number,zip from res_bank_list] 0:[ID:id] 1:[银行名称:name] 2:[电话:phone] 3:[地址:addr]]
		data := make([]string, len(input)-1)
		for key, _ := range input {
			if key == "sql" {
				continue
			}
			i, _ := strconv.Atoi(key)
			data[i] = input.Get(key)

		}
		row = sheet.AddRow()
		style := xlsx.NewStyle()
		style.Font = *xlsx.NewFont(11, "宋体")
		style.Font.Bold = true
		//字体红色
		style.Font.Color = "FFFF0000"

		style.ApplyFont = true
		for _, value := range data {
			split := strings.Split(value, ":")
			cell = row.AddCell()
			cell.Value = split[0]
			cell.SetStyle(style)
		}
		for i := 0; i < int(num); i++ {
			row = sheet.AddRow()
			for _, value := range data {
				split := strings.Split(value, ":")
				if maps[i][split[1]] == nil {
					row.AddCell().Value = ""
				} else {
					cell = row.AddCell()
					cell.Value = maps[i][split[1]].(string)
				}
			}
		}

	}else {
		beego.Error("excel is error")
	}
	file.Write(c.Ctx.ResponseWriter)
	end := time.Since(start)
	fmt.Println(end)
}

package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"github.com/tealeg/xlsx"
	"fmt"
	"time"
	"strconv"
)

type ExcelController struct {
	beego.Controller
}

func (c *ExcelController) Excel() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	sql := c.GetString("sql")
	sql2 := c.GetString("sql2")
	fmt.Println(sql)
	fmt.Println("sql2:", sql2)
	start := time.Now()

	defer func() {
		if len(sql2) > 1 {
			db.Raw(sql2).Exec()
			fmt.Println("defer run,", sql2)
		}

	}()
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")

	if err != nil {
		fmt.Printf(err.Error())
	}

	num, err := db.Raw(sql).Values(&maps)
	if err == nil && num > 0 {
		//这个地方是写入
		//写入表头
		input := c.Input()
		arr := make([]string, len(input)-2)

		//由于map是无序的 这个地方取一个有序的数组来循环
		for key, _ := range input {
			if key == "sql" || key == "sql2" {
				continue
			}

			i, _ := strconv.Atoi(key)
			arr[i] = input.Get(key)

		}

		row = sheet.AddRow()

		style := xlsx.NewStyle()
		style.Font = *xlsx.NewFont(11, "宋体")
		style.Font.Bold = true
		//字体红色
		style.Font.Color = "FFFF0000"

		style.ApplyFont = true

		//首行冻结 好像有点难度

		//		style.Fill = *xlsx.NewFill("none", "FFFF0000", "00000000")
		//		style.ApplyFill = true

		for _, value := range arr {

			split := strings.Split(value, ":")

			cell = row.AddCell()
			cell.Value = split[0]
			cell.SetStyle(style)

		}

		//style.Fill = *xlsx.DefaultFill()

		style = xlsx.NewStyle()
		style.Font = *xlsx.NewFont(11, "宋体")
		style.ApplyFont = true

		for i := 0; i < int(num); i++ {
			row = sheet.AddRow()
			for _, value := range arr {
				split := strings.Split(value, ":")
				if maps[i][split[1]] == nil {
					row.AddCell().Value = ""
				} else {
					cell = row.AddCell()
					cell.SetStyle(style)
					cell.Value = maps[i][split[1]].(string)
				}
			}
		}
	} else {

	}
	file.Write(c.Ctx.ResponseWriter)
	end := time.Since(start)
	fmt.Println(end)
}

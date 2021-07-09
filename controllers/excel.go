package controllers

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	beego "github.com/beego/beego/v2/server/web"
	"reflect"
	"strconv"
	"user/models"
	"user/utils"
)

/**
 * Excel操作
 * 文档地址：https://xuri.me/excelize/zh-hans/
 */
type ExcelController struct {
	beego.Controller
}

// @Title Write
// @Description 导出excel表格
// @Success 200
// @router /excel/write [get]
func (this *ExcelController) Write() {
	f := excelize.NewFile()
	//创建一个工作表
	index := f.NewSheet("学生")
	//设置单元格的值
	f.SetCellValue("学生", "A1", "姓名")
	f.SetCellValue("学生", "B1", "年龄")
	f.SetCellValue("学生", "C1", "邮箱")
	//设置默认的工作表
	f.SetActiveSheet(index)
	//根据指定路径保存文件
	filename := "student.xlsx"
	err := f.SaveAs(filename)
	if err != nil {
		utils.ToJson(this.Controller, err.Error(), "保存错误", "400")
		return
	}
	utils.ToJson(this.Controller, filename, "保存成功", "400")
	return

}

// @Title Read
// @Description 读取excel
// @Success 200
// @router /excel/read [get]
func (this *ExcelController) Read() {
	filename := "article.xlsx"
	f, err := excelize.OpenFile(filename)
	if err != nil {
		utils.ToJson(this.Controller, err, "读取错误", "400")
		return
	}
	//获取指定单元格的值
	//cell := f.GetCellValue("article","A1")

	//获取所有的单元格
	rows := f.GetRows("article")

	utils.ToJson(this.Controller, rows, "读取成功", "200")
	return

}

// @Title Export
// @Description 数据表导出
// @Success 200
// @router /excel/export [get]
func (this *ExcelController) Export() {
	var level []models.UserLevel
	o := models.Open(new(models.UserLevel))
	o.QueryTable(&models.UserLevel{}).All(&level)

	f := excelize.NewFile()
	for index, _ := range level {
		value := reflect.ValueOf(level[index])
		valueType := reflect.TypeOf(level[index])
		for i := 0; i < value.NumField(); i++ { //获取字段数量
			if index == 0 { //写入title
				titleaxis := string(65+i) + "1"
				f.SetCellValue("sheet1", titleaxis, valueType.Field(i).Name)
			}
			//写入数据
			y := strconv.Itoa(index + 2)
			axis := string(65+i) + y
			f.SetCellValue("sheet1", axis, value.Field(i))

		}

	}
	filename := new(models.UserLevel).TableName() + ".xlsx"
	if err := f.SaveAs(filename); err != nil {
		utils.ToJson(this.Controller, err, "保存失败", "500")
		return
	}

	utils.ToJson(this.Controller, filename, "导出成功", "200")
	return
}

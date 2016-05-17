package config

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/speedata/goxlsx"
)

var excel_path = [1]string{
	"static/test.xlsx",
}

func ExcelLoad() {
	for _, v := range excel_path {
		table := &ConfTable{
			name:   "",
			column: map[string]string{},
			rows:   map[string]map[string]interface{}{},
		}
		table.read_excel(v)
		//fmt.Println("excel load: table name: ", table.name, "table: ", table)
		conf_tables[table.name] = table
	}
}

func (table *ConfTable) read_excel(path string) {
	xlsx, err := goxlsx.OpenFile(path)
	if err != nil {
		fmt.Println("open error: ", err)
		panic(err)
	}
	num := xlsx.NumWorksheets()
	for i := 0; i < num; i++ {
		table.read_work_sheet(xlsx, i)
	}
}

func (table *ConfTable) read_work_sheet(xlsx *goxlsx.Spreadsheet, index int) {
	work_sheet, err := xlsx.GetWorksheet(index)
	if err != nil {
		fmt.Println("open work sheet error: ", err)
		panic(err)
	}
	table.name = work_sheet.Name
	if work_sheet.MaxRow <= work_sheet.MinRow+1 || work_sheet.MaxColumn <= work_sheet.MinColumn {
		fmt.Println("open work sheet error: empty config")
		return
	}
	table.excel_parser_column(work_sheet, work_sheet.MinRow)
	//第1行已经提前解析
	for i := work_sheet.MinRow + 1; i < work_sheet.MaxRow; i++ {
		table.excel_parser_rows(work_sheet, i)
	}
}

func (table *ConfTable) excel_parser_column(work_sheet *goxlsx.Worksheet, row int) {
	//第1列为标记列，此处不处理
	for i := work_sheet.MinColumn + 1; i <= work_sheet.MaxColumn; i++ {
		str := work_sheet.Cell(i, row)
		str = strings.TrimSpace(str)
		ss := strings.Split(str, ":")
		table.column[ss[0]] = ss[1]
	}
}

func (table *ConfTable) excel_parser_rows(work_sheet *goxlsx.Worksheet, row int) {
	str := work_sheet.Cell(1, row)
	if str != "#" && str != ";" {
		return
	}
	ss := make([]string, 0, work_sheet.MaxColumn-work_sheet.MinColumn)
	j := 0
	for i := work_sheet.MinColumn + 1; i <= work_sheet.MaxColumn; i++ {
		str = work_sheet.Cell(i, row)
		ss = append(ss, str)
		j++
	}
	row_map := make(map[string]interface{})
	j = 0
	for k, v := range table.column {
		var value interface{}
		switch v {
		case "int":
			k, _ := strconv.Atoi(ss[j])
			value = k
		case "string":
			value = ss[j]
		case "float":
			f, _ := strconv.ParseFloat(ss[j], 32)
			value = float32(f)
		}
		j++
		row_map[k] = value
	}
	table.rows[ss[0]] = row_map
}

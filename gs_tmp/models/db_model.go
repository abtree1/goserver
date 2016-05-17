package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"gs_tmp/config"
	. "gs_tmp/utils"
)

type TableDb struct {
	Static int
	Name   string
	Data   []map[string]interface{}
}

type Operation struct {
	table_name string
	sql        string
	params     []interface{}
	back       bool
	handler    chan<- *TableDb
}

var writes = make(chan *Operation)

func RunDb(exit chan<- bool) {
	p1 := config.GetIniString("database", "type")
	p2 := config.GetIniString("database", "conn")
	db, err := sql.Open(p1, p2)
	if err != nil {
		fmt.Println("conn db error", err.Error())
		panic(err)
	}
	defer func() {
		db.Close()
		exit <- true
	}()
	for {
		write := <-writes
		if write.back {
			write.call_handler(db)
		} else {
			write.cast_handler(db)
		}
	}
}

func (write *Operation) cast_handler(db *sql.DB) error {
	_, err := db.Exec(write.sql, write.params...)
	return err
}

func (write *Operation) call_handler(db *sql.DB) error {
	rows, err := db.Query(write.sql, write.params...)
	if err != nil {
		fmt.Println("QueryErr:", err.Error())
		rows.Close()
		return err
	}
	//fmt.Println("rows")
	stable := GetTableCat(write.table_name)
	length := len(stable)
	//fmt.Println("stable: ", length)
	refs := make([]interface{}, length)
	i := 0
	for _, v := range stable {
		switch v {
		case "int":
			var id int
			refs[i] = &id
		case "string":
			var sd string
			refs[i] = &sd
		}
		i++
	}
	//fmt.Println("refs :", refs[0])
	dts := make([]map[string]interface{}, 0, 1)
	for rows.Next() {
		rows.Scan(refs...)
		i = 0
		db_data := make(map[string]interface{})
		for k, v := range stable {
			switch v {
			case "int":
				db_data[k] = *refs[i].(*int)
			case "string":
				db_data[k] = *refs[i].(*string)
			}
			i++
		}
		dts = append(dts, db_data)
		//fmt.Println("from db id:", rets[0].(int), " name:", rets[1].(string), " pwd:", rets[2].(string), " age:", rets[3].(int))
	}
	//fmt.Println("table_name:", write.table_name, " datas:", dts)
	tb := &TableDb{
		Static: TABLE_LOAD,
		Name:   write.table_name,
		Data:   dts,
	}
	write.handler <- tb
	return nil
}

package models

import (
	. "gs_tmp/utils"
)

func LoadAllPlayerData(playerid int) map[string]*TableDb {
	length := GetTableCount()
	params := make([]interface{}, 0, 1)
	params = append(params, playerid)
	h := make(chan *TableDb, 1)
	for _, k := range GetTableNames() {
		if k == "users" {
			ps := []string{"id"}
			s := LoadTable(k, ps)
			op := &Operation{
				table_name: k,
				sql:        s,
				params:     params,
				back:       true,
				handler:    h,
			}
			writes <- op
		} else {
			ps := []string{"user_id"}
			s := LoadTable(k, ps)
			op := &Operation{
				table_name: k,
				sql:        s,
				params:     params,
				back:       true,
				handler:    h,
			}
			writes <- op
		}
	}
	datas := make(map[string]*TableDb)
	for i := 0; i < length; i++ {
		data := <-h
		//fmt.Println("receive back:", data)
		datas[data.Name] = data
	}
	//fmt.Println("sql tables:", datas)
	close(h)
	return datas
}

func LoadData(table_name string, params []string, values []interface{}) *TableDb {
	h := make(chan *TableDb, 1)
	s := LoadTable(table_name, params)
	op := &Operation{
		table_name: table_name,
		sql:        s,
		params:     values,
		back:       true,
		handler:    h,
	}
	writes <- op
	return <-h
}

package config

import (
	//"fmt"
	"strconv"

	. "gs_tmp/utils"
)

type ConfTable struct {
	name   string
	column map[string]string
	rows   map[string]map[string]interface{}
}

var conf_tables = map[string]*ConfTable{}

func GetLength(table_name string) (int, bool) {
	t, ok := conf_tables[table_name]
	if ok {
		return len(t.rows), true
	}
	return 0, false
}

func Find(table_name string, index interface{}) (map[string]interface{}, bool) {
	t, ok := conf_tables[table_name]
	if ok {
		i := ToString(index)
		r, ok := t.rows[i]
		return r, ok
	}
	return nil, false
}

func GetValue(table_name string, index interface{}, column string) (interface{}, bool) {
	r, ok := Find(table_name, index)
	if ok {
		c, ok := r[column]
		return c, ok
	} else {
		return nil, false
	}
}

func GetValueInt(table_name string, index interface{}, column string) (int, bool) {
	v, ok := GetValue(table_name, index, column)
	if ok {
		return v.(int), true
	} else {
		return 0, false
	}
}

func GetValueString(table_name string, index interface{}, column string) (string, bool) {
	v, ok := GetValue(table_name, index, column)
	if ok {
		return v.(string), true
	} else {
		return "", false
	}
}

func GetValueFloat(table_name string, index interface{}, column string) (float32, bool) {
	v, ok := GetValue(table_name, index, column)
	if ok {
		return v.(float32), true
	} else {
		return 0.0, false
	}
}

func GetIniInt(space string, key string) (int, bool) {
	if space != "" {
		key = "[" + space + "]" + key
	}
	v := ini_map[key]
	i, err := strconv.Atoi(v)
	if err != nil {
		return 0, false
	}
	return i, true
}

func GetIniString(space string, key string) string {
	if space != "" {
		key = "[" + space + "]" + key
	}
	return ini_map[key]
}

func GetInitFloat(space string, key string) (float32, bool) {
	if space != "" {
		key = "[" + space + "]" + key
	}
	v := ini_map[key]
	f, err := strconv.ParseFloat(v, 32)
	if err != nil {
		return 0.0, false
	}
	return float32(f), true
}

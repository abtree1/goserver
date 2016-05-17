package utils

var table_list = map[string]map[string]string{
	"users": map[string]string{
		"id":   "int",
		"name": "string",
		"pwd":  "string",
		"age":  "int",
	},
	"user_conns": map[string]string{
		"id":      "int",
		"phone":   "string",
		"mobile":  "string",
		"email":   "string",
		"qq":      "string",
		"user_id": "int",
	},
}

func GetTableCat(name string) map[string]string {
	return table_list[name]
}

func GetTableCount() int {
	return len(table_list)
}

func GetTableNames() []string {
	keys := make([]string, 0, len(table_list))
	for k, _ := range table_list {
		keys = append(keys, k)
	}
	return keys
}

func LoadTable(table_name string, params []string) string {
	length := len(params)
	sql := "select * from " + table_name
	if length == 1 {
		sql += " where " + params[0] + "=?"
	} else {
		sql += " where "
		for i := 0; i < length-1; i++ {
			sql += params[i] + "=? and "
		}
		sql += params[length-1] + "=?"
	}
	return sql
}

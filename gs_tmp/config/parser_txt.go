package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var txt_path = map[string]string{
	"test": "static/test.conf",
}

func TxtLoad() {
	for k, v := range txt_path {
		table := &ConfTable{
			name:   k,
			column: map[string]string{},
			rows:   map[string]map[string]interface{}{},
		}
		table.read(v)
		conf_tables[k] = table
	}
}

func (table *ConfTable) read(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("open file error: ", err)
		file.Close()
		return
	}
	buf := bufio.NewReader(file)
	l, err := buf.ReadString('\n')
	if err == io.EOF {
		if len(l) == 0 {
			fmt.Println("empty file")
			file.Close()
			return
		}
	} else if err != nil {
		fmt.Println("read file error: ", err.Error())
		file.Close()
		return
	}
	table.parser_column(l)
	for {
		l, err = buf.ReadString('\n')
		if err == io.EOF {
			if len(l) == 0 {
				break
			}
		} else if err != nil {
			fmt.Println("read file error: ", err.Error())
			file.Close()
			return
		}
		if !(strings.HasPrefix(l, "#") || strings.HasPrefix(l, ";")) {
			continue
		}

		table.parser_row(l)
	}
	file.Close()
}

func split_rule(c rune) bool {
	if c == '\t' || c == ' ' {
		return true
	} else {
		return false
	}
}

func (table *ConfTable) parser_column(l string) {
	l = strings.TrimSpace(l)
	strs := strings.FieldsFunc(l, split_rule)
	for _, v := range strs {
		ss := strings.Split(v, ":")
		table.column[ss[0]] = ss[1]
	}
}

func (table *ConfTable) parser_row(l string) {
	l = strings.TrimPrefix(l, "#")
	l = strings.TrimSpace(l)
	strs := strings.FieldsFunc(l, split_rule)
	i := 0
	row := make(map[string]interface{})
	for k, v := range table.column {
		var value interface{}
		switch v {
		case "int":
			j, _ := strconv.Atoi(strs[i])
			value = j
		case "string":
			value = strs[i]
		case "float":
			f, _ := strconv.ParseFloat(strs[i], 32)
			value = float32(f)
		}
		i++
		row[k] = value
	}
	table.rows[strs[0]] = row
}

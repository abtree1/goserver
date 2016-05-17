package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var ini_path = [1]string{
	"static/test.ini",
}

var ini_map = map[string]string{}

func IniLoad() {
	for _, v := range ini_path {
		read_ini(v)
	}
	replace_ini()
	fmt.Println(ini_map)
}

func read_ini(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("open file error: ", err)
		file.Close()
		return
	}
	buf := bufio.NewReader(file)
	title := ""
	for {
		l, err := buf.ReadString('\n')
		if err == io.EOF {
			if len(l) == 0 {
				//fmt.Println("empty file")
				file.Close()
				break
			}
		} else if err != nil {
			fmt.Println("read file error: ", err.Error())
			file.Close()
			return
		}
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			continue
		} else if strings.HasPrefix(l, "#") || strings.HasPrefix(l, ";") {
			continue
		} else if strings.HasPrefix(l, "[") {
			title = l
			continue
		}
		i := strings.Index(l, "=")
		key := strings.TrimSpace(string([]byte(l)[:i]))
		value := strings.TrimSpace(string([]byte(l)[i+1:]))
		value = strings.TrimPrefix(value, "\"")
		value = strings.TrimSuffix(value, "\"")
		if title == "" {
			ini_map[key] = value
		} else {
			key = title + key
			ini_map[key] = value
		}
	}
}

func replace_ini() {
	r, _ := regexp.Compile("%\\(.*\\)")
	for k, v := range ini_map {
		for {
			if r.MatchString(v) {
				v = regexp_ini(r, v)
			} else {
				ini_map[k] = v
				break
			}
		}
	}
}

func regexp_ini(r *regexp.Regexp, v string) string {
	ss := r.FindStringSubmatch(v)
	ii := r.FindStringSubmatchIndex(v)
	str := ""
	if ii[0] > 0 {
		str = string([]byte(v)[:ii[0]])
	}
	s := strings.TrimPrefix(ss[0], "%(")
	s = strings.TrimSuffix(s, ")")
	str += ini_map[s]
	if ii[1] < len(v) {
		str += string([]byte(v)[ii[1]:])
	}
	return str
}

package webengine

import (
	"fmt"
	"strconv"
)

const configPath = "./config.ini"

type config map[string]map[string]string

var c config

func GetConfig(section, key string) string {
	return c[section][key]
}

func GetDBInfo() (ret struct {
	Host   string
	Port   uint16
	User   string
	Pwd    string
	DBName string
}) {
	ret.Host = GetConfig("DATABASE", "Host")
	port, _ := strconv.Atoi(GetConfig("DATABASE", "Port"))
	ret.Port = uint16(port)
	ret.User = GetConfig("DATABASE", "User")
	ret.Pwd = GetConfig("DATABASE", "Pwd")
	ret.DBName = GetConfig("DATABASE", "DBName")
	return ret
}

func GetDBInfoString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset='utf8'",
		GetConfig("DATABASE", "User"),
		GetConfig("DATABASE", "Pwd"),
		GetConfig("DATABASE", "Host"),
		GetConfig("DATABASE", "Port"),
		GetConfig("DATABASE", "DBName"),
	)
}

package commonutil

import (
	"os"
	"strconv"
)

func SetEnvStr(value string, env string) string {
	if e := os.Getenv(env); e != "" {
		return e
	}
	return value
}

func SetEnvInt(value int, env string) int {
	if e := os.Getenv(env); e != "" {
		if res, err := strconv.Atoi(e); err != nil {
			panic("环境变量参数格式错误，无法注入！")
		} else {
			return res
		}
	}
	return value
}

func SetEnvBool(value bool, env string) bool {
	if e := os.Getenv(env); e != "" {
		if res, err := strconv.ParseBool(e); err != nil {
			panic("环境变量参数格式错误，无法注入！")
		} else {
			return res
		}
	}
	return value
}

package gparam

import (
	"os"
)

// GetParam() 返回内容为参数的map
func GetParam() map[string]string {
	param := make(map[string]string)
	args := os.Args[1:]
	for i := 0; i < len(args); i = i + 2 {
		param[args[i]] = args[i+1]
	}
	return param
}

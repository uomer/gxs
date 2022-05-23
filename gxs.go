package main

import (
	"fmt"
	"gxs/config"
	"gxs/gparam"
	"gxs/work"
	"os"
)

// param : -b 基本url
// : -f 写入的文件名
// : -u 页面的url在基本url后面的部分
func main() {
	conf := config.Config
	param := gparam.GetParam()
	base := param["-b"]
	if base == "" {
		base = conf.Base
	}
	purl := param["-u"]
	if purl == "" {
		purl = conf.Purl
	}
	filename := param["-f"]
	mode := param["-m"]
	fileMode := os.O_CREATE | os.O_WRONLY
	if conf.AppendMode || mode == "1" {
		fileMode = fileMode | os.O_APPEND
	} else {
		fileMode = fileMode | os.O_TRUNC
	}
	file, err := os.OpenFile(filename, fileMode, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	work.Work(base, purl, file)
}

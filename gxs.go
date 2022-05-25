package main

import (
	"flag"
	"fmt"
	"gxs/config"
	"gxs/work"
	"os"
)

// param
// : -b 基本url
// : -f 写入的文件名
// : -u 页面的url在基本url后面的部分
// : -a 追加模式
func main() {
	// 命令行参数处理
	cf := flag.String("c", "", "使用指定配置文件")
	base := flag.String("b", "", "网站域名")
	purl := flag.String("u", "", "启始章节链接")
	filename := flag.String("f", "", "生成的小说文件")
	appendMode := flag.Bool("a", false, "追加模式")
	flag.Parse()

	if *filename == "" {
		fmt.Println("没有指定生成的文件")
		return
	}

	var conf = new(config.Conf)

	if *cf != "" {
		conf.UseFile(*cf)
	} else {
		conf.GetConf()
	}

	if *base == "" {
		*base = conf.Base
	}
	if *purl == "" {
		*purl = conf.Purl
	}
	if !*appendMode {
		if conf.AppendMode {
			fmt.Println("启用配置文件中的追回模式")
			*appendMode = true
		} else {
			fmt.Println("启用重建模式")
		}
	}

	// 设置爬取解析dom
	work.SetQuery(&conf.Query)
	// 设置页面编码
	work.SetEncode(conf.Encode)

	// 设置打开文件的模式
	fileMode := os.O_CREATE | os.O_WRONLY
	if conf.AppendMode || *appendMode {
		fileMode = fileMode | os.O_APPEND
	} else {
		fileMode = fileMode | os.O_TRUNC
	}

	// 打开文件
	file, err := os.OpenFile(*filename, fileMode, 0644)
	if err != nil {
		fmt.Printf("打开要写入的文件失败，%s\n", err)
		return
	}
	// 关闭文件
	defer file.Close()

	// 开始爬取
	work.Work(*base, *purl, file)
}

package work

import (
	"fmt"
	"gxs/config"
	"gxs/convert"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var conf *config.Conf

func SetConfig(config *config.Conf) {
	conf = config
}

// 爬小说并写入文件
// 参数:
// base : 基本purl
// purl : 页面url后面的部分
// file : 要写入的文件*os.File
func Work(base, purl string, file *os.File) {
	doc, err := goquery.NewDocument(base + purl)
	if err != nil {
		log.Fatalln(err)
	}
	titleFindArr := conf.Query.Title
	// 章节标题
	titleSel := gSel(doc, titleFindArr)
	title := titleSel.Text()
	// 内容
	contentFindArr := conf.Query.Content
	contentSel := gSel(doc, contentFindArr)
	content := contentSel.Text()
	title = convert.ConvertString(title)
	content = convert.ConvertString(content)
	//
	content = strings.Split(content, "()\n")[0]
	// 写入
	if title != "" {
		file.WriteString(title)
		file.WriteString("\n\n")
		fmt.Println(title)
	}
	if content != "" {
		file.WriteString(content)
		file.WriteString("\n\n")
	}

	// 下一章
	nextFindArr := conf.Query.Next
	nextWithIndex := conf.Query.NextWithIndex
	nextIndex := conf.Query.NextIndex
	nextSel := gSelWithIndex(doc, nextFindArr, nextWithIndex, nextIndex)
	url, ok := nextSel.Attr("href")
	if !ok {
		fmt.Println("没有下一章了")
		return
	}
	Work(base, url, file)
}

package work

import (
	"fmt"
	"gxs/config"
	"gxs/convert"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

var query *config.Query
var encode string

func SetQuery(q *config.Query) {
	query = q
}
func SetEncode(code string) {
	encode = code
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
	titleFindArr := query.Title
	// 章节标题
	titleSel := gSel(doc, titleFindArr)
	title := titleSel.Text()
	// 内容
	contentFindArr := query.Content
	contentSel := gSel(doc, contentFindArr)
	content := contentSel.Text()
	title = convert.ConvertString(title, encode)
	content = convert.ConvertString(content, encode)
	//content = strings.Split(content, "()\n")[0]

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
	nextFindArr := query.Next
	nextWithIndex := query.NextWithIndex
	nextIndex := query.NextIndex
	nextSel := gSelWithIndex(doc, nextFindArr, nextWithIndex, nextIndex)
	url, ok := nextSel.Attr("href")
	if !ok {
		fmt.Println("没有下一章了")
		return
	}
	Work(base, url, file)
}

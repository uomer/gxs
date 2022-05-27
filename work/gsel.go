// gsel.go 查询方法
package work

import (
	"github.com/PuerkitoBio/goquery"
)

// 不使用索引的查询
// doc *goquery.Document doc指针
// findArr 选择器数组切片[]string
func gSel(doc *goquery.Document, findArr []string) *goquery.Selection {
	sel := doc.Find(findArr[0])
	doc.Contents()
	if len(findArr) > 1 {
		for i, v := range findArr {
			if i > 0 {
				sel = sel.Find(v)
			}
		}
	}
	return sel
}

// 使用索引的查询
// doc *goquery.Document doc指针
// findArr 选择器数组切片[]string
// isWith 是否使用索引 bool
// index 索引 uint
func gSelWithIndex(doc *goquery.Document, findArr []string, isWith bool, index uint) *goquery.Selection {
	s := gSel(doc, findArr)
	if isWith {
		for i, n := range s.Nodes {
			if i == int(index) {
				return goquery.NewDocumentFromNode(n).Selection
			}

		}
	}
	return s
}

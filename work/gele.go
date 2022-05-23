package work

import (
	"github.com/PuerkitoBio/goquery"
)

func gSel(doc *goquery.Document, findArr []string) *goquery.Selection {
	sel := doc.Find(findArr[0])
	doc.Contents()
	if len(findArr) > 1 {
		for i, v := range findArr {
			if i > 0 {
				sel.Find(v)
			}
		}
	}
	return sel
}

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

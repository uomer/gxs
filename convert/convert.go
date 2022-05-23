package convert

import (
	"strings"

	"github.com/axgle/mahonia"
)

const code = "gbk"

func ConvertString(src string) string {
	srcCoder := mahonia.NewDecoder(code)
	return strings.Replace(srcCoder.ConvertString(src), "聽", "", -1)
}

package convert

import (
	"strings"

	"github.com/axgle/mahonia"
)

func ConvertString(src string, code string) string {
	srcCoder := mahonia.NewDecoder(code)
	return strings.Replace(srcCoder.ConvertString(src), "聽", "", -1)
}

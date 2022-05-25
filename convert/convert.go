package convert

import (
	"strings"

	"github.com/axgle/mahonia"
)

func ConvertString(src string, code string) string {
	srcCoder := mahonia.NewDecoder(code)
	res := strings.ReplaceAll(srcCoder.ConvertString(src), "聽", "")
	res = strings.TrimSpace(res)
	res = strings.ReplaceAll(res, "    ", "\n")
	return res
}

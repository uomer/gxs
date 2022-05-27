package convert

import (
	"strings"

	"github.com/axgle/mahonia"
)

func ConvertString(src string, code string) string {
	// 新建解码器
	srcCoder := mahonia.NewDecoder(code)
	// 解码并去除乱码
	res := strings.ReplaceAll(srcCoder.ConvertString(src), "聽", "")
	// 去除首尾空格
	res = strings.TrimSpace(res)
	// br相关
	res = strings.ReplaceAll(res, "    ", "\n")
	return res
}

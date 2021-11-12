package string

import (
	"fmt"
	"testing"
)

func TestStringConvert(t *testing.T) {
	var s string = "LEETCODEISHIRIN"
	res := stringConvert(s, 3)
	fmt.Printf("原字符串:%s\n,转后字符串:%s\n", s, res)
	fmt.Printf("正确答案:%s\n", "LCIRETOESIIGEDHN")
}

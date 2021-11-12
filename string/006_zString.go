package string

import (
	"bytes"
)

// 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
/*
	等差数列的通项公式	an = a1 + (n-1)d
	https://blog.csdn.net/qq_33022911/article/details/83616833
*/

func stringConvert(s string, rows int) string {
	if rows == 1 || len(s) < 1 {
		return s
	}
	// bytes.Buffer
	// type Buffer struct {buf []byte ... } 字符切片 下标表示 行数
	bufs := make([]bytes.Buffer, rows)
	// flag == 1 表示行数 +1 / flag == -1 表示行数 -1
	flag := -1
	// Buffer的下标 表示行数
	curr := 0
	for i := 0; i < len(s); i++ {
		bufs[curr].WriteByte(s[i])
		if curr == 0 || curr == rows-1 {
			flag = -flag
		}
		curr += flag
		// row == 3
		// i=0  append s[0] to bufs[0]
		// curr == 0 true
		// flag = -(-1)
		// curr = 1
		// i = 1 append s[1] to bufs[1]
		// curr == 0 false
		// flag = 1
		// curr = 2
		// i = 2 append s[2] to bufs[2]
		// curr == row -1
		// flag = -1
		// curr = 2 -1 = 1
		// i = 3 append s[3] to bufs[1]
		// curr false
		// flag = -1
		// curr = 1 -1 = 0
		// i = 4 append s[4] to bufs[0]
		// ...
	}
	var ans bytes.Buffer
	// []bufs 分别写入 ans
	for _, v := range bufs {
		ans.Write(v.Bytes())
	}
	// 返回字符串
	return ans.String()
}
